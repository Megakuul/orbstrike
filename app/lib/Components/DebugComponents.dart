import 'package:flame/components.dart';
import 'package:orbstrike/Game/GameBoard.helper.dart';

class NetworkFrameCounter extends TextComponent {
  final GameCoreApi coreApi;
  final timer = Stopwatch()..start();

  NetworkFrameCounter({required this.coreApi}) : super() {
    priority = 999;
    coreApi.fetchCount = 0;
  }

  @override
  void update(double dt) {
    super.update(dt);
    if (timer.elapsedMilliseconds >= 1000) {
      text = "Network FPS: ${coreApi.fetchCount}";
      coreApi.fetchCount = 0;
      timer.reset();
    }
  }
}

class IngameFrameCounter extends FpsTextComponent {
  @override
  set text(String text) {
    super.text = "Ingame FPS: $text";
  }
}