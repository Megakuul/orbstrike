import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flame/game.dart';
import 'package:grpc/grpc.dart';

import 'package:orbstrike/Components/Dialogs/GameErrorDialog.dart';
import 'package:orbstrike/Game/Core/GameField.dart';
import 'package:orbstrike/Game/Lib/Types/MainUICallbacks.dart';

class GameOverlay extends StatefulWidget {
  final String host;
  final int port;
  final int gameId;
  final String name;
  final double lerpFactor;
  final bool showDebug;
  final ChannelCredentials? credentials;
  const GameOverlay({super.key,
    required this.host,
    required this.port,
    required this.gameId,
    required this.name,
    required this.lerpFactor,
    required this.showDebug,
    required this.credentials
  });

  @override
  State<GameOverlay> createState() => _GameOverlay();
}

class _GameOverlay extends State<GameOverlay> {
  ValueNotifier<bool> isLoading = ValueNotifier<bool>(false);
  String loadingMsg = "";
  String loadingMsgDetail = "";

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
                lerpFactor: widget.lerpFactor,
                showDebug: widget.showDebug,
                credentials: widget.credentials,
                mCallbacks: MainUICallbacks(
                    showDial: (message, color) {
                      showDialog (
                          context: context,
                          builder: (context) {
                            return GameErrorDialog(
                              message: message.toString(),
                              color: color,
                            );
                          }
                      );
                    },
                    showSnack: (message, color) {
                      ScaffoldMessenger.of(context).showSnackBar(
                          SnackBar(
                              backgroundColor: color,
                              content: Text("$message")
                          )
                      );
                    },
                    setProgress: (loading, msg, [msgDetailed]) {
                      WidgetsBinding.instance.addPostFrameCallback((_) {
                        loadingMsg = msg;
                        loadingMsgDetail = msgDetailed ?? "";
                        isLoading.value = loading;
                      });
                    }
                ),
              )
          ), // Your FlameGame widget
          ValueListenableBuilder<bool>(
            valueListenable: isLoading,
            builder: (context, loading, child) {
              if (loading) {
                return Positioned.fill(
                  child: Container(
                    color: Colors.black.withOpacity(0.8),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        const CupertinoActivityIndicator(radius: 25),
                        const SizedBox(height: 20),
                        Text(loadingMsg,style: const TextStyle(color: Colors.white60, fontSize: 20)),
                        const SizedBox(height: 15),
                        Text(loadingMsgDetail, style: const TextStyle(color: Colors.white60, fontSize: 10))
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