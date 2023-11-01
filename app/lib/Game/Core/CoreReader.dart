import 'dart:ui';

import 'package:grpc/grpc.dart';
import 'package:orbstrike/Game/Lib/Types/GameCoreApi.dart';
import 'package:orbstrike/Game/Lib/Types/GameCoreComponents.dart';
import 'package:orbstrike/Game/Lib/Types/MainUICallbacks.dart';
import 'package:orbstrike/Game/Core/CoreUpdater.dart';
import 'package:orbstrike/proto/game/game.pbgrpc.dart';

import 'package:rxdart/rxdart.dart';



void startCoreReader(GameCoreApi api, GameCoreComponents coreComponents, MainUICallbacks mCallbacks, int gameId, String host, int port) {
  mCallbacks.setProgress(false, "");
  int reconnectTimeout = 5;

  if (api.chan != null) {
    api.chan!.shutdown();
  }

  if (api.stream != null) {
    api.stream!.close();
  }

  api.stream = BehaviorSubject<Move>();

  api.chan = ClientChannel(
    host,
    port: port,
    options: ChannelOptions(
      credentials: api.credentials ?? const ChannelCredentials.insecure(),
    ),
  );

  api.client = GameServiceClient(api.chan!);

  final Map<String, String> streamHeaders = {
    "gameid": "$gameId",
  };

  final callOptions = CallOptions(metadata: streamHeaders);

  api.client!.proxyGameboard(api.stream!.stream, options: callOptions).listen((game) {
    coreComponents.board = game;

    updateGameBoard(coreComponents).forEach((key, value) {
      if (value) {
        coreComponents.world.add(key);
      } else {
        coreComponents.world.remove(key);
      }
    });

    if (api.fetchCount!=-1) { api.fetchCount++; }
  }, onError: (error) {
    if (error is! GrpcError) {
      mCallbacks.showDial(error.message, const Color.fromRGBO(66,69,73, 1));
      return;
    }

    switch (error.code) {
      case StatusCode.unknown:
      case StatusCode.unavailable:
      case StatusCode.internal:
        mCallbacks.setProgress(true, "Reconnecting...", error.message);
        Future.delayed(Duration(seconds: reconnectTimeout), () {
          startCoreReader(api, coreComponents, mCallbacks, gameId, host, port);
        });
        break;
      default:
        mCallbacks.showDial(error.message, const Color.fromRGBO(66,69,73, 1));
        break;
    }
  });
}