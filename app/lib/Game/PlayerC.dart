import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flutter/material.dart';
import 'package:orbstrike/proto/game/game.pb.dart';

/// Current player
class PlayerC extends PositionComponent with CollisionCallbacks {
  // This variable is a pointer to the player object
  final Player pPlayer;
  final List<int> collided;
  late ShapeHitbox hitbox;

  PlayerC({required this.pPlayer, required this.collided});

  @override
  void onCollisionStart(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollisionStart(intersectionPoints, other);
    if (other is PlayerO) {
      collided.add(other.pPlayer.id);
    }
  }

  @override
  void onCollisionEnd(PositionComponent other) {
    super.onCollisionEnd(other);
    if (other is PlayerO) {
      collided.remove(other.pPlayer.id);
    }
  }

  @override
  void onLoad() {
    hitbox = CircleHitbox(
        radius: pPlayer.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    final paint = Paint()..color = Colors.red;
    canvas.drawCircle(const Offset(0, 0), pPlayer.rad, paint);

    if (pPlayer.ringEnabled) {
      final ringpnt = Paint()
        ..color = const Color(0xFF000000)
        ..style = PaintingStyle.stroke
        ..strokeWidth = 3;
      canvas.drawCircle(const Offset(0,0), pPlayer.ringrad, ringpnt);
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

/// Other Players
class PlayerO extends PositionComponent {
  // This variable is a pointer to the player object
  final Player pPlayer;
  late ShapeHitbox hitbox;

  PlayerO({required this.pPlayer});

  @override
  void onLoad() {
    hitbox = CircleHitbox(
        radius: pPlayer.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    final paint = Paint()..color = Colors.red;
    canvas.drawCircle(const Offset(0, 0), pPlayer.rad, paint);

    if (pPlayer.ringEnabled) {
      final ringpnt = Paint()
        ..color = const Color(0xFF000000)
        ..style = PaintingStyle.stroke
        ..strokeWidth = 3;
      canvas.drawCircle(const Offset(0,0), pPlayer.ringrad, ringpnt);
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