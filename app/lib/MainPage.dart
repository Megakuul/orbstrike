import 'package:grpc/grpc.dart';
import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:orbstrike/proto/auth/auth.pbgrpc.dart';

import 'Game/GameBoard.dart';

class GameConfiguration {
  GameConfiguration({required this.hostname, required this.port, required this.gameID, required this.latestGameIDs});
  String hostname;
  int port;
  int gameID;
  List<int> latestGameIDs;
}

Future<int> createGame(GameConfiguration conf) async {
  final client = AuthServiceClient(
    ClientChannel(
      conf.hostname,
      port: conf.port,
      options: const ChannelOptions(credentials: ChannelCredentials.insecure())
    )
  );

  final req = CreateGameRequest(identifier: "serialnumber");
  final res = await client.createGame(req);
  return res.gameid;
}

void joinGame(BuildContext context, GameConfiguration conf) {
  Navigator.of(context).push(
      MaterialPageRoute(
          builder: (context) => GameWidget(game: GameField(
            gameID: conf.gameID,
            hudContext: context,
            host: conf.hostname,
            port: conf.port,
          ))
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

  @override
  void initState() {
    super.initState();

    gameIdController.text = widget.gameConfig.gameID.toString();
    hostController.text = widget.gameConfig.hostname;
    portController.text = widget.gameConfig.port.toString();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          TextButton(
              onPressed: () async {
                try {
                  widget.gameConfig.latestGameIDs.add(
                      await createGame(widget.gameConfig)
                  );
                  setState(() {});
                } catch (err) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      backgroundColor: Colors.red,
                      content: Text("ERROR: $err")
                    )
                  );
                }
              },
              child: const Text("Create Game")
          ),
          TextButton(
              onPressed: () {
                widget.gameConfig.gameID = int.tryParse(gameIdController.text) ?? 0;
                joinGame(context, widget.gameConfig);
              },
              child: const Text("Join Game")
          ),
          TextField(
            decoration: const InputDecoration(
              border: OutlineInputBorder(),
              labelText: "Game ID"
            ),
            controller: gameIdController,
          ),

          Row(
            children: [
              Expanded(
                child: TextField(
                  decoration: const InputDecoration(
                      border: OutlineInputBorder(),
                      labelText: "Hostname"
                  ),
                  onSubmitted: (hostname) {
                    widget.gameConfig.hostname = hostname;
                  },
                  controller: hostController,
                ),
              ),
              Expanded(
                  child: TextField(
                    decoration: const InputDecoration(
                        border: OutlineInputBorder(),
                        labelText: "Port"
                    ),
                    onSubmitted: (port) {
                      widget.gameConfig.port = int.tryParse(port) ?? widget.gameConfig.port;
                    },
                    controller: portController,
                  )
              )
            ],
          ),
          Expanded(
              child: ListView.builder(
                  itemCount: widget.gameConfig.latestGameIDs.length,
                  itemBuilder: (context, i) {
                    final id = widget.gameConfig.latestGameIDs[i];
                    return ListTile(
                        title: Text("$id")
                    );
                  }
              )
          )
        ],
      ),
    );
  }
}