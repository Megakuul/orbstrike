import 'package:grpc/grpc.dart';
import 'package:flutter/material.dart';
import 'package:orbstrike/Components/CreateGameDialog.dart';
import 'package:orbstrike/Components/ElevatedIconButton.dart';
import 'package:orbstrike/Components/GradientText.dart';
import 'package:orbstrike/Components/InputField.dart';
import 'package:orbstrike/proto/auth/auth.pbgrpc.dart';

import 'Game/GameBoard.dart';

final GlobalKey gameKey = GlobalKey();

class GameConfiguration {
  GameConfiguration({required this.hostname, required this.port, required this.gameID, required this.latestGameIDs, required this.credentials});
  String hostname;
  int port;
  int gameID;
  ChannelCredentials? credentials;
  List<int> latestGameIDs;
}

Future<int> createGame(GameConfiguration conf,
    double radius,
    double plRadius,
    double plRingRadius,
    double speed,
    int maxPlayers,
    ) async {
  final client = AuthServiceClient(
    ClientChannel(
      conf.hostname,
      port: conf.port,
      options: ChannelOptions(credentials: conf.credentials ?? const ChannelCredentials.insecure())
    )
  );

  final req = CreateGameRequest(
    identifier: "serialnumber",
    radius: radius,
    playerradius: plRadius,
    playerringradius: plRingRadius,
    speed: speed,
    maxplayers: maxPlayers,
  );
  final res = await client.createGame(req);
  return res.gameid;
}

void joinGame(BuildContext context, GameConfiguration conf, String name) {
  Navigator.of(context).push(
    MaterialPageRoute(
      builder: (context) => GameOverlay(
        host: conf.hostname,
        port: conf.port,
        gameId: conf.gameID,
        name: name,
        credentials: conf.credentials,
      )
    )
  );
}

class MainPage extends StatefulWidget {
  final String title;
  final GameConfiguration gameConfig;
  const MainPage({Key? key, required this.title, required this.gameConfig}) : super(key: key);

  @override
  State<MainPage> createState() => _MainPage();
}

class _MainPage extends State<MainPage> {
  final gameIdController = TextEditingController();
  final hostController = TextEditingController();
  final portController = TextEditingController();

  final gameRadiusController = TextEditingController();
  final gamePlayerRadiusController = TextEditingController();
  final gamePlayerRingRadiusController = TextEditingController();
  final gameSpeedController = TextEditingController();
  final gameMaxPlayerController = TextEditingController();

  final playerName = TextEditingController();

  @override
  void initState() {
    super.initState();

    final gameIdStr = widget.gameConfig.gameID.toString();
    gameIdController.text = gameIdStr=="-1" ? "" : gameIdStr;
    hostController.text = widget.gameConfig.hostname;
    portController.text = widget.gameConfig.port.toString();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromRGBO(66,69,73, 0.9),
      body: Column(
        children: [
          Padding(
            padding: const EdgeInsets.only(top: 100, bottom: 20),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                GradientText(
                  text: const Text(
                    "Orbstrike",
                    style: TextStyle(fontSize: 70, fontWeight: FontWeight.w500)
                  ),
                  gradient: const LinearGradient(
                    colors: [
                      Color.fromRGBO(128, 179, 255, 1),
                      Color.fromRGBO(152, 228, 255, 1),
                      Color.fromRGBO(182, 255, 250, 1)
                    ],
                    begin: Alignment.topLeft,
                    end: Alignment.bottomRight
                  )
                ),
                const SizedBox(width: 20),
                Image.asset(
                  "assets/icons/orbstrike.png",
                  width: 90,
                  height: 90,
                  fit: BoxFit.cover,
                )
              ],
            ),
          ),

          Container(
            alignment: Alignment.center,
            width: 250,
            child: InputField(
              mx: 5, my: 5,
              hint: "Game Identifier",
              controller: gameIdController,
              radius: 12,
              onChanged: (_) {
                widget.gameConfig.gameID =
                    int.tryParse(gameIdController.text) ?? widget.gameConfig.gameID;
              },
            ),
          ),
          Container(
            constraints: const BoxConstraints(maxWidth: 700),
            child: Wrap(
              crossAxisAlignment: WrapCrossAlignment.center,
              children: [
                ElevatedIconButton(
                  text: const Text("Create Game", style: TextStyle(fontSize: 25)),
                  icon: const Icon(Icons.create),
                  mx: 20, my: 20,
                  gradient: const LinearGradient(
                      colors: [
                        Color.fromRGBO(54,57,62, 0.7),
                        Color.fromRGBO(46,49,54, 0.7),
                        Color.fromRGBO(30,33,36, 0.5),
                      ],
                      begin: Alignment.bottomLeft,
                      end: Alignment.topRight
                  ),
                  onPressed: () async {
                    try {
                      final res = await showDialog<String>(
                        context: context,
                        builder: (BuildContext context) {
                          return CreateGameDialog(
                            radiusController: gameRadiusController,
                            plRadiusController: gamePlayerRadiusController,
                            plRingRadiusController: gamePlayerRingRadiusController,
                            speedController: gameSpeedController,
                            maxPlayerController: gameMaxPlayerController,
                          );
                        }
                      );

                      if (res!="Create") { return; }

                      final double? radius =
                        double.tryParse(gameRadiusController.text);
                      if (radius==null) {
                        throw Exception("Map Size must be a number!");
                      }
                      final double? plRadius =
                      double.tryParse(gamePlayerRadiusController.text);
                      if (plRadius==null) {
                        throw Exception("Player Size must be a number!");
                      }
                      final double? plRingRadius =
                      double.tryParse(gamePlayerRingRadiusController.text);
                      if (plRingRadius==null) {
                        throw Exception("Player Ring Size must be a number!");
                      }
                      final double? speed =
                      double.tryParse(gameSpeedController.text);
                      if (speed==null) {
                        throw Exception("Speed must be a number!");
                      }
                      final int? maxPlayers =
                        int.tryParse(gameMaxPlayerController.text);
                      if (maxPlayers==null) {
                        throw Exception("Max Players must be a number!");
                      }

                      final gameId = await createGame(
                          widget.gameConfig,
                          radius,
                          plRadius,
                          plRingRadius,
                          speed,
                          maxPlayers
                      );

                      setState(() {
                        widget.gameConfig.latestGameIDs.insert(0, gameId);
                      });
                    } catch (err) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          backgroundColor: Colors.red,
                          content: Text("ERROR: $err")
                        )
                      );
                    }
                  },
                ),
                ElevatedIconButton(
                  text: const Text("Join Game", style: TextStyle(fontSize: 25)),
                  icon: const Icon(Icons.directions_run),
                  mx: 20, my: 20,
                  gradient: const LinearGradient(
                      colors: [
                        Color.fromRGBO(54,57,62, 0.7),
                        Color.fromRGBO(46,49,54, 0.7),
                        Color.fromRGBO(30,33,36, 0.5),
                      ],
                      begin: Alignment.bottomLeft,
                      end: Alignment.topRight
                  ),
                  onPressed: () {
                    widget.gameConfig.gameID = int.tryParse(gameIdController.text) ?? 0;
                    joinGame(context, widget.gameConfig, playerName.text);
                  },
                ),
              ],
            ),
          ),

          Expanded(
            child: Container(
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(12),
                color: const Color.fromRGBO(255,255,255, 0.02)
              ),
              width: MediaQuery.of(context).size.width / 1.5,
              margin: const EdgeInsets.symmetric(vertical: 10),
              padding: const EdgeInsets.symmetric(horizontal: 4, vertical: 5),
              child: ListView.builder(
                itemCount: widget.gameConfig.latestGameIDs.length,
                itemBuilder: (context, i) {
                  final id = widget.gameConfig.latestGameIDs[i];
                  return ListTile(
                    leading: IconButton(
                      icon: const Icon(Icons.play_circle_outline_rounded),
                      color: Colors.white38,
                      splashRadius: 22,
                      hoverColor: Colors.white10,
                      onPressed: () {
                        setState(() {
                          gameIdController.text = id.toString();
                        });
                      },
                    ),
                    title: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        const Text("Game Identifier",
                            style: TextStyle(color: Colors.white60, fontSize: 20)),
                        Text("$id",
                            style: const TextStyle(color: Colors.white60, fontSize: 20)),
                      ],
                    )
                  );
                }
              ),
            )
          )
        ],
      ),
      bottomNavigationBar: Align(
        heightFactor: 1,
        alignment: Alignment.bottomCenter,
        child: Container(
          margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 15),
          width: 600,
          child: Row(
            children: [
              Expanded(
                flex: 7,
                child: InputField(
                  mx: 5, my: 5,
                  hint: "Server Hostname",
                  controller: hostController,
                  radius: 12,
                  onChanged: (_) {
                    widget.gameConfig.hostname = hostController.text;
                  },
                ),
              ),
              Expanded(
                flex: 3,
                child: InputField(
                  mx: 5, my: 5,
                  hint: "Server Port",
                  controller: portController,
                  radius: 12,
                  onChanged: (_) {
                    widget.gameConfig.port =
                        int.tryParse(portController.text) ?? widget.gameConfig.port;
                  },
                ),
              )
            ],
          ),
        ),
      )
    );
  }
}