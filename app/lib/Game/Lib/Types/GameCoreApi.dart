import 'package:grpc/grpc.dart';
import 'package:rxdart/rxdart.dart';

import 'package:orbstrike/proto/game/game.pbgrpc.dart';

class GameCoreApi {
  int fetchCount;

  ClientChannel? chan;
  GameServiceClient? client;
  BehaviorSubject<Move>? stream;
  ChannelCredentials? credentials;

  GameCoreApi({this.stream, this.chan, this.client, this.credentials, this.fetchCount=-1});
}