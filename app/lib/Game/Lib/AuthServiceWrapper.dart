import 'package:grpc/grpc.dart';

import 'package:orbstrike/proto/auth/auth.pbgrpc.dart';
import 'package:orbstrike/Game/Lib/Types/UserCredentials.dart';

Future<UserCredentials> joinGame(int gameId, String name, String host, int port, ChannelCredentials? credentials) async {
  final client = AuthServiceClient(
      ClientChannel(
          host,
          port: port,
          options: ChannelOptions(credentials: credentials ?? const ChannelCredentials.insecure())
      )
  );

  final req = JoinGameRequest(
    gameid: gameId,
    name:  name,
  );
  final res = await client.joinGame(req);
  return UserCredentials(
    id: res.userid,
    key: res.userkey,
  );
}

Future<void> exitGame(int gameId, UserCredentials? userCreds, String host, int port, ChannelCredentials? credentials) async {
  final client = AuthServiceClient(
      ClientChannel(
          host,
          port: port,
          options: ChannelOptions(credentials: credentials ?? const ChannelCredentials.insecure())
      )
  );

  final req = ExitGameRequest(
      gameid: gameId,
      userkey: userCreds?.key
  );
  await client.exitGame(req);
}