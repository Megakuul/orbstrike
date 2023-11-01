import 'package:flutter/material.dart';
import 'package:orbstrike/Components/CustomWidgets/InputField.dart';

class SettingDialog extends StatelessWidget {
  final TextEditingController showDebugController;
  final TextEditingController lerpFactorController;

  const SettingDialog({Key? key,
    required this.showDebugController,
    required this.lerpFactorController,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(16)),
      title: const Text(
        "Advanced Settings",
        style: TextStyle(color: Colors.white60),
        textAlign: TextAlign.center,
      ),
      insetPadding: const EdgeInsets.all(12),
      backgroundColor: const Color.fromRGBO(46,49,54, 1),
      content: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          InputField(
            controller: showDebugController,
            hint: "Show Debuginformation",
            keyboardType: TextInputType.number,
            my: 5,
            radius: 12,
          ),
          InputField(
            controller: lerpFactorController,
            hint: "Lerp Factor",
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
            Navigator.pop(context, "Save");
          },
          child: const Text("Save"),
        )
      ],
    );
  }
}