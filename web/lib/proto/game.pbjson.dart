//
//  Generated code. Do not modify.
//  source: game.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use gameBoardDescriptor instead')
const GameBoard$json = {
  '1': 'GameBoard',
  '2': [
    {'1': 'players', '3': 1, '4': 3, '5': 11, '6': '.game.GameBoard.PlayersEntry', '10': 'players'},
    {'1': 'id', '3': 2, '4': 1, '5': 5, '10': 'id'},
    {'1': 'rad', '3': 3, '4': 1, '5': 1, '10': 'rad'},
  ],
  '3': [GameBoard_PlayersEntry$json],
};

@$core.Deprecated('Use gameBoardDescriptor instead')
const GameBoard_PlayersEntry$json = {
  '1': 'PlayersEntry',
  '2': [
    {'1': 'key', '3': 1, '4': 1, '5': 5, '10': 'key'},
    {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.game.Player', '10': 'value'},
  ],
  '7': {'7': true},
};

/// Descriptor for `GameBoard`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List gameBoardDescriptor = $convert.base64Decode(
    'CglHYW1lQm9hcmQSNgoHcGxheWVycxgBIAMoCzIcLmdhbWUuR2FtZUJvYXJkLlBsYXllcnNFbn'
    'RyeVIHcGxheWVycxIOCgJpZBgCIAEoBVICaWQSEAoDcmFkGAMgASgBUgNyYWQaSAoMUGxheWVy'
    'c0VudHJ5EhAKA2tleRgBIAEoBVIDa2V5EiIKBXZhbHVlGAIgASgLMgwuZ2FtZS5QbGF5ZXJSBX'
    'ZhbHVlOgI4AQ==');

@$core.Deprecated('Use playerDescriptor instead')
const Player$json = {
  '1': 'Player',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 5, '10': 'id'},
    {'1': 'x', '3': 2, '4': 1, '5': 1, '10': 'x'},
    {'1': 'y', '3': 3, '4': 1, '5': 1, '10': 'y'},
    {'1': 'color', '3': 4, '4': 1, '5': 5, '10': 'color'},
    {'1': 'kills', '3': 5, '4': 1, '5': 5, '10': 'kills'},
    {'1': 'ringEnabled', '3': 6, '4': 1, '5': 8, '10': 'ringEnabled'},
  ],
};

/// Descriptor for `Player`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List playerDescriptor = $convert.base64Decode(
    'CgZQbGF5ZXISDgoCaWQYASABKAVSAmlkEgwKAXgYAiABKAFSAXgSDAoBeRgDIAEoAVIBeRIUCg'
    'Vjb2xvchgEIAEoBVIFY29sb3ISFAoFa2lsbHMYBSABKAVSBWtpbGxzEiAKC3JpbmdFbmFibGVk'
    'GAYgASgIUgtyaW5nRW5hYmxlZA==');

@$core.Deprecated('Use moveDescriptor instead')
const Move$json = {
  '1': 'Move',
  '2': [
    {'1': 'direction', '3': 1, '4': 1, '5': 14, '6': '.game.Move.Direction', '10': 'direction'},
    {'1': 'enableRing', '3': 2, '4': 1, '5': 8, '10': 'enableRing'},
    {'1': 'userkey', '3': 3, '4': 1, '5': 5, '10': 'userkey'},
    {'1': 'gameid', '3': 4, '4': 1, '5': 5, '10': 'gameid'},
  ],
  '4': [Move_Direction$json],
};

@$core.Deprecated('Use moveDescriptor instead')
const Move_Direction$json = {
  '1': 'Direction',
  '2': [
    {'1': 'NONE', '2': 0},
    {'1': 'UP', '2': 1},
    {'1': 'DOWN', '2': 2},
    {'1': 'LEFT', '2': 3},
    {'1': 'UP_LEFT', '2': 4},
    {'1': 'DOWN_LEFT', '2': 5},
    {'1': 'RIGHT', '2': 6},
    {'1': 'UP_RIGHT', '2': 7},
    {'1': 'DOWN_RIGHT', '2': 8},
  ],
};

/// Descriptor for `Move`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List moveDescriptor = $convert.base64Decode(
    'CgRNb3ZlEjIKCWRpcmVjdGlvbhgBIAEoDjIULmdhbWUuTW92ZS5EaXJlY3Rpb25SCWRpcmVjdG'
    'lvbhIeCgplbmFibGVSaW5nGAIgASgIUgplbmFibGVSaW5nEhgKB3VzZXJrZXkYAyABKAVSB3Vz'
    'ZXJrZXkSFgoGZ2FtZWlkGAQgASgFUgZnYW1laWQidgoJRGlyZWN0aW9uEggKBE5PTkUQABIGCg'
    'JVUBABEggKBERPV04QAhIICgRMRUZUEAMSCwoHVVBfTEVGVBAEEg0KCURPV05fTEVGVBAFEgkK'
    'BVJJR0hUEAYSDAoIVVBfUklHSFQQBxIOCgpET1dOX1JJR0hUEAg=');

