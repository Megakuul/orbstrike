// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: game.proto

package game

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Move_Direction int32

const (
	Move_NONE       Move_Direction = 0
	Move_UP         Move_Direction = 1
	Move_DOWN       Move_Direction = 2
	Move_LEFT       Move_Direction = 3
	Move_UP_LEFT    Move_Direction = 4
	Move_DOWN_LEFT  Move_Direction = 5
	Move_RIGHT      Move_Direction = 6
	Move_UP_RIGHT   Move_Direction = 7
	Move_DOWN_RIGHT Move_Direction = 8
)

// Enum value maps for Move_Direction.
var (
	Move_Direction_name = map[int32]string{
		0: "NONE",
		1: "UP",
		2: "DOWN",
		3: "LEFT",
		4: "UP_LEFT",
		5: "DOWN_LEFT",
		6: "RIGHT",
		7: "UP_RIGHT",
		8: "DOWN_RIGHT",
	}
	Move_Direction_value = map[string]int32{
		"NONE":       0,
		"UP":         1,
		"DOWN":       2,
		"LEFT":       3,
		"UP_LEFT":    4,
		"DOWN_LEFT":  5,
		"RIGHT":      6,
		"UP_RIGHT":   7,
		"DOWN_RIGHT": 8,
	}
)

func (x Move_Direction) Enum() *Move_Direction {
	p := new(Move_Direction)
	*p = x
	return p
}

func (x Move_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Move_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_game_proto_enumTypes[0].Descriptor()
}

func (Move_Direction) Type() protoreflect.EnumType {
	return &file_game_proto_enumTypes[0]
}

func (x Move_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Move_Direction.Descriptor instead.
func (Move_Direction) EnumDescriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2, 0}
}

// Gameboard representation
type GameBoard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players       map[int32]*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Id            int32             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Rad           float64           `protobuf:"fixed64,3,opt,name=rad,proto3" json:"rad,omitempty"`
	MaxPlayers    int32             `protobuf:"varint,4,opt,name=maxPlayers,proto3" json:"maxPlayers,omitempty"`
	Speed         float64           `protobuf:"fixed64,5,opt,name=speed,proto3" json:"speed,omitempty"`
	PlayerRad     float64           `protobuf:"fixed64,6,opt,name=playerRad,proto3" json:"playerRad,omitempty"`
	PlayerRingRad float64           `protobuf:"fixed64,7,opt,name=playerRingRad,proto3" json:"playerRingRad,omitempty"`
}

func (x *GameBoard) Reset() {
	*x = GameBoard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameBoard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameBoard) ProtoMessage() {}

func (x *GameBoard) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameBoard.ProtoReflect.Descriptor instead.
func (*GameBoard) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameBoard) GetPlayers() map[int32]*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameBoard) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GameBoard) GetRad() float64 {
	if x != nil {
		return x.Rad
	}
	return 0
}

func (x *GameBoard) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

func (x *GameBoard) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *GameBoard) GetPlayerRad() float64 {
	if x != nil {
		return x.PlayerRad
	}
	return 0
}

func (x *GameBoard) GetPlayerRingRad() float64 {
	if x != nil {
		return x.PlayerRingRad
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	X           float64 `protobuf:"fixed64,3,opt,name=x,proto3" json:"x,omitempty"`
	Y           float64 `protobuf:"fixed64,4,opt,name=y,proto3" json:"y,omitempty"`
	Rad         float64 `protobuf:"fixed64,5,opt,name=rad,proto3" json:"rad,omitempty"`
	Ringrad     float64 `protobuf:"fixed64,6,opt,name=ringrad,proto3" json:"ringrad,omitempty"`
	Color       int32   `protobuf:"varint,7,opt,name=color,proto3" json:"color,omitempty"`
	Kills       int32   `protobuf:"varint,8,opt,name=kills,proto3" json:"kills,omitempty"`
	RingEnabled bool    `protobuf:"varint,9,opt,name=ringEnabled,proto3" json:"ringEnabled,omitempty"`
	Speed       float64 `protobuf:"fixed64,10,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Player) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Player) GetRad() float64 {
	if x != nil {
		return x.Rad
	}
	return 0
}

func (x *Player) GetRingrad() float64 {
	if x != nil {
		return x.Ringrad
	}
	return 0
}

func (x *Player) GetColor() int32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *Player) GetKills() int32 {
	if x != nil {
		return x.Kills
	}
	return 0
}

func (x *Player) GetRingEnabled() bool {
	if x != nil {
		return x.RingEnabled
	}
	return false
}

func (x *Player) GetSpeed() float64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

// Direction to move
type Move struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userkey    []byte         `protobuf:"bytes,1,opt,name=userkey,proto3" json:"userkey,omitempty"`
	Gameid     int32          `protobuf:"varint,2,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Direction  Move_Direction `protobuf:"varint,3,opt,name=direction,proto3,enum=game.Move_Direction" json:"direction,omitempty"`
	EnableRing bool           `protobuf:"varint,4,opt,name=enableRing,proto3" json:"enableRing,omitempty"`
	HitPlayers []int32        `protobuf:"varint,5,rep,packed,name=hitPlayers,proto3" json:"hitPlayers,omitempty"`
}

func (x *Move) Reset() {
	*x = Move{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Move) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Move) ProtoMessage() {}

func (x *Move) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Move.ProtoReflect.Descriptor instead.
func (*Move) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

func (x *Move) GetUserkey() []byte {
	if x != nil {
		return x.Userkey
	}
	return nil
}

func (x *Move) GetGameid() int32 {
	if x != nil {
		return x.Gameid
	}
	return 0
}

func (x *Move) GetDirection() Move_Direction {
	if x != nil {
		return x.Direction
	}
	return Move_NONE
}

func (x *Move) GetEnableRing() bool {
	if x != nil {
		return x.EnableRing
	}
	return false
}

func (x *Move) GetHitPlayers() []int32 {
	if x != nil {
		return x.HitPlayers
	}
	return nil
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0xa9, 0x02, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x12, 0x36, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x72, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6d, 0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x61, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x69, 0x6e,
	0x67, 0x52, 0x61, 0x64, 0x1a, 0x48, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd8,
	0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x72, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x69, 0x6e, 0x67, 0x72, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x69,
	0x6e, 0x67, 0x72, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x22, 0xa4, 0x02, 0x0a, 0x04, 0x4d, 0x6f,
	0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x61,
	0x6d, 0x65, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x69, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x69,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x76, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x50, 0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x4f, 0x57, 0x4e,
	0x5f, 0x4c, 0x45, 0x46, 0x54, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x49, 0x47, 0x48, 0x54,
	0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x07,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x4f, 0x57, 0x4e, 0x5f, 0x52, 0x49, 0x47, 0x48, 0x54, 0x10, 0x08,
	0x32, 0x78, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x33, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x1a, 0x0f, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x47, 0x61,
	0x6d, 0x65, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d,
	0x6f, 0x76, 0x65, 0x1a, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_proto_goTypes = []interface{}{
	(Move_Direction)(0), // 0: game.Move.Direction
	(*GameBoard)(nil),   // 1: game.GameBoard
	(*Player)(nil),      // 2: game.Player
	(*Move)(nil),        // 3: game.Move
	nil,                 // 4: game.GameBoard.PlayersEntry
}
var file_game_proto_depIdxs = []int32{
	4, // 0: game.GameBoard.players:type_name -> game.GameBoard.PlayersEntry
	0, // 1: game.Move.direction:type_name -> game.Move.Direction
	2, // 2: game.GameBoard.PlayersEntry.value:type_name -> game.Player
	3, // 3: game.GameService.ProxyGameboard:input_type -> game.Move
	3, // 4: game.GameService.StreamGameboard:input_type -> game.Move
	1, // 5: game.GameService.ProxyGameboard:output_type -> game.GameBoard
	1, // 6: game.GameService.StreamGameboard:output_type -> game.GameBoard
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameBoard); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Move); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		EnumInfos:         file_game_proto_enumTypes,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}