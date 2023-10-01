import 'package:flutter/services.dart';
import 'package:orbstrike/proto/game.pb.dart';

Move_Direction? getMovementInput() {
  final w = RawKeyboard.instance.keysPressed.contains(LogicalKeyboardKey.keyW);
  final s = RawKeyboard.instance.keysPressed.contains(LogicalKeyboardKey.keyS);
  final a = RawKeyboard.instance.keysPressed.contains(LogicalKeyboardKey.keyA);
  final d = RawKeyboard.instance.keysPressed.contains(LogicalKeyboardKey.keyD);
  final space = RawKeyboard.instance.keysPressed.contains(LogicalKeyboardKey.space);

  if (space) { return Move_Direction.NONE; }
  if (w&&d&&s&&a) { return Move_Direction.NONE; }
  else if (w&&s) {
    if (a) { return Move_Direction.LEFT; }
    else if (d) { return Move_Direction.RIGHT; }
    else { return Move_Direction.NONE; }
  } else if (a&&d) {
    if (w) { return Move_Direction.UP; }
    else if (s) { return Move_Direction.DOWN; }
    else { return Move_Direction.NONE; }
  } else if (w&&a) {
    return Move_Direction.UP_LEFT;
  } else if (w&&d) {
    return Move_Direction.UP_RIGHT;
  } else if (s&&a) {
    return Move_Direction.DOWN_LEFT;
  } else if (s&&d) {
    return Move_Direction.DOWN_RIGHT;
  } else if (w) {
    return Move_Direction.UP;
  } else if (s) {
    return Move_Direction.DOWN;
  } else if (a) {
    return Move_Direction.LEFT;
  } else if (d) {
    return Move_Direction.RIGHT;
  }

  return null;
}

bool getRingState() {
  if (RawKeyboard.instance.keysPressed.contains(LogicalKeyboardKey.keyR)) {
    return true;
  }
  return false;
}