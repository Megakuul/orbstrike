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

import 'game.pbenum.dart';

export 'game.pbenum.dart';

/// Gameboard representation
class GameBoard extends $pb.GeneratedMessage {
  factory GameBoard({
    $core.Map<$core.int, Player>? players,
    $core.int? id,
    $core.double? rad,
  }) {
    final $result = create();
    if (players != null) {
      $result.players.addAll(players);
    }
    if (id != null) {
      $result.id = id;
    }
    if (rad != null) {
      $result.rad = rad;
    }
    return $result;
  }
  GameBoard._() : super();
  factory GameBoard.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GameBoard.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'GameBoard', package: const $pb.PackageName(_omitMessageNames ? '' : 'game'), createEmptyInstance: create)
    ..m<$core.int, Player>(1, _omitFieldNames ? '' : 'players', entryClassName: 'GameBoard.PlayersEntry', keyFieldType: $pb.PbFieldType.O3, valueFieldType: $pb.PbFieldType.OM, valueCreator: Player.create, valueDefaultOrMaker: Player.getDefault, packageName: const $pb.PackageName('game'))
    ..a<$core.int>(2, _omitFieldNames ? '' : 'id', $pb.PbFieldType.O3)
    ..a<$core.double>(3, _omitFieldNames ? '' : 'rad', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GameBoard clone() => GameBoard()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GameBoard copyWith(void Function(GameBoard) updates) => super.copyWith((message) => updates(message as GameBoard)) as GameBoard;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static GameBoard create() => GameBoard._();
  GameBoard createEmptyInstance() => create();
  static $pb.PbList<GameBoard> createRepeated() => $pb.PbList<GameBoard>();
  @$core.pragma('dart2js:noInline')
  static GameBoard getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GameBoard>(create);
  static GameBoard? _defaultInstance;

  @$pb.TagNumber(1)
  $core.Map<$core.int, Player> get players => $_getMap(0);

  @$pb.TagNumber(2)
  $core.int get id => $_getIZ(1);
  @$pb.TagNumber(2)
  set id($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(1);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get rad => $_getN(2);
  @$pb.TagNumber(3)
  set rad($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRad() => $_has(2);
  @$pb.TagNumber(3)
  void clearRad() => clearField(3);
}

class Player extends $pb.GeneratedMessage {
  factory Player({
    $core.int? id,
    $core.double? x,
    $core.double? y,
    $core.double? rad,
    $core.double? ringrad,
    $core.int? color,
    $core.int? kills,
    $core.bool? ringEnabled,
    $core.int? speed,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (x != null) {
      $result.x = x;
    }
    if (y != null) {
      $result.y = y;
    }
    if (rad != null) {
      $result.rad = rad;
    }
    if (ringrad != null) {
      $result.ringrad = ringrad;
    }
    if (color != null) {
      $result.color = color;
    }
    if (kills != null) {
      $result.kills = kills;
    }
    if (ringEnabled != null) {
      $result.ringEnabled = ringEnabled;
    }
    if (speed != null) {
      $result.speed = speed;
    }
    return $result;
  }
  Player._() : super();
  factory Player.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Player.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Player', package: const $pb.PackageName(_omitMessageNames ? '' : 'game'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'id', $pb.PbFieldType.O3)
    ..a<$core.double>(2, _omitFieldNames ? '' : 'x', $pb.PbFieldType.OD)
    ..a<$core.double>(3, _omitFieldNames ? '' : 'y', $pb.PbFieldType.OD)
    ..a<$core.double>(4, _omitFieldNames ? '' : 'rad', $pb.PbFieldType.OD)
    ..a<$core.double>(5, _omitFieldNames ? '' : 'ringrad', $pb.PbFieldType.OD)
    ..a<$core.int>(6, _omitFieldNames ? '' : 'color', $pb.PbFieldType.O3)
    ..a<$core.int>(7, _omitFieldNames ? '' : 'kills', $pb.PbFieldType.O3)
    ..aOB(8, _omitFieldNames ? '' : 'ringEnabled', protoName: 'ringEnabled')
    ..a<$core.int>(9, _omitFieldNames ? '' : 'speed', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Player clone() => Player()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Player copyWith(void Function(Player) updates) => super.copyWith((message) => updates(message as Player)) as Player;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Player create() => Player._();
  Player createEmptyInstance() => create();
  static $pb.PbList<Player> createRepeated() => $pb.PbList<Player>();
  @$core.pragma('dart2js:noInline')
  static Player getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Player>(create);
  static Player? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get x => $_getN(1);
  @$pb.TagNumber(2)
  set x($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasX() => $_has(1);
  @$pb.TagNumber(2)
  void clearX() => clearField(2);

  @$pb.TagNumber(3)
  $core.double get y => $_getN(2);
  @$pb.TagNumber(3)
  set y($core.double v) { $_setDouble(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasY() => $_has(2);
  @$pb.TagNumber(3)
  void clearY() => clearField(3);

  @$pb.TagNumber(4)
  $core.double get rad => $_getN(3);
  @$pb.TagNumber(4)
  set rad($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRad() => $_has(3);
  @$pb.TagNumber(4)
  void clearRad() => clearField(4);

  @$pb.TagNumber(5)
  $core.double get ringrad => $_getN(4);
  @$pb.TagNumber(5)
  set ringrad($core.double v) { $_setDouble(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasRingrad() => $_has(4);
  @$pb.TagNumber(5)
  void clearRingrad() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get color => $_getIZ(5);
  @$pb.TagNumber(6)
  set color($core.int v) { $_setSignedInt32(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasColor() => $_has(5);
  @$pb.TagNumber(6)
  void clearColor() => clearField(6);

  @$pb.TagNumber(7)
  $core.int get kills => $_getIZ(6);
  @$pb.TagNumber(7)
  set kills($core.int v) { $_setSignedInt32(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasKills() => $_has(6);
  @$pb.TagNumber(7)
  void clearKills() => clearField(7);

  @$pb.TagNumber(8)
  $core.bool get ringEnabled => $_getBF(7);
  @$pb.TagNumber(8)
  set ringEnabled($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasRingEnabled() => $_has(7);
  @$pb.TagNumber(8)
  void clearRingEnabled() => clearField(8);

  @$pb.TagNumber(9)
  $core.int get speed => $_getIZ(8);
  @$pb.TagNumber(9)
  set speed($core.int v) { $_setSignedInt32(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasSpeed() => $_has(8);
  @$pb.TagNumber(9)
  void clearSpeed() => clearField(9);
}

/// Direction to move
class Move extends $pb.GeneratedMessage {
  factory Move({
    $core.int? userkey,
    $core.int? gameid,
    Move_Direction? direction,
    $core.bool? enableRing,
    $core.Iterable<$core.int>? hitPlayers,
  }) {
    final $result = create();
    if (userkey != null) {
      $result.userkey = userkey;
    }
    if (gameid != null) {
      $result.gameid = gameid;
    }
    if (direction != null) {
      $result.direction = direction;
    }
    if (enableRing != null) {
      $result.enableRing = enableRing;
    }
    if (hitPlayers != null) {
      $result.hitPlayers.addAll(hitPlayers);
    }
    return $result;
  }
  Move._() : super();
  factory Move.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Move.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Move', package: const $pb.PackageName(_omitMessageNames ? '' : 'game'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'userkey', $pb.PbFieldType.O3)
    ..a<$core.int>(2, _omitFieldNames ? '' : 'gameid', $pb.PbFieldType.O3)
    ..e<Move_Direction>(3, _omitFieldNames ? '' : 'direction', $pb.PbFieldType.OE, defaultOrMaker: Move_Direction.NONE, valueOf: Move_Direction.valueOf, enumValues: Move_Direction.values)
    ..aOB(4, _omitFieldNames ? '' : 'enableRing', protoName: 'enableRing')
    ..p<$core.int>(5, _omitFieldNames ? '' : 'hitPlayers', $pb.PbFieldType.K3, protoName: 'hitPlayers')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Move clone() => Move()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Move copyWith(void Function(Move) updates) => super.copyWith((message) => updates(message as Move)) as Move;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Move create() => Move._();
  Move createEmptyInstance() => create();
  static $pb.PbList<Move> createRepeated() => $pb.PbList<Move>();
  @$core.pragma('dart2js:noInline')
  static Move getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Move>(create);
  static Move? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get userkey => $_getIZ(0);
  @$pb.TagNumber(1)
  set userkey($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserkey() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserkey() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get gameid => $_getIZ(1);
  @$pb.TagNumber(2)
  set gameid($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasGameid() => $_has(1);
  @$pb.TagNumber(2)
  void clearGameid() => clearField(2);

  @$pb.TagNumber(3)
  Move_Direction get direction => $_getN(2);
  @$pb.TagNumber(3)
  set direction(Move_Direction v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasDirection() => $_has(2);
  @$pb.TagNumber(3)
  void clearDirection() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get enableRing => $_getBF(3);
  @$pb.TagNumber(4)
  set enableRing($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasEnableRing() => $_has(3);
  @$pb.TagNumber(4)
  void clearEnableRing() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get hitPlayers => $_getList(4);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
