//
//  Generated code. Do not modify.
//  source: auth.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use createGameRequestDescriptor instead')
const CreateGameRequest$json = {
  '1': 'CreateGameRequest',
  '2': [
    {'1': 'identifier', '3': 1, '4': 1, '5': 9, '10': 'identifier'},
  ],
};

/// Descriptor for `CreateGameRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createGameRequestDescriptor = $convert.base64Decode(
    'ChFDcmVhdGVHYW1lUmVxdWVzdBIeCgppZGVudGlmaWVyGAEgASgJUgppZGVudGlmaWVy');

@$core.Deprecated('Use createGameResponseDescriptor instead')
const CreateGameResponse$json = {
  '1': 'CreateGameResponse',
  '2': [
    {'1': 'gameid', '3': 1, '4': 1, '5': 5, '10': 'gameid'},
  ],
};

/// Descriptor for `CreateGameResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createGameResponseDescriptor = $convert.base64Decode(
    'ChJDcmVhdGVHYW1lUmVzcG9uc2USFgoGZ2FtZWlkGAEgASgFUgZnYW1laWQ=');

@$core.Deprecated('Use joinGameRequestDescriptor instead')
const JoinGameRequest$json = {
  '1': 'JoinGameRequest',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'gameid', '3': 2, '4': 1, '5': 5, '10': 'gameid'},
  ],
};

/// Descriptor for `JoinGameRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List joinGameRequestDescriptor = $convert.base64Decode(
    'Cg9Kb2luR2FtZVJlcXVlc3QSEgoEbmFtZRgBIAEoCVIEbmFtZRIWCgZnYW1laWQYAiABKAVSBm'
    'dhbWVpZA==');

@$core.Deprecated('Use joinGameResponseDescriptor instead')
const JoinGameResponse$json = {
  '1': 'JoinGameResponse',
  '2': [
    {'1': 'userkey', '3': 1, '4': 1, '5': 5, '10': 'userkey'},
  ],
};

/// Descriptor for `JoinGameResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List joinGameResponseDescriptor = $convert.base64Decode(
    'ChBKb2luR2FtZVJlc3BvbnNlEhgKB3VzZXJrZXkYASABKAVSB3VzZXJrZXk=');

