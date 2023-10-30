import 'dart:ui' as ui;

import 'package:flame/collisions.dart';
import 'package:flame/components.dart';
import 'package:flame/sprite.dart';
import 'package:flutter/material.dart';
import 'package:orbstrike/Game/Player.helper.dart';
import 'package:orbstrike/proto/game/game.pb.dart';

class PlayerComponent extends PositionComponent with CollisionCallbacks {
  Player networkPlayerRep;
  final SpriteSheet playerNormalSS;
  final SpriteSheet playerRingedSS;
  final SpriteSheet ringAnimateSS;

  ui.Image? playerNormalRendered;
  ui.Image? playerRingedRendered;
  SpriteAnimationComponent? ringComponent;

  bool ringEnabledBuf = false;

  static const RINGANIMATEINTERVAL = 0.05;

  late ShapeHitbox hitbox;

  final TextPainter nameTextPainter = TextPainter();

  PlayerComponent({
    required this.networkPlayerRep,
    required this.playerNormalSS,
    required this.playerRingedSS,
    required this.ringAnimateSS
  });

  @override
  void onLoad() async {
    super.onLoad();
    hitbox = CircleHitbox(
        radius: networkPlayerRep.rad,
        position: Vector2.all(0),
        anchor: Anchor.center
    )..renderShape = false;
    add(hitbox);

    playerNormalRendered = await prerenderSpriteComponent(
      playerNormalSS.getSprite(0, networkPlayerRep.color),
      Size.fromRadius(networkPlayerRep.rad),
    );

    playerRingedRendered = await prerenderSpriteComponent(
      playerRingedSS.getSprite(0, networkPlayerRep.color),
      Size.fromRadius(networkPlayerRep.rad),
    );

    ringComponent = SpriteAnimationComponent(
      animation: ringAnimateSS.createAnimation(
          row: 0, from: 0, to: ringAnimateSS.columns-1,
          stepTime: RINGANIMATEINTERVAL
      ),
      size: Vector2.all(networkPlayerRep.ringrad*2),
      anchor: Anchor.center
    );
  }

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    if (playerNormalRendered!=null && playerRingedRendered!=null) {
      canvas.drawImage(
        networkPlayerRep.ringEnabled ? playerRingedRendered! : playerNormalRendered!,
        Offset(-networkPlayerRep.rad, -networkPlayerRep.rad),
        Paint()
      );
    }

    nameTextPainter
      ..text = TextSpan(
        text: "${networkPlayerRep.name} 💀 ${networkPlayerRep.kills}",
        style: const TextStyle(color: Colors.white60, fontSize: 12),
      )
      ..textDirection = TextDirection.ltr
      ..textAlign = TextAlign.center
      ..layout()
      ..paint(canvas, Offset(-(nameTextPainter.width / 2), networkPlayerRep.rad+10));
  }

  @override
  void update(double dt) {
    super.update(dt);

    if (ringComponent!=null && networkPlayerRep.ringEnabled && !ringEnabledBuf) {
      add(ringComponent!);
      ringEnabledBuf=networkPlayerRep.ringEnabled;
    } else if (ringComponent!=null && !networkPlayerRep.ringEnabled && ringEnabledBuf) {
      remove(ringComponent!);
      ringEnabledBuf=networkPlayerRep.ringEnabled;
    }

    if (x!=networkPlayerRep.x || y!=networkPlayerRep.y) {
      x = networkPlayerRep.x;
      y = networkPlayerRep.y;
    }
  }
}

class MainPlayerComponent extends PlayerComponent {
  final List<int> collided;

  MainPlayerComponent({
    required super.networkPlayerRep,
    required super.playerNormalSS,
    required super.playerRingedSS,
    required super.ringAnimateSS,
    required this.collided
  });

  @override
  void onCollisionStart(Set<Vector2> intersectionPoints, PositionComponent other) {
    super.onCollisionStart(intersectionPoints, other);
    if (other is EnemyPlayerComponent) {
      collided.add(other.networkPlayerRep.id);
    }
  }

  @override
  void onCollisionEnd(PositionComponent other) {
    super.onCollisionEnd(other);
    if (other is EnemyPlayerComponent) {
      collided.remove(other.networkPlayerRep.id);
    }
  }
}

class EnemyPlayerComponent extends PlayerComponent {
  EnemyPlayerComponent({
    required super.networkPlayerRep,
    required super.playerNormalSS,
    required super.playerRingedSS,
    required super.ringAnimateSS
  });
}