import 'package:flutter/material.dart';

class ElevatedIconButton extends StatelessWidget {
  final Text text;
  final Icon icon;
  final Gradient gradient;
  final void Function() onPressed;
  final double my;
  final double mx;

  const ElevatedIconButton({Key? key,
    required this.text,
    required this.gradient,
    required this.icon,
    required this.onPressed,
    this.my = 0,
    this.mx = 0,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.symmetric(vertical: mx, horizontal: my),
      padding: const EdgeInsets.all(10),
      decoration: BoxDecoration(
        gradient: gradient,
        borderRadius: BorderRadius.circular(12)
      ),
      child: ElevatedButton(
        onPressed: onPressed,
        style: ElevatedButton.styleFrom(
          foregroundColor: Colors.white60,
          backgroundColor: Colors.transparent,
          shadowColor: Colors.transparent,
        ),
        child: Row(
          mainAxisSize: MainAxisSize.min,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [text, icon],
        )
      ),
    );
  }
}