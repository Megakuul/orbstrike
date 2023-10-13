import 'package:flame/components.dart';
import 'package:flutter/material.dart';

class WorldBorder extends PositionComponent {
  final List<Color> colors;
  final double stroke;
  final double radius;
  late final Paint paint;
  static const offset = Offset(0,0);
  static const strokew = 50.0;

  WorldBorder({required this.colors, required this.radius, required this.stroke}) {
    final rect = Rect.fromCircle(center: offset, radius: radius);
    paint = Paint()
      ..style = PaintingStyle.stroke
      ..strokeWidth = strokew
      ..shader = LinearGradient(
        colors: colors,
        begin: Alignment.topLeft,
        end: Alignment.bottomRight
      ).createShader(rect);
  }

  @override
  double get z => 1000;

  @override
  void render(Canvas canvas) {
    canvas.drawCircle(offset, radius, paint);
  }

  @override
  bool get isHud => false;
}