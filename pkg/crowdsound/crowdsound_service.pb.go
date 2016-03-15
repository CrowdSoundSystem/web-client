// Code generated by protoc-gen-go.
// source: pkg/crowdsound/crowdsound_service.proto
// DO NOT EDIT!

/*
Package crowdsound is a generated protocol buffer package.

It is generated from these files:
	pkg/crowdsound/crowdsound_service.proto
	pkg/crowdsound/crowdsound_admin_service.proto

It has these top-level messages:
	PingRequest
	PingResponse
	GetSessionDataRequest
	GetSessionDataResponse
	GetPlayingRequest
	GetPlayingResponse
	GetQueueRequest
	GetQueueResponse
	ListTrendingArtistsRequest
	ListTrendingArtistsResponse
	PostSongRequest
	PostSongResponse
	VoteSongRequest
	VoteSongResponse
	VoteSkipRequest
	VoteSkipResponse
	SkipStatusRequest
	SkipStatusResponse
	SkipRequest
	SkipResponse
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PingRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingResponse struct {
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetSessionDataRequest struct {
}

func (m *GetSessionDataRequest) Reset()                    { *m = GetSessionDataRequest{} }
func (m *GetSessionDataRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSessionDataRequest) ProtoMessage()               {}
func (*GetSessionDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetSessionDataResponse struct {
	SessionName string `protobuf:"bytes,1,opt,name=session_name" json:"session_name,omitempty"`
	NumUsers    int32  `protobuf:"varint,2,opt,name=num_users" json:"num_users,omitempty"`
}

func (m *GetSessionDataResponse) Reset()                    { *m = GetSessionDataResponse{} }
func (m *GetSessionDataResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSessionDataResponse) ProtoMessage()               {}
func (*GetSessionDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type GetPlayingRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *GetPlayingRequest) Reset()                    { *m = GetPlayingRequest{} }
func (m *GetPlayingRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPlayingRequest) ProtoMessage()               {}
func (*GetPlayingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetPlayingResponse struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Artist string `protobuf:"bytes,2,opt,name=artist" json:"artist,omitempty"`
	Genre  string `protobuf:"bytes,3,opt,name=genre" json:"genre,omitempty"`
}

func (m *GetPlayingResponse) Reset()                    { *m = GetPlayingResponse{} }
func (m *GetPlayingResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPlayingResponse) ProtoMessage()               {}
func (*GetPlayingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetQueueRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *GetQueueRequest) Reset()                    { *m = GetQueueRequest{} }
func (m *GetQueueRequest) String() string            { return proto.CompactTextString(m) }
func (*GetQueueRequest) ProtoMessage()               {}
func (*GetQueueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type GetQueueResponse struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Artist     string `protobuf:"bytes,2,opt,name=artist" json:"artist,omitempty"`
	Genre      string `protobuf:"bytes,3,opt,name=genre" json:"genre,omitempty"`
	IsPlaying  bool   `protobuf:"varint,4,opt,name=isPlaying" json:"isPlaying,omitempty"`
	IsBuffered bool   `protobuf:"varint,5,opt,name=isBuffered" json:"isBuffered,omitempty"`
}

func (m *GetQueueResponse) Reset()                    { *m = GetQueueResponse{} }
func (m *GetQueueResponse) String() string            { return proto.CompactTextString(m) }
func (*GetQueueResponse) ProtoMessage()               {}
func (*GetQueueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ListTrendingArtistsRequest struct {
}

func (m *ListTrendingArtistsRequest) Reset()                    { *m = ListTrendingArtistsRequest{} }
func (m *ListTrendingArtistsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTrendingArtistsRequest) ProtoMessage()               {}
func (*ListTrendingArtistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ListTrendingArtistsResponse struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Score int32  `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
}

func (m *ListTrendingArtistsResponse) Reset()                    { *m = ListTrendingArtistsResponse{} }
func (m *ListTrendingArtistsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTrendingArtistsResponse) ProtoMessage()               {}
func (*ListTrendingArtistsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PostSongRequest struct {
	UserId string   `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Artist []string `protobuf:"bytes,3,rep,name=artist" json:"artist,omitempty"`
	Genre  string   `protobuf:"bytes,4,opt,name=genre" json:"genre,omitempty"`
}

func (m *PostSongRequest) Reset()                    { *m = PostSongRequest{} }
func (m *PostSongRequest) String() string            { return proto.CompactTextString(m) }
func (*PostSongRequest) ProtoMessage()               {}
func (*PostSongRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type PostSongResponse struct {
}

func (m *PostSongResponse) Reset()                    { *m = PostSongResponse{} }
func (m *PostSongResponse) String() string            { return proto.CompactTextString(m) }
func (*PostSongResponse) ProtoMessage()               {}
func (*PostSongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type VoteSongRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Artist string `protobuf:"bytes,3,opt,name=artist" json:"artist,omitempty"`
	Like   bool   `protobuf:"varint,4,opt,name=like" json:"like,omitempty"`
}

func (m *VoteSongRequest) Reset()                    { *m = VoteSongRequest{} }
func (m *VoteSongRequest) String() string            { return proto.CompactTextString(m) }
func (*VoteSongRequest) ProtoMessage()               {}
func (*VoteSongRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type VoteSongResponse struct {
}

func (m *VoteSongResponse) Reset()                    { *m = VoteSongResponse{} }
func (m *VoteSongResponse) String() string            { return proto.CompactTextString(m) }
func (*VoteSongResponse) ProtoMessage()               {}
func (*VoteSongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type VoteSkipRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *VoteSkipRequest) Reset()                    { *m = VoteSkipRequest{} }
func (m *VoteSkipRequest) String() string            { return proto.CompactTextString(m) }
func (*VoteSkipRequest) ProtoMessage()               {}
func (*VoteSkipRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type VoteSkipResponse struct {
}

func (m *VoteSkipResponse) Reset()                    { *m = VoteSkipResponse{} }
func (m *VoteSkipResponse) String() string            { return proto.CompactTextString(m) }
func (*VoteSkipResponse) ProtoMessage()               {}
func (*VoteSkipResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func init() {
	proto.RegisterType((*PingRequest)(nil), "CrowdSound.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "CrowdSound.PingResponse")
	proto.RegisterType((*GetSessionDataRequest)(nil), "CrowdSound.GetSessionDataRequest")
	proto.RegisterType((*GetSessionDataResponse)(nil), "CrowdSound.GetSessionDataResponse")
	proto.RegisterType((*GetPlayingRequest)(nil), "CrowdSound.GetPlayingRequest")
	proto.RegisterType((*GetPlayingResponse)(nil), "CrowdSound.GetPlayingResponse")
	proto.RegisterType((*GetQueueRequest)(nil), "CrowdSound.GetQueueRequest")
	proto.RegisterType((*GetQueueResponse)(nil), "CrowdSound.GetQueueResponse")
	proto.RegisterType((*ListTrendingArtistsRequest)(nil), "CrowdSound.ListTrendingArtistsRequest")
	proto.RegisterType((*ListTrendingArtistsResponse)(nil), "CrowdSound.ListTrendingArtistsResponse")
	proto.RegisterType((*PostSongRequest)(nil), "CrowdSound.PostSongRequest")
	proto.RegisterType((*PostSongResponse)(nil), "CrowdSound.PostSongResponse")
	proto.RegisterType((*VoteSongRequest)(nil), "CrowdSound.VoteSongRequest")
	proto.RegisterType((*VoteSongResponse)(nil), "CrowdSound.VoteSongResponse")
	proto.RegisterType((*VoteSkipRequest)(nil), "CrowdSound.VoteSkipRequest")
	proto.RegisterType((*VoteSkipResponse)(nil), "CrowdSound.VoteSkipResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for CrowdSound service

type CrowdSoundClient interface {
	// Ping sends a ping to the server to indicate the client is
	// alive and connected.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	GetSessionData(ctx context.Context, in *GetSessionDataRequest, opts ...grpc.CallOption) (*GetSessionDataResponse, error)
	// GetPlaying returns the current song that is playing.
	GetPlaying(ctx context.Context, in *GetPlayingRequest, opts ...grpc.CallOption) (*GetPlayingResponse, error)
	// GetQueue streams the queue of songs.
	GetQueue(ctx context.Context, in *GetQueueRequest, opts ...grpc.CallOption) (CrowdSound_GetQueueClient, error)
	// ListTrendingArtists returns an ordered list of artists, based on trending
	// value.
	//
	// As with ListSongs, it is up to the client to control pagination and stream
	// termination (with the exception of end-of-stream).
	ListTrendingArtists(ctx context.Context, in *ListTrendingArtistsRequest, opts ...grpc.CallOption) (CrowdSound_ListTrendingArtistsClient, error)
	// PostSong informs the CrowdSound service of a 'Song' that the client
	// has. The 'Song' may or may not have all of the fields filled or present.
	PostSong(ctx context.Context, opts ...grpc.CallOption) (CrowdSound_PostSongClient, error)
	// VoteSong informs the CrowdSound service of a vote for a Song.
	VoteSong(ctx context.Context, in *VoteSongRequest, opts ...grpc.CallOption) (*VoteSongResponse, error)
	// VoteSkip votes to skip the currently playing song.
	VoteSkip(ctx context.Context, in *VoteSkipRequest, opts ...grpc.CallOption) (*VoteSkipResponse, error)
}

type crowdSoundClient struct {
	cc *grpc.ClientConn
}

func NewCrowdSoundClient(cc *grpc.ClientConn) CrowdSoundClient {
	return &crowdSoundClient{cc}
}

func (c *crowdSoundClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.CrowdSound/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowdSoundClient) GetSessionData(ctx context.Context, in *GetSessionDataRequest, opts ...grpc.CallOption) (*GetSessionDataResponse, error) {
	out := new(GetSessionDataResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.CrowdSound/GetSessionData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowdSoundClient) GetPlaying(ctx context.Context, in *GetPlayingRequest, opts ...grpc.CallOption) (*GetPlayingResponse, error) {
	out := new(GetPlayingResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.CrowdSound/GetPlaying", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowdSoundClient) GetQueue(ctx context.Context, in *GetQueueRequest, opts ...grpc.CallOption) (CrowdSound_GetQueueClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CrowdSound_serviceDesc.Streams[0], c.cc, "/CrowdSound.CrowdSound/GetQueue", opts...)
	if err != nil {
		return nil, err
	}
	x := &crowdSoundGetQueueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrowdSound_GetQueueClient interface {
	Recv() (*GetQueueResponse, error)
	grpc.ClientStream
}

type crowdSoundGetQueueClient struct {
	grpc.ClientStream
}

func (x *crowdSoundGetQueueClient) Recv() (*GetQueueResponse, error) {
	m := new(GetQueueResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crowdSoundClient) ListTrendingArtists(ctx context.Context, in *ListTrendingArtistsRequest, opts ...grpc.CallOption) (CrowdSound_ListTrendingArtistsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CrowdSound_serviceDesc.Streams[1], c.cc, "/CrowdSound.CrowdSound/ListTrendingArtists", opts...)
	if err != nil {
		return nil, err
	}
	x := &crowdSoundListTrendingArtistsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CrowdSound_ListTrendingArtistsClient interface {
	Recv() (*ListTrendingArtistsResponse, error)
	grpc.ClientStream
}

type crowdSoundListTrendingArtistsClient struct {
	grpc.ClientStream
}

func (x *crowdSoundListTrendingArtistsClient) Recv() (*ListTrendingArtistsResponse, error) {
	m := new(ListTrendingArtistsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crowdSoundClient) PostSong(ctx context.Context, opts ...grpc.CallOption) (CrowdSound_PostSongClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_CrowdSound_serviceDesc.Streams[2], c.cc, "/CrowdSound.CrowdSound/PostSong", opts...)
	if err != nil {
		return nil, err
	}
	x := &crowdSoundPostSongClient{stream}
	return x, nil
}

type CrowdSound_PostSongClient interface {
	Send(*PostSongRequest) error
	CloseAndRecv() (*PostSongResponse, error)
	grpc.ClientStream
}

type crowdSoundPostSongClient struct {
	grpc.ClientStream
}

func (x *crowdSoundPostSongClient) Send(m *PostSongRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *crowdSoundPostSongClient) CloseAndRecv() (*PostSongResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PostSongResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *crowdSoundClient) VoteSong(ctx context.Context, in *VoteSongRequest, opts ...grpc.CallOption) (*VoteSongResponse, error) {
	out := new(VoteSongResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.CrowdSound/VoteSong", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crowdSoundClient) VoteSkip(ctx context.Context, in *VoteSkipRequest, opts ...grpc.CallOption) (*VoteSkipResponse, error) {
	out := new(VoteSkipResponse)
	err := grpc.Invoke(ctx, "/CrowdSound.CrowdSound/VoteSkip", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CrowdSound service

type CrowdSoundServer interface {
	// Ping sends a ping to the server to indicate the client is
	// alive and connected.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	GetSessionData(context.Context, *GetSessionDataRequest) (*GetSessionDataResponse, error)
	// GetPlaying returns the current song that is playing.
	GetPlaying(context.Context, *GetPlayingRequest) (*GetPlayingResponse, error)
	// GetQueue streams the queue of songs.
	GetQueue(*GetQueueRequest, CrowdSound_GetQueueServer) error
	// ListTrendingArtists returns an ordered list of artists, based on trending
	// value.
	//
	// As with ListSongs, it is up to the client to control pagination and stream
	// termination (with the exception of end-of-stream).
	ListTrendingArtists(*ListTrendingArtistsRequest, CrowdSound_ListTrendingArtistsServer) error
	// PostSong informs the CrowdSound service of a 'Song' that the client
	// has. The 'Song' may or may not have all of the fields filled or present.
	PostSong(CrowdSound_PostSongServer) error
	// VoteSong informs the CrowdSound service of a vote for a Song.
	VoteSong(context.Context, *VoteSongRequest) (*VoteSongResponse, error)
	// VoteSkip votes to skip the currently playing song.
	VoteSkip(context.Context, *VoteSkipRequest) (*VoteSkipResponse, error)
}

func RegisterCrowdSoundServer(s *grpc.Server, srv CrowdSoundServer) {
	s.RegisterService(&_CrowdSound_serviceDesc, srv)
}

func _CrowdSound_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CrowdSoundServer).Ping(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CrowdSound_GetSessionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetSessionDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CrowdSoundServer).GetSessionData(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CrowdSound_GetPlaying_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetPlayingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CrowdSoundServer).GetPlaying(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CrowdSound_GetQueue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetQueueRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrowdSoundServer).GetQueue(m, &crowdSoundGetQueueServer{stream})
}

type CrowdSound_GetQueueServer interface {
	Send(*GetQueueResponse) error
	grpc.ServerStream
}

type crowdSoundGetQueueServer struct {
	grpc.ServerStream
}

func (x *crowdSoundGetQueueServer) Send(m *GetQueueResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrowdSound_ListTrendingArtists_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTrendingArtistsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CrowdSoundServer).ListTrendingArtists(m, &crowdSoundListTrendingArtistsServer{stream})
}

type CrowdSound_ListTrendingArtistsServer interface {
	Send(*ListTrendingArtistsResponse) error
	grpc.ServerStream
}

type crowdSoundListTrendingArtistsServer struct {
	grpc.ServerStream
}

func (x *crowdSoundListTrendingArtistsServer) Send(m *ListTrendingArtistsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CrowdSound_PostSong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CrowdSoundServer).PostSong(&crowdSoundPostSongServer{stream})
}

type CrowdSound_PostSongServer interface {
	SendAndClose(*PostSongResponse) error
	Recv() (*PostSongRequest, error)
	grpc.ServerStream
}

type crowdSoundPostSongServer struct {
	grpc.ServerStream
}

func (x *crowdSoundPostSongServer) SendAndClose(m *PostSongResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *crowdSoundPostSongServer) Recv() (*PostSongRequest, error) {
	m := new(PostSongRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CrowdSound_VoteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(VoteSongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CrowdSoundServer).VoteSong(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CrowdSound_VoteSkip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(VoteSkipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CrowdSoundServer).VoteSkip(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CrowdSound_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CrowdSound.CrowdSound",
	HandlerType: (*CrowdSoundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CrowdSound_Ping_Handler,
		},
		{
			MethodName: "GetSessionData",
			Handler:    _CrowdSound_GetSessionData_Handler,
		},
		{
			MethodName: "GetPlaying",
			Handler:    _CrowdSound_GetPlaying_Handler,
		},
		{
			MethodName: "VoteSong",
			Handler:    _CrowdSound_VoteSong_Handler,
		},
		{
			MethodName: "VoteSkip",
			Handler:    _CrowdSound_VoteSkip_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQueue",
			Handler:       _CrowdSound_GetQueue_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListTrendingArtists",
			Handler:       _CrowdSound_ListTrendingArtists_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostSong",
			Handler:       _CrowdSound_PostSong_Handler,
			ClientStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x73, 0x2c, 0xc9, 0x10, 0x92, 0x66, 0x38, 0xd4, 0x72, 0x4b, 0x55, 0x56, 0x88, 0xf6,
	0x2a, 0x20, 0xb8, 0x83, 0xab, 0xb4, 0x48, 0x15, 0x12, 0x87, 0x40, 0x10, 0x12, 0xdc, 0x44, 0x26,
	0x9e, 0x44, 0x4b, 0x5a, 0x6f, 0xf0, 0xda, 0x20, 0x9e, 0x80, 0xf7, 0xe0, 0x49, 0x59, 0xdb, 0xeb,
	0x78, 0xed, 0xd8, 0xad, 0xd4, 0xbb, 0x68, 0x67, 0xe6, 0xfb, 0xc7, 0x9b, 0xcf, 0x86, 0xe3, 0xf5,
	0x6a, 0xf9, 0x74, 0xee, 0x8b, 0xdf, 0xae, 0x14, 0xa1, 0xe7, 0x1a, 0x3f, 0x67, 0x92, 0xfc, 0x5f,
	0x7c, 0x4e, 0xa3, 0xb5, 0x2f, 0x02, 0x81, 0x70, 0x16, 0x55, 0xa6, 0x51, 0x85, 0x1d, 0xc2, 0xed,
	0x09, 0xf7, 0x96, 0x9f, 0xe8, 0x67, 0x48, 0x32, 0xc0, 0x01, 0xdc, 0x0a, 0x55, 0xf3, 0x8c, 0xbb,
	0x56, 0xfd, 0xa8, 0x7e, 0xd2, 0x65, 0x7d, 0xe8, 0x25, 0x75, 0xb9, 0x16, 0x9e, 0x24, 0xb6, 0x07,
	0xf7, 0xcf, 0x29, 0x98, 0x92, 0x94, 0x5c, 0x78, 0xaf, 0x9d, 0xc0, 0xd1, 0x93, 0x6c, 0x0c, 0x0f,
	0x8a, 0x85, 0x64, 0x04, 0xef, 0x41, 0x4f, 0x26, 0xc7, 0x33, 0xcf, 0xb9, 0xa4, 0x04, 0x8c, 0x43,
	0xe8, 0x7a, 0xe1, 0xe5, 0x2c, 0x4a, 0x93, 0x56, 0x43, 0x1d, 0xb5, 0xd9, 0x63, 0x18, 0x2a, 0xc4,
	0xe4, 0xc2, 0xf9, 0x73, 0xd5, 0x46, 0x63, 0x40, 0xb3, 0x4b, 0x87, 0xf4, 0xa0, 0x65, 0xc0, 0xfb,
	0xb0, 0xe3, 0xf8, 0x01, 0x97, 0x41, 0x4c, 0xee, 0xe2, 0x1d, 0x68, 0x2f, 0xc9, 0xf3, 0xc9, 0x6a,
	0xc6, 0x08, 0x06, 0x03, 0x85, 0xf8, 0x18, 0x52, 0x48, 0x95, 0x31, 0x0b, 0xd8, 0xcd, 0x7a, 0x6e,
	0x10, 0x12, 0x3d, 0x20, 0x97, 0x7a, 0x4d, 0xab, 0xa5, 0x8e, 0x3a, 0x88, 0x00, 0x5c, 0x9e, 0x86,
	0x8b, 0x05, 0xf9, 0xe4, 0x5a, 0xed, 0xe8, 0x8c, 0x1d, 0x80, 0xfd, 0x56, 0x31, 0x3e, 0xfb, 0xe4,
	0xb9, 0xaa, 0x73, 0x1c, 0x13, 0x65, 0x7a, 0xab, 0x2f, 0x61, 0xbf, 0xb4, 0x5a, 0xba, 0x90, 0x5a,
	0x40, 0xce, 0x85, 0x5a, 0x20, 0xb9, 0xce, 0x0f, 0x30, 0x98, 0x08, 0x19, 0x4c, 0x45, 0xf5, 0x65,
	0x6e, 0x00, 0x8d, 0xc2, 0x13, 0x35, 0x8f, 0x9a, 0xe6, 0x13, 0xb5, 0xe2, 0x2b, 0x41, 0xd8, 0xcd,
	0x80, 0xda, 0x87, 0xf7, 0x30, 0xf8, 0x22, 0x02, 0xba, 0x61, 0x88, 0xae, 0x5e, 0xf0, 0x55, 0x92,
	0xd1, 0x89, 0x32, 0x32, 0x9e, 0xce, 0x60, 0x3a, 0x63, 0xc5, 0xd7, 0x95, 0x7f, 0x57, 0x3a, 0x17,
	0xf7, 0x24, 0x73, 0xcf, 0xff, 0xb6, 0xc1, 0x50, 0x1d, 0x5f, 0x41, 0x2b, 0x52, 0x19, 0xf7, 0x46,
	0xd9, 0xe1, 0xc8, 0x90, 0xdf, 0xb6, 0xb6, 0x0b, 0x7a, 0x83, 0x1a, 0x7e, 0x85, 0x7e, 0x5e, 0x6f,
	0x7c, 0x64, 0x76, 0x97, 0xbe, 0x13, 0x36, 0xbb, 0xaa, 0x65, 0x83, 0x7e, 0x07, 0x90, 0x09, 0x8d,
	0x0f, 0x0b, 0x33, 0xf9, 0xd7, 0xc1, 0x3e, 0xac, 0x2a, 0x6f, 0x70, 0x6f, 0xa0, 0x93, 0x8a, 0x8b,
	0xfb, 0x85, 0x6e, 0x53, 0x79, 0xfb, 0xa0, 0xbc, 0x98, 0x82, 0x9e, 0xd5, 0xf1, 0x07, 0xdc, 0x2d,
	0xb1, 0x0f, 0x9f, 0x98, 0x83, 0xd5, 0xf2, 0xda, 0xc7, 0xd7, 0xf6, 0x19, 0x59, 0x6a, 0xed, 0x54,
	0xae, 0xfc, 0xda, 0x05, 0x87, 0xf3, 0x6b, 0x6f, 0xf9, 0x58, 0x3b, 0xa9, 0xe3, 0x39, 0x74, 0x52,
	0x87, 0xf2, 0xa8, 0x82, 0xa9, 0x79, 0xd4, 0x96, 0x76, 0xb5, 0x0d, 0x48, 0x49, 0x55, 0x02, 0xca,
	0x74, 0x2c, 0x01, 0x19, 0x1e, 0xb2, 0xda, 0xe9, 0xf0, 0x1b, 0x64, 0x5f, 0xe3, 0x7f, 0x8d, 0xc6,
	0xd9, 0xf4, 0xfb, 0x4e, 0xfc, 0x2d, 0x7e, 0xf1, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x79, 0xf8, 0x72,
	0x28, 0xb6, 0x05, 0x00, 0x00,
}
