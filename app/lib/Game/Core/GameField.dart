import 'dart:async';

import 'package:flame/events.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flame/game.dart';
import 'package:flame/components.dart';
import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';

import 'package:orbstrike/Game/Core/CoreReader.dart';
import 'package:orbstrike/proto/game/game.pbgrpc.dart';
import 'package:orbstrike/Game/Lib/InputParser.dart';
import 'package:orbstrike/Components/DebugComponents.dart';
import 'package:orbstrike/Game/Lib/AuthServiceWrapper.dart';
import 'package:orbstrike/Game/Lib/Types/GameCoreApi.dart';
import 'package:orbstrike/Game/Lib/Types/GameCoreComponents.dart';
import 'package:orbstrike/Game/Lib/Types/MainUICallbacks.dart';

class GameField extends FlameGame with KeyboardEvents, HasCollisionDetection {
  final String host;
  final int port;
  final int gameId;
  final String name;
  final double lerpFactor;
  final bool showDebug;
  final ChannelCredentials? credentials;
  final MainUICallbacks mCallbacks;

  static const P_SPRITESHEETSIZE = 335.0;
  static const P_NORMALSSPATH = "playerNormalSS.png";
  static const P_RINGEDSSPATH = "playerRingedSS.png";
  static const P_RINGANIMATESSPATH = "ringAnimateSS.png";

  GameField({
    required this.gameId,
    required this.name,
    required this.lerpFactor,
    required this.showDebug,
    required this.credentials,
    required this.host,
    required this.port,
    required this.mCallbacks,
  });

  final GameCoreApi coreApi = GameCoreApi();

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


  @override
  Future<void> onLoad() async {
    super.onLoad();
    coreComp.lerpFactor = lerpFactor;

    if (showDebug) {
      final textPaint = TextPaint(
          style: const TextStyle(
            fontSize: 15,
            color: Colors.white60,
          )
      );
      add(IngameFrameCounter()
        ..textRenderer=textPaint);
      add(NetworkFrameCounter(coreApi: coreApi)
        ..textRenderer=textPaint
        ..position=Vector2(0, 30));
    }

    coreComp.mainCamera = CameraComponent(world: coreComp.world);
    addAll([coreComp.mainCamera, coreComp.world]);

    try {
      mCallbacks.setProgress(true, "Loading Assets...");
      coreComp.playerNameSS = SpriteSheet(
          image: await images.load(P_NORMALSSPATH),
          srcSize: Vector2.all(P_SPRITESHEETSIZE)
      );
      coreComp.playerRingedSS = SpriteSheet(
          image: await images.load(P_RINGEDSSPATH),
          srcSize: Vector2.all(P_SPRITESHEETSIZE)
      );
      coreComp.ringAnimateSS = SpriteSheet(
          image: await images.load(P_RINGANIMATESSPATH),
          srcSize: Vector2.all(P_SPRITESHEETSIZE)
      );
      mCallbacks.setProgress(false, "");

      mCallbacks.setProgress(true, "Connecting to Orbstrike Proxy...");
      coreComp.mainPlayerCreds =
          await joinGame(gameId, name, host, port, credentials);

      startCoreReader(coreApi, coreComp, mCallbacks, gameId, host, port);
    } catch (err) {
      String errMsg;
      dynamic dynErr = err;
      if (dynErr is Exception && (dynErr as dynamic).message != null) {
        errMsg = (dynErr as dynamic).message;
      } else {
        errMsg = err.toString();
      }
      mCallbacks.showDial(errMsg, const Color.fromRGBO(222, 4, 4, 1));
      return;
    }

    coreApi.stream?.add(Move(
        direction: Move_Direction.NONE,
        enableRing: false,
        gameid: gameId,
        userkey: coreComp.mainPlayerCreds?.key
    ));
  }

  // Accumulator for time to achieve a constant sendinginterval
  double deltaAccumulator = 0.0;
  final double senderInterval = 0.10;

  @override
  void update(double dt) {
    super.update(dt);

    deltaAccumulator += dt;
    while (deltaAccumulator >= senderInterval) {
      coreApi.stream?.add(Move(
          direction: coreComp.mainPlayerDirection,
          enableRing: coreComp.mainPlayerRingState,
          hitPlayers: coreComp.mainPlayerCollided,
          gameid: gameId, userkey: coreComp.mainPlayerCreds?.key
      ));
      deltaAccumulator -= senderInterval;
    }

    if (coreComp.mainPlayerComponent!=null) {
      coreComp.mainCamera.follow(coreComp.mainPlayerComponent!);
    }
  }

  @override
  void onDispose() async {
    try {
      coreApi.stream?.close();
      coreApi.chan?.shutdown();
      if (coreComp.mainPlayerCreds!=null) {
        await exitGame(gameId, coreComp.mainPlayerCreds, host, port, credentials);
      }
    } catch (err) {
      if (kDebugMode) {
        print("$err");
      }
    }
    super.onDispose();
  }
}