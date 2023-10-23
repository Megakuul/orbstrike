import 'dart:ui';

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/material.dart';
import 'package:orbstrike/proto/game/game.pb.dart';

class MainPlayer extends PositionComponent with CollisionCallbacks {
  Player networkPlayerRep;
  final List<int> collided;
  final SpriteSheet spriteSheet;
  late ShapeHitbox hitbox;

  Sprite? playerSprite;

  MainPlayer({required this.networkPlayerRep, required this.collided, required this.spriteSheet});

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
  void onLoad() async {
    hitbox = CircleHitbox(
        radius: networkPlayerRep.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);

    playerSprite = spriteSheet.getSprite(0, networkPlayerRep.color);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    if (playerSprite!=null) {
      playerSprite?.render(
        canvas,
        position: Vector2(-networkPlayerRep.rad, -networkPlayerRep.rad),
        size: Vector2(networkPlayerRep.rad*2, networkPlayerRep.rad*2),
      );
    }

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
  final SpriteSheet spriteSheet;
  late ShapeHitbox hitbox;

  Sprite? playerSprite;

  final TextPainter textPainter = TextPainter();

  EnemyPlayer({required this.networkPlayerRep, required this.spriteSheet});

  @override
  void onLoad() {
    hitbox = CircleHitbox(
        radius: networkPlayerRep.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);

    playerSprite = spriteSheet.getSprite(0, networkPlayerRep.color);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    if (playerSprite!=null) {
      playerSprite?.render(
        canvas,
        position: Vector2(-networkPlayerRep.rad, -networkPlayerRep.rad),
        size: Vector2(networkPlayerRep.rad*2, networkPlayerRep.rad*2),
      );
    }

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