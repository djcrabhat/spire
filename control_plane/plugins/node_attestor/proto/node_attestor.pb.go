// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_attestor.proto

/*
Package control_plane_proto is a generated protocol buffer package.

It is generated from these files:
	node_attestor.proto

It has these top-level messages:
	AttestedData
	AttestRequest
	AttestResponse
*/
package control_plane_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import control_plane_proto1 "github.com/spiffe/sri/control_plane/plugins/common/proto"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureRequest control_plane_proto1.ConfigureRequest

func (m *ConfigureRequest) Reset() { (*control_plane_proto1.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string {
	return (*control_plane_proto1.ConfigureRequest)(m).String()
}
func (*ConfigureRequest) ProtoMessage() {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*control_plane_proto1.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type ConfigureResponse control_plane_proto1.ConfigureResponse

func (m *ConfigureResponse) Reset() { (*control_plane_proto1.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string {
	return (*control_plane_proto1.ConfigureResponse)(m).String()
}
func (*ConfigureResponse) ProtoMessage() {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*control_plane_proto1.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoRequest control_plane_proto1.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset() { (*control_plane_proto1.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string {
	return (*control_plane_proto1.GetPluginInfoRequest)(m).String()
}
func (*GetPluginInfoRequest) ProtoMessage() {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/control_plane/plugins/common/proto/common.proto
type GetPluginInfoResponse control_plane_proto1.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*control_plane_proto1.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*control_plane_proto1.GetPluginInfoResponse)(m).GetCompany()
}

// *A type which contains attestation data for specific platform.
type AttestedData struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AttestedData) Reset()                    { *m = AttestedData{} }
func (m *AttestedData) String() string            { return proto.CompactTextString(m) }
func (*AttestedData) ProtoMessage()               {}
func (*AttestedData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AttestedData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AttestedData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// *Represents a request to attest a node.
type AttestRequest struct {
	AttestedData   *AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	AttestedBefore bool          `protobuf:"varint,2,opt,name=attestedBefore" json:"attestedBefore,omitempty"`
}

func (m *AttestRequest) Reset()                    { *m = AttestRequest{} }
func (m *AttestRequest) String() string            { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()               {}
func (*AttestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttestRequest) GetAttestedData() *AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *AttestRequest) GetAttestedBefore() bool {
	if m != nil {
		return m.AttestedBefore
	}
	return false
}

// *Represents a response when attesting a node.
type AttestResponse struct {
	Valid        bool   `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	BaseSPIFFEID string `protobuf:"bytes,2,opt,name=baseSPIFFEID" json:"baseSPIFFEID,omitempty"`
}

func (m *AttestResponse) Reset()                    { *m = AttestResponse{} }
func (m *AttestResponse) String() string            { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()               {}
func (*AttestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AttestResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *AttestResponse) GetBaseSPIFFEID() string {
	if m != nil {
		return m.BaseSPIFFEID
	}
	return ""
}

func init() {
	proto.RegisterType((*AttestedData)(nil), "control_plane_proto.AttestedData")
	proto.RegisterType((*AttestRequest)(nil), "control_plane_proto.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "control_plane_proto.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	// *Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *control_plane_proto1.ConfigureRequest, opts ...grpc.CallOption) (*control_plane_proto1.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *control_plane_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*control_plane_proto1.GetPluginInfoResponse, error)
	// *Attesta a node.
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *control_plane_proto1.ConfigureRequest, opts ...grpc.CallOption) (*control_plane_proto1.ConfigureResponse, error) {
	out := new(control_plane_proto1.ConfigureResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *control_plane_proto1.GetPluginInfoRequest, opts ...grpc.CallOption) (*control_plane_proto1.GetPluginInfoResponse, error) {
	out := new(control_plane_proto1.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.NodeAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := grpc.Invoke(ctx, "/control_plane_proto.NodeAttestor/Attest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	// *Responsible for configuration of the plugin.
	Configure(context.Context, *control_plane_proto1.ConfigureRequest) (*control_plane_proto1.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *control_plane_proto1.GetPluginInfoRequest) (*control_plane_proto1.GetPluginInfoResponse, error)
	// *Attesta a node.
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(control_plane_proto1.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*control_plane_proto1.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(control_plane_proto1.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*control_plane_proto1.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/control_plane_proto.NodeAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "control_plane_proto.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
		{
			MethodName: "Attest",
			Handler:    _NodeAttestor_Attest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_attestor.proto",
}

func init() { proto.RegisterFile("node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x16, 0xa2, 0x04, 0xc6, 0xc2, 0x61, 0xf1, 0x40, 0x38, 0x61, 0x8d, 0x04, 0x3d, 0xb4, 0x09,
	0x26, 0xde, 0x51, 0xc0, 0xe0, 0xc1, 0xe0, 0x7a, 0xf3, 0x42, 0x16, 0x3a, 0xc5, 0x26, 0x65, 0xa7,
	0xee, 0x6e, 0x4d, 0xbc, 0xf8, 0x36, 0xbe, 0xa7, 0x71, 0xdb, 0x1a, 0x6a, 0x08, 0x7a, 0x9b, 0xf9,
	0xf6, 0xfb, 0x99, 0xd9, 0x81, 0xb6, 0xa4, 0x00, 0x17, 0xc2, 0x18, 0xd4, 0x86, 0x94, 0x97, 0x28,
	0x32, 0xc4, 0xda, 0x2b, 0x92, 0x46, 0x51, 0xbc, 0x48, 0x62, 0x21, 0x71, 0x61, 0xc1, 0xee, 0x64,
	0x1d, 0x99, 0x97, 0x74, 0xe9, 0xad, 0x68, 0xe3, 0xeb, 0x24, 0x0a, 0x43, 0xf4, 0xb5, 0x8a, 0xfc,
	0x12, 0xd5, 0x4f, 0xe2, 0x74, 0x1d, 0x49, 0xed, 0xaf, 0x68, 0xb3, 0x21, 0xe9, 0x5b, 0x65, 0xde,
	0x64, 0xde, 0xee, 0x35, 0x38, 0x23, 0x9b, 0x86, 0xc1, 0x58, 0x18, 0xc1, 0x18, 0x1c, 0x9a, 0xf7,
	0x04, 0x3b, 0x95, 0x5e, 0x65, 0xd0, 0xe0, 0xb6, 0xfe, 0xc6, 0x02, 0x61, 0x44, 0xa7, 0xda, 0xab,
	0x0c, 0x1c, 0x6e, 0x6b, 0xf7, 0x03, 0x9a, 0x99, 0x8e, 0xe3, 0x6b, 0x8a, 0xda, 0xb0, 0x09, 0x38,
	0x62, 0xcb, 0xc8, 0x1a, 0x1c, 0x0f, 0x4f, 0xbd, 0x1d, 0xb3, 0x7b, 0xdb, 0x89, 0xbc, 0x24, 0x63,
	0x7d, 0x68, 0x15, 0xfd, 0x0d, 0x86, 0xa4, 0xd0, 0xa6, 0xd6, 0xf9, 0x2f, 0xd4, 0xbd, 0x87, 0x56,
	0x91, 0xaf, 0x13, 0x92, 0x1a, 0xd9, 0x09, 0x1c, 0xbd, 0x89, 0x38, 0x0a, 0x6c, 0x72, 0x9d, 0x67,
	0x0d, 0x73, 0xc1, 0x59, 0x0a, 0x8d, 0x4f, 0xf3, 0xd9, 0x74, 0x3a, 0x99, 0x8d, 0xad, 0x5b, 0x83,
	0x97, 0xb0, 0xe1, 0x67, 0x15, 0x9c, 0x07, 0x0a, 0x70, 0x94, 0x7f, 0x3b, 0x7b, 0x86, 0xc6, 0x2d,
	0xc9, 0x30, 0x5a, 0xa7, 0x0a, 0xd9, 0xf9, 0xce, 0x15, 0x7e, 0xde, 0xf3, 0xfd, 0xbb, 0xfd, 0xbf,
	0x68, 0xf9, 0x98, 0x21, 0x34, 0xef, 0xd0, 0xcc, 0xed, 0x61, 0x66, 0x32, 0x24, 0x76, 0xb1, 0x53,
	0x58, 0xe2, 0x14, 0x19, 0x97, 0xff, 0xa1, 0xe6, 0x39, 0x8f, 0x50, 0xcb, 0xf6, 0x61, 0xee, 0x9e,
	0x1b, 0x14, 0xce, 0x67, 0x7b, 0x39, 0x99, 0xe5, 0xfc, 0x60, 0x59, 0xb3, 0xf8, 0xd5, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x03, 0x46, 0x43, 0x64, 0xa7, 0x02, 0x00, 0x00,
}