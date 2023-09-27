import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'GameBoard.dart';


void main() {
  final game = GameField(host: "localhost", port: 8080);

  runApp(GameWidget(game: game));
}
