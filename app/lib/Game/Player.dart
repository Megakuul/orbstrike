import 'dart:ui';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flutter/material.dart';
import 'package:orbstrike/proto/game/game.pb.dart';

class MainPlayer extends PositionComponent with CollisionCallbacks {
  Player networkPlayerRep;
  final List<int> collided;
  late ShapeHitbox hitbox;

  MainPlayer({required this.networkPlayerRep, required this.collided});

  @override
  void onCollisionStart(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollisionStart(intersectionPoints, other);
    if (other is EnemyPlayer) {
      collided.add(other.networkPlayerRep.id);
    }
  }

  @override
  void onCollisionEnd(PositionComponent other) {
    super.onCollisionEnd(other);
    if (other is EnemyPlayer) {
      collided.remove(other.networkPlayerRep.id);
    }
  }

  @override
  void onLoad() {
    hitbox = CircleHitbox(
        radius: networkPlayerRep.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    final paint = Paint()..color = Colors.red;
    canvas.drawCircle(const Offset(0, 0), networkPlayerRep.rad, paint);

    if (networkPlayerRep.ringEnabled) {
      final ringpnt = Paint()
        ..color = const Color(0xFF000000)
        ..style = PaintingStyle.stroke
        ..strokeWidth = 3;
      canvas.drawCircle(const Offset(0,0), networkPlayerRep.ringrad, ringpnt);
    }
  }

  @override
  void update(double dt) {
    super.update(dt);

    if (x!=networkPlayerRep.x || y!=networkPlayerRep.y) {
      x = networkPlayerRep.x;
      y = networkPlayerRep.y;
    }
  }
}

class EnemyPlayer extends PositionComponent {
  Player networkPlayerRep;
  late ShapeHitbox hitbox;

  final TextPainter textPainter = TextPainter();

  EnemyPlayer({required this.networkPlayerRep});

  @override
  void onLoad() {
    hitbox = CircleHitbox(
        radius: networkPlayerRep.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    final paint = Paint()..color = Colors.red;

    canvas.drawCircle(const Offset(0, 0), networkPlayerRep.rad, paint);

    if (networkPlayerRep.ringEnabled) {
      final ringpnt = Paint()
        ..color = const Color(0xFF000000)
        ..style = PaintingStyle.stroke
        ..strokeWidth = 3;
      canvas.drawCircle(const Offset(0,0), networkPlayerRep.ringrad, ringpnt);
    }

    // Draw text
    textPainter
      ..text = TextSpan(
        text: networkPlayerRep.name,
        style: const TextStyle(color: Colors.white, fontSize: 12),
      )
      ..textDirection = TextDirection.ltr
      ..textAlign = TextAlign.center
      ..layout()
      ..paint(canvas, Offset(-(textPainter.width / 2), networkPlayerRep.rad+10));
  }

  @override
  void update(double dt) {
    super.update(dt);

    if (x!=networkPlayerRep.x || y!=networkPlayerRep.y) {
      x = networkPlayerRep.x;
      y = networkPlayerRep.y;
    }
  }

}