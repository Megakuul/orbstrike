import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'GameBoard.dart';

class GameConfiguration {
  GameConfiguration({required this.hostname, required this.port, required this.gameID});

  String? hostname;
  int? port;
  int? gameID;
}

GameConfiguration gameConfig = GameConfiguration(
    hostname: "localhost",
    port: 8080,
    gameID: 0
);

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
      home: const MainPage(title: "Orbstrike"),
    );
  }
}

class MainPage extends StatefulWidget {
  final String title;
  const MainPage({Key? key, required this.title}) : super(key: key);

  @override
  State<MainPage> createState() => _MainPage();
}

class _MainPage extends State<MainPage> {
  Widget curPage = Home();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: curPage,
      floatingActionButton: FloatingActionButton(
        onPressed: () {

        },
        child: const Icon(Icons.ac_unit_sharp),
      ),
    );
  }
}

class Home extends StatefulWidget {
  const Home({Key? key}) : super(key: key);

  @override
  State<Home> createState() => _Home();
}

class _Home extends State<Home> {
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        TextButton(
            onPressed: onPressed,
            child: child
        ),
        TextButton(
            onPressed: onPressed,
            child: child
        ),
        TextField(

        ),
      ],
    );
  }
}
