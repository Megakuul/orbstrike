import 'package:flame/components.dart';
import 'package:orbstrike/proto/game.pb.dart';
import 'package:flutter/material.dart';
import 'proto/game.pbgrpc.dart';
import 'package:orbstrike/GameBoard.dart';

class PlayerC extends PositionComponent with HasGameRef<GameField> {
  // This variable is a pointer to the player object
  final Player pPlayer;

  PlayerC({required this.pPlayer});

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    final paint = Paint()..color = Colors.red;
    canvas.drawCircle(const Offset(0, 0), 40, paint);

    if (pPlayer.ringEnabled) {
      final ringpnt = Paint()
        ..color = const Color(0xFF000000)
        ..style = PaintingStyle.stroke
        ..strokeWidth = 3;
      canvas.drawCircle(const Offset(0,0), 50, ringpnt);
    }
  }

  @override
  void update(double dt) {
    super.update(dt);

    if (x!=pPlayer.x || y!=pPlayer.y) {
      x = pPlayer.x;
      y = pPlayer.y;
    }
  }

}