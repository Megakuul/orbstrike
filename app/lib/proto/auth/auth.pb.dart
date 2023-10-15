//
//  Generated code. Do not modify.
//  source: auth.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class CreateGameRequest extends $pb.GeneratedMessage {
  factory CreateGameRequest({
    $core.String? identifier,
    $core.double? radius,
    $core.int? maxplayers,
    $core.double? speed,
    $core.double? playerradius,
    $core.double? playerringradius,
  }) {
    final $result = create();
    if (identifier != null) {
      $result.identifier = identifier;
    }
    if (radius != null) {
      $result.radius = radius;
    }
    if (maxplayers != null) {
      $result.maxplayers = maxplayers;
    }
    if (speed != null) {
      $result.speed = speed;
    }
    if (playerradius != null) {
      $result.playerradius = playerradius;
    }
    if (playerringradius != null) {
      $result.playerringradius = playerringradius;
    }
    return $result;
  }
  CreateGameRequest._() : super();
  factory CreateGameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateGameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateGameRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'auth'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'identifier')
    ..a<$core.double>(2, _omitFieldNames ? '' : 'radius', $pb.PbFieldType.OD)
    ..a<$core.int>(3, _omitFieldNames ? '' : 'maxplayers', $pb.PbFieldType.O3)
    ..a<$core.double>(4, _omitFieldNames ? '' : 'speed', $pb.PbFieldType.OD)
    ..a<$core.double>(5, _omitFieldNames ? '' : 'playerradius', $pb.PbFieldType.OD)
    ..a<$core.double>(6, _omitFieldNames ? '' : 'playerringradius', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateGameRequest clone() => CreateGameRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateGameRequest copyWith(void Function(CreateGameRequest) updates) => super.copyWith((message) => updates(message as CreateGameRequest)) as CreateGameRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateGameRequest create() => CreateGameRequest._();
  CreateGameRequest createEmptyInstance() => create();
  static $pb.PbList<CreateGameRequest> createRepeated() => $pb.PbList<CreateGameRequest>();
  @$core.pragma('dart2js:noInline')
  static CreateGameRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateGameRequest>(create);
  static CreateGameRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get identifier => $_getSZ(0);
  @$pb.TagNumber(1)
  set identifier($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIdentifier() => $_has(0);
  @$pb.TagNumber(1)
  void clearIdentifier() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get radius => $_getN(1);
  @$pb.TagNumber(2)
  set radius($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRadius() => $_has(1);
  @$pb.TagNumber(2)
  void clearRadius() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get maxplayers => $_getIZ(2);
  @$pb.TagNumber(3)
  set maxplayers($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMaxplayers() => $_has(2);
  @$pb.TagNumber(3)
  void clearMaxplayers() => clearField(3);

  @$pb.TagNumber(4)
  $core.double get speed => $_getN(3);
  @$pb.TagNumber(4)
  set speed($core.double v) { $_setDouble(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasSpeed() => $_has(3);
  @$pb.TagNumber(4)
  void clearSpeed() => clearField(4);

  @$pb.TagNumber(5)
  $core.double get playerradius => $_getN(4);
  @$pb.TagNumber(5)
  set playerradius($core.double v) { $_setDouble(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPlayerradius() => $_has(4);
  @$pb.TagNumber(5)
  void clearPlayerradius() => clearField(5);

  @$pb.TagNumber(6)
  $core.double get playerringradius => $_getN(5);
  @$pb.TagNumber(6)
  set playerringradius($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasPlayerringradius() => $_has(5);
  @$pb.TagNumber(6)
  void clearPlayerringradius() => clearField(6);
}

class CreateGameResponse extends $pb.GeneratedMessage {
  factory CreateGameResponse({
    $core.int? gameid,
  }) {
    final $result = create();
    if (gameid != null) {
      $result.gameid = gameid;
    }
    return $result;
  }
  CreateGameResponse._() : super();
  factory CreateGameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateGameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'CreateGameResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'auth'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'gameid', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateGameResponse clone() => CreateGameResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateGameResponse copyWith(void Function(CreateGameResponse) updates) => super.copyWith((message) => updates(message as CreateGameResponse)) as CreateGameResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateGameResponse create() => CreateGameResponse._();
  CreateGameResponse createEmptyInstance() => create();
  static $pb.PbList<CreateGameResponse> createRepeated() => $pb.PbList<CreateGameResponse>();
  @$core.pragma('dart2js:noInline')
  static CreateGameResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateGameResponse>(create);
  static CreateGameResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get gameid => $_getIZ(0);
  @$pb.TagNumber(1)
  set gameid($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGameid() => $_has(0);
  @$pb.TagNumber(1)
  void clearGameid() => clearField(1);
}

class JoinGameRequest extends $pb.GeneratedMessage {
  factory JoinGameRequest({
    $core.String? name,
    $core.int? gameid,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (gameid != null) {
      $result.gameid = gameid;
    }
    return $result;
  }
  JoinGameRequest._() : super();
  factory JoinGameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory JoinGameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'JoinGameRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'auth'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'name')
    ..a<$core.int>(2, _omitFieldNames ? '' : 'gameid', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  JoinGameRequest clone() => JoinGameRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  JoinGameRequest copyWith(void Function(JoinGameRequest) updates) => super.copyWith((message) => updates(message as JoinGameRequest)) as JoinGameRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static JoinGameRequest create() => JoinGameRequest._();
  JoinGameRequest createEmptyInstance() => create();
  static $pb.PbList<JoinGameRequest> createRepeated() => $pb.PbList<JoinGameRequest>();
  @$core.pragma('dart2js:noInline')
  static JoinGameRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<JoinGameRequest>(create);
  static JoinGameRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get gameid => $_getIZ(1);
  @$pb.TagNumber(2)
  set gameid($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasGameid() => $_has(1);
  @$pb.TagNumber(2)
  void clearGameid() => clearField(2);
}

class JoinGameResponse extends $pb.GeneratedMessage {
  factory JoinGameResponse({
    $core.int? userid,
    $core.List<$core.int>? userkey,
  }) {
    final $result = create();
    if (userid != null) {
      $result.userid = userid;
    }
    if (userkey != null) {
      $result.userkey = userkey;
    }
    return $result;
  }
  JoinGameResponse._() : super();
  factory JoinGameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory JoinGameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'JoinGameResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'auth'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'userid', $pb.PbFieldType.O3)
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'userkey', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  JoinGameResponse clone() => JoinGameResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  JoinGameResponse copyWith(void Function(JoinGameResponse) updates) => super.copyWith((message) => updates(message as JoinGameResponse)) as JoinGameResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static JoinGameResponse create() => JoinGameResponse._();
  JoinGameResponse createEmptyInstance() => create();
  static $pb.PbList<JoinGameResponse> createRepeated() => $pb.PbList<JoinGameResponse>();
  @$core.pragma('dart2js:noInline')
  static JoinGameResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<JoinGameResponse>(create);
  static JoinGameResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get userid => $_getIZ(0);
  @$pb.TagNumber(1)
  set userid($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUserid() => $_has(0);
  @$pb.TagNumber(1)
  void clearUserid() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get userkey => $_getN(1);
  @$pb.TagNumber(2)
  set userkey($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUserkey() => $_has(1);
  @$pb.TagNumber(2)
  void clearUserkey() => clearField(2);
}

class ExitGameRequest extends $pb.GeneratedMessage {
  factory ExitGameRequest({
    $core.int? gameid,
    $core.List<$core.int>? userkey,
  }) {
    final $result = create();
    if (gameid != null) {
      $result.gameid = gameid;
    }
    if (userkey != null) {
      $result.userkey = userkey;
    }
    return $result;
  }
  ExitGameRequest._() : super();
  factory ExitGameRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ExitGameRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ExitGameRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'auth'), createEmptyInstance: create)
    ..a<$core.int>(1, _omitFieldNames ? '' : 'gameid', $pb.PbFieldType.O3)
    ..a<$core.List<$core.int>>(2, _omitFieldNames ? '' : 'userkey', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ExitGameRequest clone() => ExitGameRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ExitGameRequest copyWith(void Function(ExitGameRequest) updates) => super.copyWith((message) => updates(message as ExitGameRequest)) as ExitGameRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ExitGameRequest create() => ExitGameRequest._();
  ExitGameRequest createEmptyInstance() => create();
  static $pb.PbList<ExitGameRequest> createRepeated() => $pb.PbList<ExitGameRequest>();
  @$core.pragma('dart2js:noInline')
  static ExitGameRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ExitGameRequest>(create);
  static ExitGameRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get gameid => $_getIZ(0);
  @$pb.TagNumber(1)
  set gameid($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasGameid() => $_has(0);
  @$pb.TagNumber(1)
  void clearGameid() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<$core.int> get userkey => $_getN(1);
  @$pb.TagNumber(2)
  set userkey($core.List<$core.int> v) { $_setBytes(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUserkey() => $_has(1);
  @$pb.TagNumber(2)
  void clearUserkey() => clearField(2);
}

class ExitGameResponse extends $pb.GeneratedMessage {
  factory ExitGameResponse() => create();
  ExitGameResponse._() : super();
  factory ExitGameResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ExitGameResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'ExitGameResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'auth'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ExitGameResponse clone() => ExitGameResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ExitGameResponse copyWith(void Function(ExitGameResponse) updates) => super.copyWith((message) => updates(message as ExitGameResponse)) as ExitGameResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ExitGameResponse create() => ExitGameResponse._();
  ExitGameResponse createEmptyInstance() => create();
  static $pb.PbList<ExitGameResponse> createRepeated() => $pb.PbList<ExitGameResponse>();
  @$core.pragma('dart2js:noInline')
  static ExitGameResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ExitGameResponse>(create);
  static ExitGameResponse? _defaultInstance;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
