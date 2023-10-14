import 'package:flutter/material.dart';

class InputField extends StatelessWidget {

  final double radius;
  final double my;
  final double mx;
  final String hint;
  final TextEditingController controller;
  final void Function(String)? onChanged;
  final TextInputType? keyboardType;

  const InputField({
    Key? key,
    required this.radius,
    required this.controller,
    required this.hint,
    this.onChanged,
    this.keyboardType,
    this.my = 0,
    this.mx = 0,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.symmetric(horizontal: mx, vertical: my),
      padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 2),
      decoration: BoxDecoration(
        color: const Color.fromRGBO(74,76,81, 1),
        borderRadius: BorderRadius.all(Radius.circular(radius)),
      ),
      child: TextFormField(
        keyboardType: keyboardType,
        cursorColor: Colors.white60,
        style: const TextStyle(color: Colors.white60),
        decoration: InputDecoration(
          labelText: hint,
          labelStyle: const TextStyle(color: Colors.white30),
          border: InputBorder.none,
        ),
        controller: controller,
        onChanged: onChanged,
      ),
    );
  }
}
