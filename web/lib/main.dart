import 'dart:async';

import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc_web.dart';
import 'proto/game.pbgrpc.dart';
import 'proto/game.pb.dart';
import 'GameBoard.dart';


void main() {
  final game = GameField(apiUri: Uri.parse("http://localhost:8080"));

  runApp(GameWidget(game: game));
  // runApp(const MyApp());
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
