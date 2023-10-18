import 'dart:async';
import 'dart:io';

import 'package:flame/events.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flame/game.dart';
import 'package:flame/components.dart';
import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';
import 'package:orbstrike/Components/GameErrorDialog.dart';
import 'package:orbstrike/proto/game/game.pbgrpc.dart';
import 'package:rxdart/rxdart.dart';

import 'package:orbstrike/Game/KeyboardHandler.dart';
import 'package:orbstrike/Game/GameBoard.helper.dart';

class GameField extends FlameGame with KeyboardEvents, HasCollisionDetection {
  final String host;
  final int port;
  final int gameId;
  final String name;
  final ChannelCredentials? credentials;
  final MainUICallbacks mCallbacks;

  GameField({
    required this.gameId,
    required this.name,
    required this.credentials,
    required this.host,
    required this.port,
    required this.mCallbacks,
  });

  final GameCoreApi coreApi = GameCoreApi(
    stream: BehaviorSubject<Move>(),
  );

  final GameCoreComponents coreComp = GameCoreComponents(
    board: GameBoard(),
    mainPlayerDirection: Move_Direction.NONE,
    mainPlayerRingState: false,
  );

  @override
  Color backgroundColor() => Colors.black12;

  @override
  KeyEventResult onKeyEvent(
    RawKeyEvent event,
    Set<LogicalKeyboardKey> keysPressed,
  ) {
    coreComp.mainPlayerRingState = getRingState();
    final dirBuf = getMovementInput();
    if (dirBuf==null) {
      return KeyEventResult.ignored;
    }
    coreComp.mainPlayerDirection = dirBuf;
    return KeyEventResult.handled;
  }

  double keystrokeAccumulator = 0.0;
  final double keystrokeUpdateInterval = 0.10;

  @override
  void update(double dt) {
    super.update(dt);

    keystrokeAccumulator += dt;
    while (keystrokeAccumulator >= keystrokeUpdateInterval) {
      coreApi.stream.add(Move(
          direction: coreComp.mainPlayerDirection,
          enableRing: coreComp.mainPlayerRingState,
          hitPlayers: coreComp.mainPlayerCollided,
          gameid: gameId, userkey: coreComp.mainPlayerCreds?.key
      ));
      keystrokeAccumulator -= keystrokeUpdateInterval;
    }

    if (coreComp.mainPlayerComponent!=null) {
      coreComp.mainCamera.follow(coreComp.mainPlayerComponent!);
    }
  }

  @override
  Future<void> onLoad() async {
    super.onLoad();
    coreComp.mainCamera = CameraComponent(world: coreComp.world);
    addAll([coreComp.mainCamera, coreComp.world]);

    try {
      coreComp.mainPlayerCreds = await joinGame(gameId, name, host, port, credentials);
      startGameSocket(coreApi, coreComp, mCallbacks, gameId, host, port);
    } catch (err) {
      mCallbacks.showDial(err.toString(), Colors.red);
      mCallbacks.closeGame();
    }

    coreApi.stream.add(Move(
        direction: Move_Direction.NONE,
        enableRing: false,
        gameid: gameId,
        userkey: coreComp.mainPlayerCreds?.key
    ));
  }

  @override
  Future<void> onDispose() async {
    super.onDispose();

    try {
      if (coreComp.mainPlayerCreds!=null) {
        await exitGame(gameId, coreComp.mainPlayerCreds, host, port, credentials);
      }
      coreApi.stream.close();
      coreApi.chan?.shutdown();
    } catch (err) {
      if (kDebugMode) {
        print("$err");
      }
    }
  }
}






