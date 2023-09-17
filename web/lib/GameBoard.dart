import 'dart:async';

import 'package:flutter/material.dart';

import 'package:flame/game.dart';
import 'package:flame/components.dart';
import 'package:grpc/grpc_web.dart';
import 'proto/game.pbgrpc.dart';

// This will be in cookie or something like that
const myid = 32523;

class GameField extends FlameGame {
  final Uri apiUri;
  final World world = World();
  final Map<int, PlayerC> playerComponents = {};
  late final CameraComponent mainCamera;

  static GrpcWebClientChannel? apiChan;
  static GameServiceClient? apiClient;
  static StreamController<Move> apiReqStream = StreamController<Move>();

  // Border size is loaded when connected to the gRPC endpoint
  static GameBoard board = GameBoard();
  static WorldBorder? border;
  static PlayerC mainPlayerComponent = PlayerC(pPlayer: Player());

  GameField({required this.apiUri});

  void _updateGameBoard(GameBoard board, Map<int, PlayerC> playerComps) {
    final playerIDs = board.players.map((p) => p.id).toSet();

    // Remove players that are not existent anymore
    playerComps.removeWhere((id, component) {
      if (!playerIDs.contains(id)) {
        world.remove(component);
        return true;
      }
      return false;
    });

    for (var pPlayer in board.players) {
      if (!playerComps.containsKey(pPlayer.id)) {
        final player = PlayerC(pPlayer: pPlayer);
        if (pPlayer.id==myid) {
          mainPlayerComponent = player;
        } else {
          playerComponents[pPlayer.id] = player;
        }
        world.add(player);
      }
    }
  }

  @override
  Color backgroundColor() => Colors.blue;

  @override
  void render(Canvas canvas) {
    super.render(canvas);
  }

  @override
  void update(double dt) {
    super.update(dt);

    mainCamera.follow(mainPlayerComponent);
  }

  @override
  Future<void> onLoad() async {
    mainCamera = CameraComponent(world: world);
    addAll([mainCamera, world]);

    apiChan = GrpcWebClientChannel.xhr(apiUri);
    apiClient = GameServiceClient(apiChan!);
    apiReqStream = StreamController<Move>();

    apiClient!.streamGameboard(apiReqStream.stream)
        .listen((game) {
      board = game;
      _updateGameBoard(board, playerComponents);
      if (border==null) {
        border = WorldBorder(color: Colors.green, width: board.y, height: board.x, stroke: 10);
        add(border!);
      }
    });
    apiReqStream.add(Move(direction: Move_Direction.RIGHT));
  }
}

class PlayerC extends PositionComponent with HasGameRef<GameField> {
  // This variable is a pointer to the player object
  final Player pPlayer;

  PlayerC({required this.pPlayer});

  @override
  void render(Canvas canvas) {
    super.render(canvas);

    x = pPlayer.x;
    y = pPlayer.y;

    final paint = Paint()..color = Colors.red;
    canvas.drawCircle(const Offset(0, 0), 15, paint);

    if (pPlayer.ringEnabled) {
      final ringpnt = Paint()
          ..color = const Color(0xFF000000)
          ..style = PaintingStyle.stroke
          ..strokeWidth = 3;
      canvas.drawCircle(const Offset(0,0), 20, ringpnt);
    }
  }

  @override
  void update(double dt) {
    super.update(dt);
  }
}

class WorldBorder extends PositionComponent {
  final Color color;
  final double stroke;
  final double width;
  final double height;

  WorldBorder({required this.color, required this.width, required this.height, required this.stroke});

  @override
  double get z => 1000;

  @override
  void render(Canvas canvas) {
    final paint = Paint()
        ..color = color
        ..style = PaintingStyle.stroke
        ..strokeWidth = stroke;

    final rect = Rect.fromLTWH(0,0, width, height);
    canvas.drawRect(rect, paint);
  }

  @override
  bool get isHud => false;
}

