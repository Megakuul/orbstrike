import 'dart:async';
import 'dart:ui';

import 'package:flame/components.dart';
import 'package:image_scaler/image_scaler.dart';
import 'package:image_scaler/types.dart';

/// Prerenders a Sprite to Image
///
/// Downscaling Sprites commes with huge quality impact
/// this function converts the sprite to a image and scales it down.
///
/// It gives a high quality image, but make sure to not run this function in the render-loop
/// it is really slow compared to native sprite downscaling as it uses a pictureRecorder.
Future<Image> prerenderSpriteComponent(Sprite sprite, Size size) async {
  final spriteComp = SpriteComponent(sprite: sprite);

  // Original Image
  final oiRecorder = PictureRecorder();
  final oiCanvas = Canvas(oiRecorder);
  spriteComp.render(oiCanvas);

  final oiImage = await oiRecorder.endRecording().toImage(
    sprite.srcSize.x.toInt(),
    sprite.srcSize.y.toInt(),
  );

  return await scale(
    image: oiImage,
    newSize: IntSize(size.width.toInt(), size.height.toInt()),
    algorithm: ScaleAlgorithm.lanczos,
    areaRadius: 3,
  );
}