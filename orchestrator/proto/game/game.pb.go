// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package game

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

var Move_Direction_name = map[int32]string{
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

var Move_Direction_value = map[string]int32{
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

func (x Move_Direction) String() string {
	return proto.EnumName(Move_Direction_name, int32(x))
}

func (Move_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2, 0}
}

// Gameboard representation
type GameBoard struct {
	Players              map[int32]*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Id                   int32             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Rad                  float64           `protobuf:"fixed64,3,opt,name=rad,proto3" json:"rad,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GameBoard) Reset()         { *m = GameBoard{} }
func (m *GameBoard) String() string { return proto.CompactTextString(m) }
func (*GameBoard) ProtoMessage()    {}
func (*GameBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *GameBoard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameBoard.Unmarshal(m, b)
}
func (m *GameBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameBoard.Marshal(b, m, deterministic)
}
func (m *GameBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameBoard.Merge(m, src)
}
func (m *GameBoard) XXX_Size() int {
	return xxx_messageInfo_GameBoard.Size(m)
}
func (m *GameBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_GameBoard.DiscardUnknown(m)
}

var xxx_messageInfo_GameBoard proto.InternalMessageInfo

func (m *GameBoard) GetPlayers() map[int32]*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *GameBoard) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GameBoard) GetRad() float64 {
	if m != nil {
		return m.Rad
	}
	return 0
}

type Player struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X                    float64  `protobuf:"fixed64,2,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,3,opt,name=y,proto3" json:"y,omitempty"`
	Rad                  float64  `protobuf:"fixed64,4,opt,name=rad,proto3" json:"rad,omitempty"`
	Ringrad              float64  `protobuf:"fixed64,5,opt,name=ringrad,proto3" json:"ringrad,omitempty"`
	Color                int32    `protobuf:"varint,6,opt,name=color,proto3" json:"color,omitempty"`
	Kills                int32    `protobuf:"varint,7,opt,name=kills,proto3" json:"kills,omitempty"`
	RingEnabled          bool     `protobuf:"varint,8,opt,name=ringEnabled,proto3" json:"ringEnabled,omitempty"`
	Speed                int32    `protobuf:"varint,9,opt,name=speed,proto3" json:"speed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Player) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Player) GetRad() float64 {
	if m != nil {
		return m.Rad
	}
	return 0
}

func (m *Player) GetRingrad() float64 {
	if m != nil {
		return m.Ringrad
	}
	return 0
}

func (m *Player) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Player) GetKills() int32 {
	if m != nil {
		return m.Kills
	}
	return 0
}

func (m *Player) GetRingEnabled() bool {
	if m != nil {
		return m.RingEnabled
	}
	return false
}

func (m *Player) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

// Direction to move
type Move struct {
	Userkey              int32          `protobuf:"varint,1,opt,name=userkey,proto3" json:"userkey,omitempty"`
	Gameid               int32          `protobuf:"varint,2,opt,name=gameid,proto3" json:"gameid,omitempty"`
	Direction            Move_Direction `protobuf:"varint,3,opt,name=direction,proto3,enum=game.Move_Direction" json:"direction,omitempty"`
	EnableRing           bool           `protobuf:"varint,4,opt,name=enableRing,proto3" json:"enableRing,omitempty"`
	HitPlayers           []int32        `protobuf:"varint,5,rep,packed,name=hitPlayers,proto3" json:"hitPlayers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Move) Reset()         { *m = Move{} }
func (m *Move) String() string { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()    {}
func (*Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Move.Unmarshal(m, b)
}
func (m *Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Move.Marshal(b, m, deterministic)
}
func (m *Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Move.Merge(m, src)
}
func (m *Move) XXX_Size() int {
	return xxx_messageInfo_Move.Size(m)
}
func (m *Move) XXX_DiscardUnknown() {
	xxx_messageInfo_Move.DiscardUnknown(m)
}

var xxx_messageInfo_Move proto.InternalMessageInfo

func (m *Move) GetUserkey() int32 {
	if m != nil {
		return m.Userkey
	}
	return 0
}

func (m *Move) GetGameid() int32 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

func (m *Move) GetDirection() Move_Direction {
	if m != nil {
		return m.Direction
	}
	return Move_NONE
}

func (m *Move) GetEnableRing() bool {
	if m != nil {
		return m.EnableRing
	}
	return false
}

func (m *Move) GetHitPlayers() []int32 {
	if m != nil {
		return m.HitPlayers
	}
	return nil
}

func init() {
	proto.RegisterEnum("game.Move_Direction", Move_Direction_name, Move_Direction_value)
	proto.RegisterType((*GameBoard)(nil), "game.GameBoard")
	proto.RegisterMapType((map[int32]*Player)(nil), "game.GameBoard.PlayersEntry")
	proto.RegisterType((*Player)(nil), "game.Player")
	proto.RegisterType((*Move)(nil), "game.Move")
}

func init() {
	proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769)
}

var fileDescriptor_38fc58335341d769 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6b, 0xdb, 0x4e,
	0x10, 0xcd, 0xca, 0xfa, 0x3b, 0xf2, 0xcf, 0x11, 0x43, 0xf8, 0xb1, 0x84, 0x52, 0x84, 0x4e, 0x3a,
	0xa9, 0x45, 0x29, 0xa5, 0xf4, 0x18, 0xe2, 0x26, 0x85, 0xd6, 0x11, 0x9b, 0x98, 0x42, 0x2f, 0x41,
	0xb6, 0x16, 0x57, 0x44, 0x96, 0xcc, 0x5a, 0x31, 0xd6, 0xf7, 0xe9, 0xbd, 0x5f, 0xa2, 0x1f, 0xac,
	0xec, 0xae, 0x25, 0x9b, 0xde, 0x7a, 0x9b, 0xf7, 0xe6, 0xbd, 0x99, 0xd9, 0xd1, 0x08, 0x60, 0x95,
	0xaf, 0x79, 0xb2, 0x11, 0x4d, 0xdb, 0xa0, 0x29, 0xe3, 0xe8, 0x17, 0x01, 0xef, 0x36, 0x5f, 0xf3,
	0xeb, 0x26, 0x17, 0x05, 0xbe, 0x07, 0x67, 0x53, 0xe5, 0x1d, 0x17, 0x5b, 0x4a, 0xc2, 0x51, 0xec,
	0xa7, 0xaf, 0x12, 0xe5, 0x18, 0x14, 0x49, 0xa6, 0xd3, 0xd3, 0xba, 0x15, 0x1d, 0xeb, 0xc5, 0x38,
	0x01, 0xa3, 0x2c, 0xa8, 0x11, 0x92, 0xd8, 0x62, 0x46, 0x59, 0x60, 0x00, 0x23, 0x91, 0x17, 0x74,
	0x14, 0x92, 0x98, 0x30, 0x19, 0x5e, 0xde, 0xc1, 0xf8, 0xd4, 0x2a, 0x15, 0xcf, 0xbc, 0xa3, 0x44,
	0x59, 0x64, 0x88, 0x11, 0x58, 0xbb, 0xbc, 0x7a, 0xe1, 0xaa, 0x8c, 0x9f, 0x8e, 0x75, 0x67, 0x6d,
	0x62, 0x3a, 0xf5, 0xd1, 0xf8, 0x40, 0xa2, 0xdf, 0x04, 0x6c, 0xcd, 0x1e, 0xda, 0x92, 0xa1, 0xed,
	0x18, 0xc8, 0x5e, 0xd9, 0x09, 0x23, 0x7b, 0x89, 0xba, 0xc3, 0x08, 0xa4, 0xeb, 0x47, 0x32, 0x87,
	0x91, 0x90, 0x82, 0x23, 0xca, 0x7a, 0x25, 0x59, 0x4b, 0xb1, 0x3d, 0xc4, 0x0b, 0xb0, 0x96, 0x4d,
	0xd5, 0x08, 0x6a, 0xab, 0xd2, 0x1a, 0x48, 0xf6, 0xb9, 0xac, 0xaa, 0x2d, 0x75, 0x34, 0xab, 0x00,
	0x86, 0xe0, 0x4b, 0xdb, 0xb4, 0xce, 0x17, 0x15, 0x2f, 0xa8, 0x1b, 0x92, 0xd8, 0x65, 0xa7, 0x94,
	0xf4, 0x6d, 0x37, 0x9c, 0x17, 0xd4, 0xd3, 0x3e, 0x05, 0xa2, 0x9f, 0x06, 0x98, 0x5f, 0x9b, 0x1d,
	0x97, 0x63, 0xbc, 0x6c, 0xb9, 0x38, 0x6e, 0xa3, 0x87, 0xf8, 0x3f, 0xd8, 0x72, 0x07, 0xc3, 0x66,
	0x0f, 0x08, 0x53, 0xf0, 0x8a, 0x52, 0xf0, 0x65, 0x5b, 0x36, 0xb5, 0x7a, 0xe0, 0x24, 0xbd, 0xd0,
	0xdb, 0x92, 0x05, 0x93, 0x9b, 0x3e, 0xc7, 0x8e, 0x32, 0x7c, 0x0d, 0xc0, 0xd5, 0x3c, 0xac, 0xac,
	0x57, 0x6a, 0x0b, 0x2e, 0x3b, 0x61, 0x64, 0xfe, 0x47, 0xd9, 0x1e, 0x3e, 0x11, 0xb5, 0xc2, 0x51,
	0x6c, 0xb1, 0x13, 0x26, 0xda, 0x81, 0x37, 0xd4, 0x45, 0x17, 0xcc, 0xd9, 0xfd, 0x6c, 0x1a, 0x9c,
	0xa1, 0x0d, 0xc6, 0x3c, 0x0b, 0x88, 0x64, 0x6e, 0xee, 0xbf, 0xcd, 0x02, 0x43, 0x46, 0x5f, 0xa6,
	0x9f, 0x1e, 0x83, 0x11, 0xfa, 0xe0, 0xcc, 0xb3, 0x27, 0x05, 0x4c, 0xfc, 0x0f, 0x3c, 0x29, 0xd0,
	0xd0, 0x42, 0x0f, 0x2c, 0xf6, 0xf9, 0xf6, 0xee, 0x31, 0xb0, 0x71, 0x0c, 0xee, 0x3c, 0x7b, 0xd2,
	0xc8, 0xc1, 0x09, 0x80, 0xd2, 0x69, 0xec, 0xa6, 0x7b, 0xf0, 0xe5, 0xf1, 0x3d, 0x70, 0xb1, 0x2b,
	0x97, 0x1c, 0xaf, 0x60, 0x92, 0x89, 0x66, 0xdf, 0x49, 0x6e, 0xa1, 0x4e, 0x16, 0x8e, 0x2f, 0xbf,
	0x3c, 0xff, 0xeb, 0x5a, 0xa3, 0xb3, 0x98, 0xbc, 0x25, 0xf8, 0x0e, 0xce, 0x1f, 0x5a, 0xc1, 0xf3,
	0xf5, 0xbf, 0xb8, 0xae, 0xdd, 0xef, 0x76, 0xf2, 0x46, 0x66, 0x16, 0xb6, 0xfa, 0x61, 0xae, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xa0, 0x82, 0xd0, 0x3e, 0x03, 0x00, 0x00,
}
