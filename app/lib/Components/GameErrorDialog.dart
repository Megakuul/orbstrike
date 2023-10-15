import 'package:flutter/material.dart';

class GameErrorDialog extends StatelessWidget {
  final String message;
  final Color color;

  const GameErrorDialog({Key? key,
    required this.message,
    required this.color,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      title: const Text(
        "Game Error",
        style: TextStyle(color: Colors.white60),
        textAlign: TextAlign.center,
      ),
      insetPadding: const EdgeInsets.all(12),
      backgroundColor: const Color.fromRGBO(46,49,54, 1),
      content: Text(message),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.pop(context);
          },
          child: const Text("Exit"),
        ),
      ],
    );
  }
}