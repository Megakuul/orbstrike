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
    return MaterialApp(
      title: "Orbstrike",
      theme: ThemeData(
        primaryColor: createCustomSwatch(const Color.fromARGB(255, 20, 30, 48)),
      ),
      home: MainPage(title: "Orbstrike", gameConfig: GameConfiguration(
        hostname: "",
        port: 54331,
        gameID: 0,
        latestGameIDs: []
      )),
    );
  }
}
