import 'package:flame/components.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';
import 'package:orbstrike/Components/GameErrorDialog.dart';

import 'package:orbstrike/Game/PlayerC.dart';
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
  final void Function() closeGame;

  MainUICallbacks({required this.showDial, required this.closeGame});
}

class GameCoreApi {
  ClientChannel? chan;
  GameServiceClient? client;
  BehaviorSubject<Move> stream;
  ChannelCredentials? credentials;

  GameCoreApi({required this.stream, this.chan, this.client, this.credentials});
}

class GameCoreComponents {
  late final CameraComponent mainCamera;

  GameBoard board;
  final World world = World();
  WorldBorder? border;
  PlayerC? mainPlayerComponent;
  Move_Direction mainPlayerDirection;
  UserCredentials? mainPlayerCreds;
  bool mainPlayerRingState;
  final List<int> mainPlayerCollided = [];
  final Map<int, PlayerO> playerComponents = {};

  GameCoreComponents({
    required this.board,
    required this.mainPlayerDirection,
    required this.mainPlayerRingState,
    this.border,
    this.mainPlayerComponent,
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
  int reconnectTimeout = 5;

  if (api.chan != null) {
    api.chan!.shutdown();
  }

  if (coreComponents.mainPlayerCreds==null) {
    // Happens only when wrong developed, so that's why I throw a exception that is caught by flutter
    throw Exception("Main Player Credentials are not initialized!");
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
    print(game.players.length);
    try {
      print("${game.players[0]!.x}");
    } catch (e){}


    if (coreComponents.mainPlayerComponent==null && coreComponents.board.players[coreComponents.mainPlayerCreds!.id]!=null) {
      coreComponents.mainPlayerComponent
          = PlayerC(pPlayer: coreComponents.board.players[coreComponents.mainPlayerCreds!.id]!, collided: coreComponents.mainPlayerCollided);

      coreComponents.world.add(coreComponents.mainPlayerComponent!);
    }
    updateGameBoard(
        coreComponents.board,
        coreComponents.playerComponents,
        coreComponents.mainPlayerCreds!.id,
        coreComponents.mainPlayerComponent
    ).forEach((key, value) {
      if (value) {
        coreComponents.world.add(key);
      } else {
        coreComponents.world.remove(key);
      }
    });

    if (coreComponents.border==null) {
      coreComponents.border = WorldBorder(colors: [Colors.orange, Colors.red], radius: coreComponents.board.rad, stroke: 10);
      coreComponents.world.add(coreComponents.border!);
    }
  }, onError: (error) {
    // gRPC endpoint sends errors for things like "Game Over", "Game not found" etc. -> 400 Errors
    mCallbacks.showDial(error.message, Colors.red);
    mCallbacks.closeGame();
    isExpectedDone = true;
  }, onDone: () {
    if (!isExpectedDone) {
      mCallbacks.showDial("Reconnecting...", Colors.red);
      // gRPC endpoint closes connection when facing a unexpected issue (e.g. read EOF, major failure etc.) -> 500 Errors
      // Wait 5 seconds, on major failures this gives the cluster time to rebuild pods without being spammed.
      Future.delayed(Duration(seconds: reconnectTimeout), () {
        startGameSocket(api, coreComponents, mCallbacks, gameId, host, port);
      });
    }
  });
}

/// Changes the state of the playerComponents and the mainPlayerComponent based on the new GameBoard
///
/// Returns a Map containing Components to add (true) and remove (false) from the world
Map<Component, bool> updateGameBoard(final GameBoard board, final Map<int, PlayerO> playerComps, final int playerID, PlayerC? mainPlayerComponent) {
  Map<Component, bool> componentBuffer = {};

  final player = board.players[playerID];
  if (player!=null && mainPlayerComponent!=null) {
    mainPlayerComponent.pPlayer.x = player.x;
    mainPlayerComponent.pPlayer.y = player.y;
    mainPlayerComponent.pPlayer.kills = player.kills;
    mainPlayerComponent.pPlayer.color = player.color;
    mainPlayerComponent.pPlayer.ringEnabled = player.ringEnabled;
  }

  // Remove players that are not existent anymore
  playerComps.removeWhere((id, component) {
    if (!board.players.containsKey(id)) {
      componentBuffer.putIfAbsent(component, () => false);
      return true;
    }
    return false;
  });

  for (var pPlayer in board.players.values) {
    if (!playerComps.containsKey(pPlayer.id)) {
      final player = PlayerO(pPlayer: pPlayer);
      if (pPlayer.id!=playerID) {
        playerComps[pPlayer.id] = player;
        componentBuffer.putIfAbsent(player, () => true);
      }
    }
  }

  return componentBuffer;
}