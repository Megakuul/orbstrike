import 'dart:ui';

class MainUICallbacks {
  final void Function(String? message, Color color) showDial;
  final void Function(String? message, Color color) showSnack;
  final void Function(bool loading, String message, [String? detailedMessage]) setProgress;

  MainUICallbacks({required this.showDial, required this.showSnack, required this.setProgress});
}