import 'package:flame/components.dart';
import 'package:flutter/material.dart';


// TODO: Do something else, its absolutely ugly and not what i expected + its in foreground
class WorldBackground extends PositionComponent {
  final double radius;
  static const offset = Offset(0, 0);
  Paint paint = Paint()
      ..color=Colors.orange;

  WorldBackground({required this.radius});

  @override
  double get z => -1000;

  @override
  bool get isHud => false;

  @override
  void render(Canvas canvas) {
    const double rectSize = 20;  // Adjust size as needed
    const double borderRadius = 5;

    for (double x = -radius; x <= radius; x += rectSize) {
      for (double y = -radius; y <= radius; y += rectSize) {
        if ((Offset(x, y) - offset).distance <= radius) {
          final rect = RRect.fromLTRBR(
              x, y, x + rectSize, y + rectSize,
              const Radius.circular(borderRadius)
          );
          canvas.drawRRect(rect, paint);
        }
      }
    }
  }
}
