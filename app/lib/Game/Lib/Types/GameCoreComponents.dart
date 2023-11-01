import 'package:flame/camera.dart';
import 'package:flame/sprite.dart';

import 'package:orbstrike/proto/game/game.pb.dart';
import 'package:orbstrike/Game/Player/Player.dart';
import 'package:orbstrike/Game/World/WorldBackground.dart';
import 'package:orbstrike/Game/World/WorldBorder.dart';
import 'UserCredentials.dart';

class GameCoreComponents {
  late final CameraComponent mainCamera;
  late final double lerpFactor;

  GameBoard board;
  final World world = World();
  WorldBorder? border;
  WorldBackground? background;

  MainPlayerComponent? mainPlayerComponent;
  Move_Direction mainPlayerDirection;
  UserCredentials? mainPlayerCreds;
  bool mainPlayerRingState;

  final List<int> mainPlayerCollided = [];
  final Map<int, EnemyPlayerComponent> playerComponents = {};

  late SpriteSheet playerNameSS;
  late SpriteSheet playerRingedSS;
  late SpriteSheet ringAnimateSS;

  GameCoreComponents({
    required this.board,
    required this.mainPlayerDirection,
    required this.mainPlayerRingState,
    this.mainPlayerCreds,
  });
}