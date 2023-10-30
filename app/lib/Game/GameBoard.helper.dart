import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';

import 'package:orbstrike/Game/Player.dart';
import 'package:orbstrike/Game/WorldBackground.dart';
import 'package:orbstrike/Game/WorldBorder.dart';
import 'package:rxdart/rxdart.dart';

import '../proto/auth/auth.pbgrpc.dart';
import '../proto/game/game.pbgrpc.dart';

class UserCredentials {
  int id;
  List<int> key;

  UserCredentials({required this.id, required this.key});
}

class MainUICallbacks {
  final void Function(String message, Color color) showDial;
  final void Function(String message, Color color) showSnack;
  final void Function(bool loading, String message) setProgress;

  MainUICallbacks({required this.showDial, required this.showSnack, required this.setProgress});
}

class GameCoreApi {
  ClientChannel? chan;
  GameServiceClient? client;
  BehaviorSubject<Move> stream;
  ChannelCredentials? credentials;
  bool shutdown = false;

  GameCoreApi({required this.stream, this.chan, this.client, this.credentials});
}

class GameCoreComponents {
  late final CameraComponent mainCamera;

  GameBoard board;
  final World world = World();
  WorldBorder? border;
  WorldBackground? background;
  MainPlayerComponent? mainPlayerComponent;
  Move_Direction mainPlayerDirection;
  UserCredentials? mainPlayerCreds;
  bool mainPlayerRingState;
  final List<int> mainPlayerCollided = [];
  final Map<int, EnemyPlayerComponent> playerComponents = {};
  late SpriteSheet playerNameSS;
  late SpriteSheet playerRingedSS;
  late SpriteSheet ringAnimateSS;

  GameCoreComponents({
    required this.board,
    required this.mainPlayerDirection,
    required this.mainPlayerRingState,
    this.mainPlayerCreds,
  });
}

Future<UserCredentials> joinGame(int gameId, String name, String host, int port, ChannelCredentials? credentials) async {
  final client = AuthServiceClient(
      ClientChannel(
          host,
          port: port,
          options: ChannelOptions(credentials: credentials ?? const ChannelCredentials.insecure())
      )
  );

  final req = JoinGameRequest(
    gameid: gameId,
    name:  name,
  );
  final res = await client.joinGame(req);
  return UserCredentials(
      id: res.userid,
      key: res.userkey,
  );
}

Future<void> exitGame(int gameId, UserCredentials? userCreds, String host, int port, ChannelCredentials? credentials) async {
  final client = AuthServiceClient(
    ClientChannel(
      host,
      port: port,
      options: ChannelOptions(credentials: credentials ?? const ChannelCredentials.insecure())
    )
  );

  final req = ExitGameRequest(
    gameid: gameId,
    userkey: userCreds?.key
  );
  await client.exitGame(req);
}

void startGameSocket(GameCoreApi api, GameCoreComponents coreComponents, MainUICallbacks mCallbacks, int gameId, String host, int port) {
  bool isExpectedDone = false;
  int reconnectTimeout = 3;

  if (api.chan != null) {
    api.chan!.shutdown();
  }

  api.chan = ClientChannel(
    host,
    port: port,
    options: ChannelOptions(
      credentials: api.credentials ?? const ChannelCredentials.insecure(),
    ),
  );

  api.client = GameServiceClient(api.chan!);

  final Map<String, String> streamHeaders = {
    "gameid": "$gameId",
  };

  final callOptions = CallOptions(metadata: streamHeaders);

  api.client!.proxyGameboard(api.stream.stream, options: callOptions)
      .listen((game) {
    coreComponents.board = game;

    updateGameBoard(coreComponents).forEach((key, value) {
      if (value) {
        coreComponents.world.add(key);
      } else {
        coreComponents.world.remove(key);
      }
    });
  }, onError: (error) {
    // gRPC endpoint sends errors for things like "Game Over", "Game not found" etc. -> 400 Errors
    mCallbacks.showDial(error.message, const Color.fromRGBO(222, 4, 4, 1));
    api.shutdown = true;
  }, onDone: () {
    if (!api.shutdown) {
      mCallbacks.showSnack("Reconnecting...", Colors.red);
      // gRPC endpoint closes connection when facing a unexpected issue (e.g. read EOF, major failure etc.) -> 500 Errors
      // Wait {reconnectTimeout} seconds, on major failures this gives the cluster time to rebuild pods without being spammed.
      Future.delayed(Duration(seconds: reconnectTimeout), () {
        startGameSocket(api, coreComponents, mCallbacks, gameId, host, port);
      });
    }
  });
}

/// Changes the state of the playerComponents and the mainPlayerComponent based on the new GameBoard
///
/// Returns a Map containing Components to add (true) and remove (false) from the world
Map<Component, bool> updateGameBoard(final GameCoreComponents coreComp) {
  Map<Component, bool> componentBuffer = {};

  // Create main component if not existent
  if (coreComp.mainPlayerComponent==null && coreComp.board.players[coreComp.mainPlayerCreds?.id]!=null) {
    coreComp.mainPlayerComponent
      = MainPlayerComponent(
      networkPlayerRep: coreComp.board.players[coreComp.mainPlayerCreds?.id]!,
      collided: coreComp.mainPlayerCollided,
      playerNormalSS: coreComp.playerNameSS,
      playerRingedSS: coreComp.playerRingedSS,
      ringAnimateSS: coreComp.ringAnimateSS,
    );

    coreComp.world.add(coreComp.mainPlayerComponent!);
  }

  // Update main component
  final player = coreComp.board.players[coreComp.mainPlayerCreds?.id];
  if (player!=null && coreComp.mainPlayerComponent!=null) {
    coreComp.mainPlayerComponent?.networkPlayerRep = player;
  }

  // Remove players that are not existent anymore
  coreComp.playerComponents.removeWhere((id, component) {
    if (!coreComp.board.players.containsKey(id)) {
      componentBuffer.putIfAbsent(component, () => false);
      return true;
    }
    return false;
  });

  // Add players that were added and update the existing
  for (var networkPlayerRep in coreComp.board.players.values) {
    if (!coreComp.playerComponents.containsKey(networkPlayerRep.id)) {
      final player = EnemyPlayerComponent(
        networkPlayerRep: networkPlayerRep,
        playerNormalSS: coreComp.playerNameSS,
        playerRingedSS: coreComp.playerRingedSS,
        ringAnimateSS: coreComp.ringAnimateSS,
      );
      if (networkPlayerRep.id!=coreComp.mainPlayerCreds?.id) {
        coreComp.playerComponents[networkPlayerRep.id] = player;
        componentBuffer.putIfAbsent(player, () => true);
      }
    } else {
      if (networkPlayerRep.id!=coreComp.mainPlayerCreds?.id) {
        coreComp.playerComponents[networkPlayerRep.id]?.networkPlayerRep = networkPlayerRep;
      }
    }
  }

  // Update Game border if necessary
  if (coreComp.border==null) {
    coreComp.border = WorldBorder(colors: [Colors.orange, Colors.red], radius: coreComp.board.rad, stroke: 10);
    coreComp.world.add(coreComp.border!);
  }

  // Update Game background if necessary
  if (coreComp.background==null) {
    coreComp.background = WorldBackground(
      radius: coreComp.board.rad,
      rectSize: 5,
      rectBorderRadius: 12,
      rectPaint: Paint()..color=const Color.fromRGBO(54,57,62, 0.7),
      rectSpacing: 50,
    );
    coreComp.world.add(coreComp.background!);
  }

  return componentBuffer;
}