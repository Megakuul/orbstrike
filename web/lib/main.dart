import 'dart:async';

import 'package:flutter/material.dart';
import 'package:grpc/grpc_or_grpcweb.dart';
import 'package:grpc/grpc_web.dart';
import 'package:grpc/grpc.dart';
import 'proto/game.pbgrpc.dart';
import 'proto/game.pb.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'orbstrike',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});
  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int kills = 0;

  void StartChan() {
    final chan = GrpcWebClientChannel.xhr(Uri.parse("http://localhost:8080"));

    final client = GameServiceClient(chan);

    final request = StreamController<Move>();

    final response = client.streamGameboard(request.stream);

    response.listen((gameboard) {
      print(gameboard.players[1].kills);
      kills = gameboard.players[1].kills;
    });

    request.add(Move(direction: Move_Direction.RIGHT));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        alignment: Alignment.center,
        decoration: BoxDecoration(
          color: Colors.black54,
          border: Border.all(width: 10, color: Colors.orange)
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text("Meister;"),
            TextButton(
                onPressed: () {
                  setState(() {
                    StartChan();
                  });
                },
                child: const Text("Clyck me")
            ),
            Text("Hallo $kills")
          ],
        ),
      )
    );
  }
}
