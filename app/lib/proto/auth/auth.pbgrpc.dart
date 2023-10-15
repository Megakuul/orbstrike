//
//  Generated code. Do not modify.
//  source: auth.proto
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

import 'auth.pb.dart' as $0;

export 'auth.pb.dart';

@$pb.GrpcServiceName('auth.AuthService')
class AuthServiceClient extends $grpc.Client {
  static final _$createGame = $grpc.ClientMethod<$0.CreateGameRequest, $0.CreateGameResponse>(
      '/auth.AuthService/CreateGame',
      ($0.CreateGameRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.CreateGameResponse.fromBuffer(value));
  static final _$joinGame = $grpc.ClientMethod<$0.JoinGameRequest, $0.JoinGameResponse>(
      '/auth.AuthService/JoinGame',
      ($0.JoinGameRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.JoinGameResponse.fromBuffer(value));
  static final _$exitGame = $grpc.ClientMethod<$0.ExitGameRequest, $0.ExitGameResponse>(
      '/auth.AuthService/ExitGame',
      ($0.ExitGameRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.ExitGameResponse.fromBuffer(value));

  AuthServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.CreateGameResponse> createGame($0.CreateGameRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createGame, request, options: options);
  }

  $grpc.ResponseFuture<$0.JoinGameResponse> joinGame($0.JoinGameRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$joinGame, request, options: options);
  }

  $grpc.ResponseFuture<$0.ExitGameResponse> exitGame($0.ExitGameRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$exitGame, request, options: options);
  }
}

@$pb.GrpcServiceName('auth.AuthService')
abstract class AuthServiceBase extends $grpc.Service {
  $core.String get $name => 'auth.AuthService';

  AuthServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CreateGameRequest, $0.CreateGameResponse>(
        'CreateGame',
        createGame_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateGameRequest.fromBuffer(value),
        ($0.CreateGameResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.JoinGameRequest, $0.JoinGameResponse>(
        'JoinGame',
        joinGame_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.JoinGameRequest.fromBuffer(value),
        ($0.JoinGameResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ExitGameRequest, $0.ExitGameResponse>(
        'ExitGame',
        exitGame_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ExitGameRequest.fromBuffer(value),
        ($0.ExitGameResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CreateGameResponse> createGame_Pre($grpc.ServiceCall call, $async.Future<$0.CreateGameRequest> request) async {
    return createGame(call, await request);
  }

  $async.Future<$0.JoinGameResponse> joinGame_Pre($grpc.ServiceCall call, $async.Future<$0.JoinGameRequest> request) async {
    return joinGame(call, await request);
  }

  $async.Future<$0.ExitGameResponse> exitGame_Pre($grpc.ServiceCall call, $async.Future<$0.ExitGameRequest> request) async {
    return exitGame(call, await request);
  }

  $async.Future<$0.CreateGameResponse> createGame($grpc.ServiceCall call, $0.CreateGameRequest request);
  $async.Future<$0.JoinGameResponse> joinGame($grpc.ServiceCall call, $0.JoinGameRequest request);
  $async.Future<$0.ExitGameResponse> exitGame($grpc.ServiceCall call, $0.ExitGameRequest request);
}
