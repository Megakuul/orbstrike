import 'dart:async';

import 'package:flame/events.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flame/game.dart';
import 'package:flame/components.dart';
import 'package:flutter/services.dart';
import 'package:grpc/grpc.dart';
import 'package:orbstrike/proto/game/game.pbgrpc.dart';
import 'package:rxdart/rxdart.dart';

import 'package:orbstrike/Game/KeyboardHandler.dart';
import 'package:orbstrike/Game/GameBoard.helper.dart';

import '../Components/GameErrorDialog.dart';

class GameOverlay extends StatefulWidget {
  final String host;
  final int port;
  final int gameId;
  final String name;
  final ChannelCredentials? credentials;
  const GameOverlay({super.key,
    required this.host,
    required this.port,
    required this.gameId,
    required this.name,
    required this.credentials
  });

  @override
  State<GameOverlay> createState() => _GameOverlay();
}

class _GameOverlay extends State<GameOverlay> {
  ValueNotifier<bool> isLoading = ValueNotifier<bool>(false);
  String loadingMsg = "";

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromRGBO(30,33,36, 1),
      body: Stack(
        children: [
          GameWidget(
            game: GameField(
              gameId: widget.gameId,
              name: widget.name,
              host: widget.host,
              port: widget.port,
              credentials: widget.credentials,
              mCallbacks: MainUICallbacks(
                showDial: (message, color) {
                  showDialog (
                    context: context,
                    builder: (context) {
                      return GameErrorDialog(
                        message: message,
                        color: color,
                      );
                    }
                  );
                },
                showSnack: (message, color) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      backgroundColor: color,
                      content: Text(message)
                    )
                  );
                },
                setProgress: (loading, msg) {
                  WidgetsBinding.instance.addPostFrameCallback((_) {
                    loadingMsg = msg;
                    isLoading.value = loading;
                  });
                }
              )
            )
          ), // Your FlameGame widget
          ValueListenableBuilder<bool>(
            valueListenable: isLoading,
            builder: (context, loading, child) {
              if (loading) {
                return Positioned.fill(
                  child: Container(
                    color: Colors.black.withOpacity(0.5),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        const CupertinoActivityIndicator(radius: 25),
                        const SizedBox(height: 20),
                        Text(loadingMsg,style: const TextStyle(color: Colors.white60))
                      ],
                    ),
                  ),
                );
              } else {
                return const SizedBox.shrink();
              }
            },
          )
        ],
      ),
      floatingActionButton: FloatingActionButton(
        tooltip: "Leave the game",
        child: const Icon(Icons.exit_to_app_outlined),
        onPressed: () {
          Navigator.of(context).pop();
        },
      ),
    );
  }
}


class GameField extends FlameGame with KeyboardEvents, HasCollisionDetection {
  final String host;
  final int port;
  final int gameId;
  final String name;
  final ChannelCredentials? credentials;
  final MainUICallbacks mCallbacks;

  static const P_SPRITESHEETSIZE = 335.0;
  static const P_SPRITESHEETPATH = "playerSpriteSheet.png";

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
    coreComp.mainCamera = CameraComponent(world: coreComp.world);
    addAll([coreComp.mainCamera, coreComp.world]);

    try {
      mCallbacks.setProgress(true, "Loading Assets...");
      final imageInstance = await images.load(P_SPRITESHEETPATH);
      coreComp.playerSpriteSheet = SpriteSheet(image: imageInstance,srcSize: Vector2.all(P_SPRITESHEETSIZE));
      mCallbacks.setProgress(false, "");

      mCallbacks.setProgress(true, "Connecting to Orbstrike Proxy...");
      coreComp.mainPlayerCreds = await joinGame(gameId, name, host, port, credentials);
      mCallbacks.setProgress(false, "");

      startGameSocket(coreApi, coreComp, mCallbacks, gameId, host, port);
    } catch (err) {
      String errMsg;
      dynamic dynErr = err;
      if (dynErr is Exception && (dynErr as dynamic).message != null) {
        errMsg = (dynErr as dynamic).message;
      } else { errMsg = err.toString(); }
      mCallbacks.showDial(errMsg, const Color.fromRGBO(222, 4, 4, 1));
      return;
    }

    coreApi.stream.add(Move(
        direction: Move_Direction.NONE,
        enableRing: false,
        gameid: gameId,
        userkey: coreComp.mainPlayerCreds?.key
    ));
    super.onLoad();
  }

  @override
  void onDispose() async {
    coreApi.shutdown = true;
    try {
      coreApi.stream.close();
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






