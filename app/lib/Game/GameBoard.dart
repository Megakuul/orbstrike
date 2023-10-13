import 'dart:async';

import 'package:flame/events.dart';
import 'package:flutter/material.dart';
import 'package:flame/game.dart';
import 'package:flame/components.dart';
import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';
import 'package:orbstrike/proto/game/game.pbgrpc.dart';
import 'package:rxdart/rxdart.dart';

import 'package:orbstrike/Game/KeyboardHandler.dart';
import 'package:orbstrike/Game/PlayerC.dart';
import 'package:orbstrike/Game/WorldBorder.dart';
import 'package:orbstrike/Game/GameBoard.helper.dart';

// This will be somewhere else
const playerID = 34234;

class GameField extends FlameGame with KeyboardEvents, HasCollisionDetection {
  final String host;
  final int port;
  final int gameID;
  final BuildContext hudContext;

  final World world = World();
  final Map<int, PlayerO> playerComponents = {};

  GameField({required this.hudContext, required this.gameID, required this.host, required this.port});

  late final CameraComponent mainCamera;

  List<int>? userKey;
  ClientChannel? apiChan;
  GameServiceClient? apiClient;
  BehaviorSubject<Move> apiReqStream = BehaviorSubject<Move>();

  // Border size is loaded when connected to the gRPC endpoint
  GameBoard board = GameBoard();
  WorldBorder? border;
  PlayerC? mainPlayerComponent;
  Move_Direction mainPlayerDirection = Move_Direction.NONE;
  bool mainPlayerRingState = false;
  List<int> mainPlayerCollided = [];

  @override
  Color backgroundColor() => Colors.black12;

  @override
  KeyEventResult onKeyEvent(
    RawKeyEvent event,
    Set<LogicalKeyboardKey> keysPressed,
  ) {
    mainPlayerRingState = getRingState();
    final dirBuf = getMovementInput();
    if (dirBuf==null) {
      return KeyEventResult.ignored;
    }
    mainPlayerDirection = dirBuf;
    return KeyEventResult.handled;
  }

  double keystrokeAccumulator = 0.0;
  final double keystrokeUpdateInterval = 0.10;

  @override
  void update(double dt) {
    super.update(dt);

    keystrokeAccumulator += dt;
    while (keystrokeAccumulator >= keystrokeUpdateInterval) {
      apiReqStream.add(Move(
          direction: mainPlayerDirection,
          enableRing: mainPlayerRingState,
          hitPlayers: mainPlayerCollided,
          gameid: gameID, userkey: userKey
      ));
      keystrokeAccumulator -= keystrokeUpdateInterval;
    }

    if (mainPlayerComponent!=null) {
      mainCamera.follow(mainPlayerComponent!);
    }
  }

  @override
  Future<void> onLoad() async {
    mainCamera = CameraComponent(world: world);
    addAll([mainCamera, world]);

    _startGrpcConnection();

    apiReqStream.add(Move(direction: Move_Direction.NONE, enableRing: false, gameid: gameID, userkey: userKey));
  }

  void _startGrpcConnection() {
    bool isExpectedDone = false;

    if (apiChan != null) {
      apiChan!.shutdown();
    }

    apiChan = ClientChannel(
      host,
      port: port,
      options: const ChannelOptions(
        credentials: ChannelCredentials.insecure(),
      ),
    );

    apiClient = GameServiceClient(apiChan!);

    final Map<String, String> streamHeaders = {
      "gameid": "187",
    };

    final callOptions = CallOptions(metadata: streamHeaders);

    apiClient!.proxyGameboard(apiReqStream.stream, options: callOptions)
        .listen((game) {
      board = game;
      if (mainPlayerComponent==null && board.players[playerID]!=null) {
        mainPlayerComponent = PlayerC(pPlayer: board.players[playerID]!, collided: mainPlayerCollided);
        world.add(mainPlayerComponent!);
      }
      updateGameBoard(board, playerComponents, playerID, mainPlayerComponent).forEach((key, value) {
        if (value) {
          world.add(key);
        } else {
          world.remove(key);
        }
      });

      if (border==null) {
        border = WorldBorder(colors: [Colors.orange, Colors.red], radius: board.rad, stroke: 10);
        world.add(border!);
      }
    }, onError: (error) {
      // gRPC endpoint sends errors for things like "Game Over", "Game not found" etc. -> 400 Errors
      // TODO: These things are displayed with widgets using the context from the top level ui (but this is not implemented right now).
      print('Error encountered: ${error.message}');
      isExpectedDone = true;
    }, onDone: () {
      if (!isExpectedDone) {
        // gRPC endpoint closes connection when facing a unexpected issue (e.g. read EOF, major failure etc.) -> 500 Errors
        // Wait around 5 seconds, on major failures this gives the cluster time to rebuild pods without being spammed.
        print("Server closed connection unexpected");
        Future.delayed(const Duration(seconds: 5), () {
          _startGrpcConnection();
        });
      }
    });
  }
}






