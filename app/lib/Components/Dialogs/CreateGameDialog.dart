import 'package:flutter/material.dart';
import 'package:orbstrike/Components/CustomWidgets/InputField.dart';

class CreateGameDialog extends StatelessWidget {
  final TextEditingController radiusController;
  final TextEditingController plRadiusController;
  final TextEditingController plRingRadiusController;
  final TextEditingController speedController;
  final TextEditingController maxPlayerController;

  const CreateGameDialog({Key? key,
    required this.radiusController,
    required this.maxPlayerController,
    required this.plRadiusController,
    required this.plRingRadiusController,
    required this.speedController
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      title: const Text(
        "Create Game",
        style: TextStyle(color: Colors.white60),
        textAlign: TextAlign.center,
      ),
      insetPadding: const EdgeInsets.all(12),
      backgroundColor: const Color.fromRGBO(46,49,54, 1),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          InputField(
            controller: radiusController,
            hint: "Map Size",
            keyboardType: TextInputType.number,
            my: 5,
            radius: 12,
          ),
          InputField(
            controller: plRadiusController,
            hint: "Player Size",
            keyboardType: TextInputType.number,
            my: 5,
            radius: 12,
          ),
          InputField(
            controller: plRingRadiusController,
            hint: "Player Ring Size",
            keyboardType: TextInputType.number,
            my: 5,
            radius: 12,
          ),
          InputField(
            controller: speedController,
            hint: "Speed",
            keyboardType: TextInputType.number,
            my: 5,
            radius: 12,
          ),
          InputField(
            controller: maxPlayerController,
            hint: "Max Players",
            keyboardType: TextInputType.number,
            my: 5,
            radius: 12,
          ),
        ],
      ),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.pop(context, "Cancel");
          },
          child: const Text("Cancel"),
        ),
        TextButton(
          onPressed: () {
            Navigator.pop(context, "Create");
          },
          child: const Text("Create"),
        )
      ],
    );
  }
}