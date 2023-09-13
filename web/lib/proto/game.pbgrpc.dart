//
//  Generated code. Do not modify.
//  source: game.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'game.pb.dart' as $0;

export 'game.pb.dart';

@$pb.GrpcServiceName('game.GameService')
class GameServiceClient extends $grpc.Client {
  static final _$gameUpdate = $grpc.ClientMethod<$0.Move, $0.GameBoard>(
      '/game.GameService/GameUpdate',
      ($0.Move value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.GameBoard.fromBuffer(value));

  GameServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.GameBoard> gameUpdate($0.Move request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$gameUpdate, request, options: options);
  }
}

@$pb.GrpcServiceName('game.GameService')
abstract class GameServiceBase extends $grpc.Service {
  $core.String get $name => 'game.GameService';

  GameServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Move, $0.GameBoard>(
        'GameUpdate',
        gameUpdate_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Move.fromBuffer(value),
        ($0.GameBoard value) => value.writeToBuffer()));
  }

  $async.Future<$0.GameBoard> gameUpdate_Pre($grpc.ServiceCall call, $async.Future<$0.Move> request) async {
    return gameUpdate(call, await request);
  }

  $async.Future<$0.GameBoard> gameUpdate($grpc.ServiceCall call, $0.Move request);
}
