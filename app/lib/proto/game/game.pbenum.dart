//
//  Generated code. Do not modify.
//  source: game.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class Move_Direction extends $pb.ProtobufEnum {
  static const Move_Direction NONE = Move_Direction._(0, _omitEnumNames ? '' : 'NONE');
  static const Move_Direction UP = Move_Direction._(1, _omitEnumNames ? '' : 'UP');
  static const Move_Direction DOWN = Move_Direction._(2, _omitEnumNames ? '' : 'DOWN');
  static const Move_Direction LEFT = Move_Direction._(3, _omitEnumNames ? '' : 'LEFT');
  static const Move_Direction UP_LEFT = Move_Direction._(4, _omitEnumNames ? '' : 'UP_LEFT');
  static const Move_Direction DOWN_LEFT = Move_Direction._(5, _omitEnumNames ? '' : 'DOWN_LEFT');
  static const Move_Direction RIGHT = Move_Direction._(6, _omitEnumNames ? '' : 'RIGHT');
  static const Move_Direction UP_RIGHT = Move_Direction._(7, _omitEnumNames ? '' : 'UP_RIGHT');
  static const Move_Direction DOWN_RIGHT = Move_Direction._(8, _omitEnumNames ? '' : 'DOWN_RIGHT');

  static const $core.List<Move_Direction> values = <Move_Direction> [
    NONE,
    UP,
    DOWN,
    LEFT,
    UP_LEFT,
    DOWN_LEFT,
    RIGHT,
    UP_RIGHT,
    DOWN_RIGHT,
  ];

  static final $core.Map<$core.int, Move_Direction> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Move_Direction? valueOf($core.int value) => _byValue[value];

  const Move_Direction._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
