import 'package:flutter/material.dart';

import 'MainPage.dart';

void main() {
  runApp(const App());
}

class App extends StatelessWidget {
  const App({Key? key}) : super(key: key);

  MaterialColor createCustomSwatch(Color color) {
    final swatch = <int, Color>{
      50: color.withOpacity(.1),
      100: color.withOpacity(.2),
      200: color.withOpacity(.3),
      300: color.withOpacity(.4),
      400: color.withOpacity(.5),
      500: color.withOpacity(.6),
      600: color.withOpacity(.7),
      700: color.withOpacity(.8),
      800: color.withOpacity(.9),
      900: color.withOpacity(1),
    };
    return MaterialColor(color.value, swatch);
  }

  @override
  Widget build(BuildContext context) {
    const primaryColor = Colors.white60;

    return MaterialApp(
      title: "Orbstrike",
      theme: ThemeData(
        primarySwatch: createCustomSwatch(primaryColor),
      ),
      home: MainPage(title: "Orbstrike", gameConfig: GameConfiguration(
        hostname: "127.0.0.1",
        port: 8080,
        gameID: -1,
        credentials: null,
        lerpFactor: 8,
        showDebug: false,
        latestGameIDs: [],
        uid: (DateTime.now().toUtc().millisecondsSinceEpoch ~/ 1000).toString()
      )),
    );
  }
}
