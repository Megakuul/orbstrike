import 'dart:ui' as ui;
import 'package:flame/components.dart';
import 'package:flutter/material.dart';

class WorldBackground extends PositionComponent {
  final double radius;
  final double rectSize;
  final double rectSpacing;
  final double rectBorderRadius;
  final Paint rectPaint;

  ui.Picture? bgPicture;
  ui.Image? bgImage;
  ui.PictureRecorder bgRecorder = ui.PictureRecorder();

  WorldBackground({required this.radius, required this.rectSize, required this.rectSpacing, required this.rectBorderRadius, required this.rectPaint}) {
    bgPicture = prerenderBackground(
      bgRecorder,
      radius,
      rectSize,
      rectSpacing,
      rectBorderRadius
    );
  }

  /// Prerender Background matrix
  ///
  /// Background matrix is prerendered, as it is static,
  /// we don't want it to be calculated on every render frame of the screen
  ui.Picture prerenderBackground(ui.PictureRecorder rec, double radius, double rectSize, double spacing, double borderRadius) {
    final offscreenCanvas = Canvas(rec);
    final size = radius*2;

    for (double x = 0; x <= size; x += rectSize+spacing) {
      for (double y = 0; y <= size; y += rectSize+spacing) {
        if ((Offset(x, y) - Offset(radius, radius)).distance <= radius) {
          final rect = RRect.fromLTRBR(
              x, y, x + rectSize, y + rectSize,
              Radius.circular(borderRadius)
          );
          offscreenCanvas.drawRRect(rect, rectPaint);
        }
      }
    }
    return rec.endRecording();
  }

  @override
  int get priority => -1000;

  @override
  void render(Canvas canvas) async {
    if (bgImage==null) {
      final size = (radius*2).toInt() + 1;
      bgImage = await bgPicture?.toImage(size, size);
    } else {
      canvas.drawImage(bgImage!, Offset(-radius, -radius), Paint());
    }
  }
}
