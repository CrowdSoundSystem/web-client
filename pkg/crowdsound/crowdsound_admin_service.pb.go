// Code generated by protoc-gen-go.
// source: pkg/crowdsound/crowdsound_admin_service.proto
// DO NOT EDIT!

package crowdsound

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SkipStatusRequest struct {
}

func (m *SkipStatusRequest) Reset()                    { *m = SkipStatusRequest{} }
func (m *SkipStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*SkipStatusRequest) ProtoMessage()               {}
func (*SkipStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type SkipStatusResponse struct {
	VotesToSkip int32   `protobuf:"varint,1,opt,name=votes_to_skip" json:"votes_to_skip,omitempty"`
	TotalUsers  int32   `protobuf:"varint,2,opt,name=total_users" json:"total_users,omitempty"`
	Threshold   float32 `protobuf:"fixed32,3,opt,name=threshold" json:"threshold,omitempty"`
}

func (m *SkipStatusResponse) Reset()                    { *m = SkipStatusResponse{} }
func (m *SkipStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*SkipStatusResponse) ProtoMessage()               {}
func (*SkipStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type SkipRequest struct {
}

func (m *SkipRequest) Reset()                    { *m = SkipRequest{} }
func (m *SkipRequest) String() string            { return proto.CompactTextString(m) }
func (*SkipRequest) ProtoMessage()               {}
func (*SkipRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type SkipResponse struct {
}

func (m *SkipResponse) Reset()                    { *m = SkipResponse{} }
func (m *SkipResponse) String() string            { return proto.CompactTextString(m) }
func (*SkipResponse) ProtoMessage()               {}
func (*SkipResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type GetSettingsRequest struct {
}

func (m *GetSettingsRequest) Reset()                    { *m = GetSettingsRequest{} }
func (m *GetSettingsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSettingsRequest) ProtoMessage()               {}
func (*GetSettingsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type GetSettingsResponse struct {
	FilterBuffered      bool    `protobuf:"varint,1,opt,name=filter_buffered" json:"filter_buffered,omitempty"`
	InactivityThreshold int32   `protobuf:"varint,2,opt,name=inactivity_threshold" json:"inactivity_threshold,omitempty"`
	ResultLimit         int32   `protobuf:"varint,3,opt,name=result_limit" json:"result_limit,omitempty"`
	SessionName         string  `protobuf:"bytes,4,opt,name=session_name" json:"session_name,omitempty"`
	QueueSize           int32   `protobuf:"varint,5,opt,name=queue_size" json:"queue_size,omitempty"`
	TrendingArtistsSize int32   `protobuf:"varint,6,opt,name=trending_artists_size" json:"trending_artists_size,omitempty"`
	SkipThreshold       float32 `protobuf:"fixed32,7,opt,name=skip_threshold" json:"skip_threshold,omitempty"`
	CountWeight         float32 `protobuf:"fixed32,8,opt,name=count_weight" json:"count_weight,omitempty"`
	VoteWeight          float32 `protobuf:"fixed32,9,opt,name=vote_weight" json:"vote_weight,omitempty"`
	GenreWeight         float32 `protobuf:"fixed32,10,opt,name=genre_weight" json:"genre_weight,omitempty"`
	ArtistWeight        float32 `protobuf:"fixed32,11,opt,name=artist_weight" json:"artist_weight,omitempty"`
	PlayedAgainMult     float32 `protobuf:"fixed32,12,opt,name=played_again_mult" json:"played_again_mult,omitempty"`
	MinRepeatWindow     float32 `protobuf:"fixed32,13,opt,name=min_repeat_window" json:"min_repeat_window,omitempty"`
}

func (m *GetSettingsResponse) Reset()                    { *m = GetSettingsResponse{} }
func (m *GetSettingsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSettingsResponse) ProtoMessage()               {}
func (*GetSettingsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type SetSettingRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*SetSettingRequest_BoolVal
	//	*SetSettingRequest_StrVal
	//	*SetSettingRequest_IntVal
	//	*SetSettingRequest_FloatVal
	Value isSetSettingRequest_Value `protobuf_oneof:"value"`
}

func (m *SetSettingRequest) Reset()                    { *m = SetSettingRequest{} }
func (m *SetSettingRequest) String() string            { return proto.CompactTextString(m) }
func (*SetSettingRequest) ProtoMessage()               {}
func (*SetSettingRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type isSetSettingRequest_Value interface {
	isSetSettingRequest_Value()
}

type SetSettingRequest_BoolVal struct {
	BoolVal bool `protobuf:"varint,2,opt,name=bool_val,oneof"`
}
type SetSettingRequest_StrVal struct {
	StrVal string `protobuf:"bytes,3,opt,name=str_val,oneof"`
}
type SetSettingRequest_IntVal struct {
	IntVal int32 `protobuf:"varint,4,opt,name=int_val,oneof"`
}
type SetSettingRequest_FloatVal struct {
	FloatVal float32 `protobuf:"fixed32,5,opt,name=float_val,oneof"`
}

func (*SetSettingRequest_BoolVal) isSetSettingRequest_Value()  {}
func (*SetSettingRequest_StrVal) isSetSettingRequest_Value()   {}
func (*SetSettingRequest_IntVal) isSetSettingRequest_Value()   {}
func (*SetSettingRequest_FloatVal) isSetSettingRequest_Value() {}

func (m *SetSettingRequest) GetValue() isSetSettingRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SetSettingRequest) GetBoolVal() bool {
	if x, ok := m.GetValue().(*SetSettingRequest_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (m *SetSettingRequest) GetStrVal() string {
	if x, ok := m.GetValue().(*SetSettingRequest_StrVal); ok {
		return x.StrVal
	}
	return ""
}

func (m *SetSettingRequest) GetIntVal() int32 {
	if x, ok := m.GetValue().(*SetSettingRequest_IntVal); ok {
		return x.IntVal
	}
	return 0
}

func (m *SetSettingRequest) GetFloatVal() float32 {
	if x, ok := m.GetValue().(*SetSettingRequest_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SetSettingRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SetSettingRequest_OneofMarshaler, _SetSettingRequest_OneofUnmarshaler, _SetSettingRequest_OneofSizer, []interface{}{
		(*SetSettingRequest_BoolVal)(nil),
		(*SetSettingRequest_StrVal)(nil),
		(*SetSettingRequest_IntVal)(nil),
		(*SetSettingRequest_FloatVal)(nil),
	}
}

func _SetSettingRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SetSettingRequest)
	// value
	switch x := m.Value.(type) {
	case *SetSettingRequest_BoolVal:
		t := uint64(0)
		if x.BoolVal {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *SetSettingRequest_StrVal:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StrVal)
	case *SetSettingRequest_IntVal:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.IntVal))
	case *SetSettingRequest_FloatVal:
		b.EncodeVarint(5<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatVal)))
	case nil:
	default:
		return fmt.Errorf("SetSettingRequest.Value has unexpected type %T", x)
	}
	return nil
}

func _SetSettingRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SetSettingRequest)
	switch tag {
	case 2: // value.bool_val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &SetSettingRequest_BoolVal{x != 0}
		return true, err
	case 3: // value.str_val
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &SetSettingRequest_StrVal{x}
		return true, err
	case 4: // value.int_val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &SetSettingRequest_IntVal{int32(x)}
		return true, err
	case 5: // value.float_val
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &SetSettingRequest_FloatVal{math.Float32frombits(uint32(x))}
		return true, err
	default:
		return false, nil
	}
}

func _SetSettingRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SetSettingRequest)
	// value
	switch x := m.Value.(type) {
	case *SetSettingRequest_BoolVal:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *SetSettingRequest_StrVal:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StrVal)))
		n += len(x.StrVal)
	case *SetSettingRequest_IntVal:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.IntVal))
	case *SetSettingRequest_FloatVal:
		n += proto.SizeVarint(5<<3 | proto.WireFixed32)
		n += 4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SetSettingResponse struct {
}

func (m *SetSettingResponse) Reset()                    { *m = SetSettingResponse{} }
func (m *SetSettingResponse) String() string            { return proto.CompactTextString(m) }
func (*SetSettingResponse) ProtoMessage()               {}
func (*SetSettingResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func init() {
	proto.RegisterType((*SkipStatusRequest)(nil), "CrowdSound.SkipStatusRequest")
	proto.RegisterType((*SkipStatusResponse)(nil), "CrowdSound.SkipStatusResponse")
	proto.RegisterType((*SkipRequest)(nil), "CrowdSound.SkipRequest")
	proto.RegisterType((*SkipResponse)(nil), "CrowdSound.SkipResponse")
	proto.RegisterType((*GetSettingsRequest)(nil), "CrowdSound.GetSettingsRequest")
	proto.RegisterType((*GetSettingsResponse)(nil), "CrowdSound.GetSettingsResponse")
	proto.RegisterType((*SetSettingRequest)(nil), "CrowdSound.SetSettingRequest")
	proto.RegisterType((*SetSettingResponse)(nil), "CrowdSound.SetSettingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Admin service

type AdminClient interface {
	SkipStatus(ctx context.Context, in *SkipStatusRequest, opts ...grpc.CallOption) (*SkipStatusResponse, error)
	Skip(ctx context.Context, in *SkipRequest, opts ...grpc.CallOption) (*SkipResponse, error)
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error)
	SetSetting(ctx context.Context, in *SetSettingRequest, opts ...grpc.CallOption) (*SetSettingResponse, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) SkipStatus(ctx context.Context, in *SkipStatusRequest, opts ...grpc.CallOption) (*SkipStatusResponse, error) {
	out := new(SkipStatusResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.Admin/SkipStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Skip(ctx context.Context, in *SkipRequest, opts ...grpc.CallOption) (*SkipResponse, error) {
	out := new(SkipResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.Admin/Skip", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error) {
	out := new(GetSettingsResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.Admin/GetSettings", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) SetSetting(ctx context.Context, in *SetSettingRequest, opts ...grpc.CallOption) (*SetSettingResponse, error) {
	out := new(SetSettingResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.Admin/SetSetting", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	SkipStatus(context.Context, *SkipStatusRequest) (*SkipStatusResponse, error)
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	SetSetting(context.Context, *SetSettingRequest) (*SetSettingResponse, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_SkipStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SkipStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).SkipStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_Skip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SkipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).Skip(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).GetSettings(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Admin_SetSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SetSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AdminServer).SetSetting(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CrowdSound.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SkipStatus",
			Handler:    _Admin_SkipStatus_Handler,
		},
		{
			MethodName: "Skip",
			Handler:    _Admin_Skip_Handler,
		},
		{
			MethodName: "GetSettings",
			Handler:    _Admin_GetSettings_Handler,
		},
		{
			MethodName: "SetSetting",
			Handler:    _Admin_SetSetting_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xc1, 0x72, 0x12, 0x41,
	0x10, 0x86, 0x09, 0x09, 0x01, 0x1a, 0x88, 0xc5, 0x40, 0xcc, 0x9a, 0x32, 0x6a, 0x71, 0xf2, 0x22,
	0x56, 0xe9, 0xd1, 0x93, 0xf1, 0x60, 0x2e, 0x56, 0x59, 0x72, 0xf3, 0x32, 0x35, 0xb0, 0x0d, 0x4c,
	0xb1, 0xcc, 0xac, 0x33, 0xbd, 0x20, 0x3e, 0x97, 0x2f, 0xe2, 0x1b, 0x39, 0x33, 0xcb, 0x2e, 0x1b,
	0xa8, 0xdc, 0x76, 0xbf, 0x7f, 0x76, 0xba, 0xfb, 0xff, 0x7b, 0xe1, 0x5d, 0xba, 0x5a, 0xbc, 0x9f,
	0x19, 0xbd, 0x8d, 0xad, 0xce, 0x54, 0x5c, 0x79, 0xe4, 0x22, 0x5e, 0x4b, 0xc5, 0x2d, 0x9a, 0x8d,
	0x9c, 0xe1, 0x38, 0x35, 0x9a, 0x34, 0x83, 0x2f, 0x5e, 0x9f, 0x78, 0x7d, 0x34, 0x80, 0xfe, 0x64,
	0x25, 0xd3, 0x09, 0x09, 0xca, 0xec, 0x0f, 0xfc, 0x95, 0xa1, 0xa5, 0xd1, 0x04, 0x58, 0x15, 0xda,
	0x54, 0x2b, 0x8b, 0xec, 0x1a, 0x7a, 0x1b, 0x4d, 0x68, 0x39, 0x69, 0x6e, 0x9d, 0x1c, 0x9d, 0xbd,
	0x39, 0x7b, 0xdb, 0x60, 0x03, 0xe8, 0x90, 0x26, 0x91, 0xf0, 0xcc, 0x55, 0xb1, 0x51, 0x3d, 0xc0,
	0x3e, 0xb4, 0x69, 0x69, 0xd0, 0x2e, 0x75, 0x12, 0x47, 0xe7, 0x0e, 0xd5, 0x47, 0x3d, 0xe8, 0xf8,
	0x4b, 0x8b, 0x1a, 0x57, 0xd0, 0xcd, 0x5f, 0xf3, 0xdb, 0x47, 0x43, 0x60, 0x5f, 0x91, 0x26, 0x48,
	0x24, 0xd5, 0xa2, 0xec, 0xe4, 0x5f, 0x1d, 0x06, 0x8f, 0xf0, 0xbe, 0x97, 0x1b, 0x78, 0x36, 0x97,
	0x09, 0xa1, 0xe1, 0xd3, 0x6c, 0x3e, 0x47, 0x83, 0x71, 0xe8, 0xa6, 0xc5, 0x5e, 0xc2, 0x50, 0x2a,
	0x31, 0x23, 0xb9, 0x91, 0xb4, 0xe3, 0x87, 0x1e, 0xf2, 0xb6, 0x86, 0xd0, 0x75, 0x20, 0x4b, 0x88,
	0x27, 0x72, 0x2d, 0x29, 0x74, 0x16, 0xa8, 0x45, 0x6b, 0xa5, 0x56, 0x5c, 0x89, 0x35, 0x46, 0x17,
	0x8e, 0xb6, 0x19, 0x03, 0x70, 0x3d, 0x64, 0xc8, 0xad, 0xfc, 0x83, 0x51, 0x23, 0x9c, 0xbc, 0x83,
	0x6b, 0x32, 0xa8, 0x62, 0xd7, 0x0b, 0x17, 0x86, 0xa4, 0x25, 0x9b, 0xcb, 0x97, 0x41, 0x7e, 0x0e,
	0x57, 0xde, 0x98, 0x4a, 0xd9, 0xa6, 0x1f, 0xdd, 0x17, 0x98, 0x39, 0xb7, 0x89, 0x6f, 0x51, 0x2e,
	0x96, 0x14, 0xb5, 0x02, 0x75, 0xc6, 0x79, 0x3f, 0x0b, 0xd8, 0x2e, 0x8e, 0x2e, 0x50, 0x99, 0x92,
	0x42, 0xa0, 0xce, 0xfa, 0xbc, 0x5c, 0x81, 0x3b, 0x01, 0xbf, 0x80, 0x7e, 0x9a, 0x88, 0x1d, 0xba,
	0x98, 0x17, 0xc2, 0xc5, 0xbc, 0x76, 0x93, 0x45, 0xdd, 0x42, 0xf2, 0xc1, 0x1b, 0x4c, 0x51, 0xb8,
	0xaf, 0xa4, 0x8a, 0xf5, 0x36, 0xea, 0x85, 0x20, 0x7e, 0xbb, 0xc8, 0x4b, 0x4b, 0xf7, 0x46, 0xb3,
	0x0e, 0x9c, 0xaf, 0x70, 0x17, 0x4c, 0xf4, 0xa3, 0xb7, 0xa6, 0x5a, 0x27, 0x7c, 0x23, 0x92, 0x60,
	0x5c, 0xeb, 0xa1, 0xe6, 0x12, 0x6d, 0x5a, 0x32, 0x01, 0x79, 0xd7, 0xda, 0x39, 0x92, 0x6e, 0x28,
	0x8f, 0xbc, 0x65, 0x0d, 0x87, 0x06, 0xd0, 0x9e, 0x27, 0x5a, 0xe4, 0xd0, 0x7b, 0x56, 0x7f, 0xa8,
	0xdd, 0x37, 0xa1, 0xe1, 0x5e, 0xb3, 0x90, 0x71, 0xb5, 0x72, 0x9e, 0xe5, 0x87, 0xbf, 0x75, 0x68,
	0x7c, 0xf6, 0x6b, 0xca, 0xbe, 0x01, 0x1c, 0xf6, 0x8e, 0xdd, 0x8d, 0x0f, 0x7b, 0x3a, 0x3e, 0x59,
	0xd2, 0xdb, 0x57, 0x4f, 0xc9, 0xfb, 0x85, 0xaa, 0xb1, 0x4f, 0x70, 0xe1, 0x39, 0xbb, 0x39, 0x3e,
	0x59, 0x5c, 0x11, 0x9d, 0x0a, 0xe5, 0xc7, 0xdf, 0xa1, 0x53, 0x59, 0x3c, 0xf6, 0xa8, 0xda, 0xe9,
	0xa2, 0xde, 0xbe, 0x7e, 0x52, 0x2f, 0x6f, 0xf4, 0xd3, 0x95, 0xc2, 0xd1, 0x74, 0xc7, 0x79, 0x1c,
	0x4d, 0x77, 0x62, 0xda, 0xa8, 0x76, 0xdf, 0xfd, 0x09, 0x87, 0xff, 0x7c, 0x7a, 0x19, 0x7e, 0xed,
	0x8f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xec, 0xdc, 0x84, 0x0b, 0x04, 0x00, 0x00,
}
