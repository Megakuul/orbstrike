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
        style: TextStyle(color: Colors.white60, fontSize: 15),
        textAlign: TextAlign.center,
      ),
      insetPadding: const EdgeInsets.all(12),
      backgroundColor: color,
      content: Padding(
        padding: const EdgeInsets.all(30),
        child: Text(
          message,
          style: TextStyle(
            color: Colors.white60,
            fontSize: message.length<20 ? 40 : 15,
            fontWeight: FontWeight.w600
          ),
          textAlign: TextAlign.center,
        ),
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.pop(context);
            Navigator.pop(context);
          },
          child: const Text("Exit"),
        ),
      ],
    );
  }
}