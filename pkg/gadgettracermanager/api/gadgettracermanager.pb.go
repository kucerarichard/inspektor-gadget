// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gadgettracermanager.proto

/*
Package gadgettracermanager is a generated protocol buffer package.

It is generated from these files:
	gadgettracermanager.proto

It has these top-level messages:
	Label
	AddTracerRequest
	RemoveTracerResponse
	AddContainerResponse
	RemoveContainerResponse
	ContainerSelector
	TracerID
	ContainerDefinition
	DumpStateRequest
	Dump
*/
package gadgettracermanager

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Label struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Label) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AddTracerRequest struct {
	Id       string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Selector *ContainerSelector `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
}

func (m *AddTracerRequest) Reset()                    { *m = AddTracerRequest{} }
func (m *AddTracerRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTracerRequest) ProtoMessage()               {}
func (*AddTracerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddTracerRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddTracerRequest) GetSelector() *ContainerSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

type RemoveTracerResponse struct {
	Debug string `protobuf:"bytes,1,opt,name=debug" json:"debug,omitempty"`
}

func (m *RemoveTracerResponse) Reset()                    { *m = RemoveTracerResponse{} }
func (m *RemoveTracerResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveTracerResponse) ProtoMessage()               {}
func (*RemoveTracerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemoveTracerResponse) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
}

type AddContainerResponse struct {
	Debug string `protobuf:"bytes,1,opt,name=debug" json:"debug,omitempty"`
}

func (m *AddContainerResponse) Reset()                    { *m = AddContainerResponse{} }
func (m *AddContainerResponse) String() string            { return proto.CompactTextString(m) }
func (*AddContainerResponse) ProtoMessage()               {}
func (*AddContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddContainerResponse) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
}

type RemoveContainerResponse struct {
	Debug string `protobuf:"bytes,1,opt,name=debug" json:"debug,omitempty"`
}

func (m *RemoveContainerResponse) Reset()                    { *m = RemoveContainerResponse{} }
func (m *RemoveContainerResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveContainerResponse) ProtoMessage()               {}
func (*RemoveContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RemoveContainerResponse) GetDebug() string {
	if m != nil {
		return m.Debug
	}
	return ""
}

type ContainerSelector struct {
	Namespace      string   `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Podname        string   `protobuf:"bytes,2,opt,name=podname" json:"podname,omitempty"`
	Labels         []*Label `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty"`
	ContainerIndex int32    `protobuf:"varint,4,opt,name=container_index,json=containerIndex" json:"container_index,omitempty"`
}

func (m *ContainerSelector) Reset()                    { *m = ContainerSelector{} }
func (m *ContainerSelector) String() string            { return proto.CompactTextString(m) }
func (*ContainerSelector) ProtoMessage()               {}
func (*ContainerSelector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ContainerSelector) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ContainerSelector) GetPodname() string {
	if m != nil {
		return m.Podname
	}
	return ""
}

func (m *ContainerSelector) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ContainerSelector) GetContainerIndex() int32 {
	if m != nil {
		return m.ContainerIndex
	}
	return 0
}

type TracerID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TracerID) Reset()                    { *m = TracerID{} }
func (m *TracerID) String() string            { return proto.CompactTextString(m) }
func (*TracerID) ProtoMessage()               {}
func (*TracerID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TracerID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ContainerDefinition struct {
	ContainerId    string   `protobuf:"bytes,1,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	CgroupPath     string   `protobuf:"bytes,2,opt,name=cgroup_path,json=cgroupPath" json:"cgroup_path,omitempty"`
	CgroupId       uint64   `protobuf:"varint,3,opt,name=cgroup_id,json=cgroupId" json:"cgroup_id,omitempty"`
	Mntns          uint64   `protobuf:"varint,4,opt,name=mntns" json:"mntns,omitempty"`
	Namespace      string   `protobuf:"bytes,5,opt,name=namespace" json:"namespace,omitempty"`
	Podname        string   `protobuf:"bytes,6,opt,name=podname" json:"podname,omitempty"`
	ContainerIndex int32    `protobuf:"varint,7,opt,name=container_index,json=containerIndex" json:"container_index,omitempty"`
	Labels         []*Label `protobuf:"bytes,8,rep,name=labels" json:"labels,omitempty"`
}

func (m *ContainerDefinition) Reset()                    { *m = ContainerDefinition{} }
func (m *ContainerDefinition) String() string            { return proto.CompactTextString(m) }
func (*ContainerDefinition) ProtoMessage()               {}
func (*ContainerDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ContainerDefinition) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *ContainerDefinition) GetCgroupPath() string {
	if m != nil {
		return m.CgroupPath
	}
	return ""
}

func (m *ContainerDefinition) GetCgroupId() uint64 {
	if m != nil {
		return m.CgroupId
	}
	return 0
}

func (m *ContainerDefinition) GetMntns() uint64 {
	if m != nil {
		return m.Mntns
	}
	return 0
}

func (m *ContainerDefinition) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ContainerDefinition) GetPodname() string {
	if m != nil {
		return m.Podname
	}
	return ""
}

func (m *ContainerDefinition) GetContainerIndex() int32 {
	if m != nil {
		return m.ContainerIndex
	}
	return 0
}

func (m *ContainerDefinition) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type DumpStateRequest struct {
}

func (m *DumpStateRequest) Reset()                    { *m = DumpStateRequest{} }
func (m *DumpStateRequest) String() string            { return proto.CompactTextString(m) }
func (*DumpStateRequest) ProtoMessage()               {}
func (*DumpStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Dump struct {
	State string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
}

func (m *Dump) Reset()                    { *m = Dump{} }
func (m *Dump) String() string            { return proto.CompactTextString(m) }
func (*Dump) ProtoMessage()               {}
func (*Dump) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Dump) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Label)(nil), "gadgettracermanager.Label")
	proto.RegisterType((*AddTracerRequest)(nil), "gadgettracermanager.AddTracerRequest")
	proto.RegisterType((*RemoveTracerResponse)(nil), "gadgettracermanager.RemoveTracerResponse")
	proto.RegisterType((*AddContainerResponse)(nil), "gadgettracermanager.AddContainerResponse")
	proto.RegisterType((*RemoveContainerResponse)(nil), "gadgettracermanager.RemoveContainerResponse")
	proto.RegisterType((*ContainerSelector)(nil), "gadgettracermanager.ContainerSelector")
	proto.RegisterType((*TracerID)(nil), "gadgettracermanager.TracerID")
	proto.RegisterType((*ContainerDefinition)(nil), "gadgettracermanager.ContainerDefinition")
	proto.RegisterType((*DumpStateRequest)(nil), "gadgettracermanager.DumpStateRequest")
	proto.RegisterType((*Dump)(nil), "gadgettracermanager.Dump")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GadgetTracerManager service

type GadgetTracerManagerClient interface {
	AddTracer(ctx context.Context, in *AddTracerRequest, opts ...grpc.CallOption) (*TracerID, error)
	RemoveTracer(ctx context.Context, in *TracerID, opts ...grpc.CallOption) (*RemoveTracerResponse, error)
	AddContainer(ctx context.Context, in *ContainerDefinition, opts ...grpc.CallOption) (*AddContainerResponse, error)
	RemoveContainer(ctx context.Context, in *ContainerDefinition, opts ...grpc.CallOption) (*RemoveContainerResponse, error)
	DumpState(ctx context.Context, in *DumpStateRequest, opts ...grpc.CallOption) (*Dump, error)
}

type gadgetTracerManagerClient struct {
	cc *grpc.ClientConn
}

func NewGadgetTracerManagerClient(cc *grpc.ClientConn) GadgetTracerManagerClient {
	return &gadgetTracerManagerClient{cc}
}

func (c *gadgetTracerManagerClient) AddTracer(ctx context.Context, in *AddTracerRequest, opts ...grpc.CallOption) (*TracerID, error) {
	out := new(TracerID)
	err := grpc.Invoke(ctx, "/gadgettracermanager.GadgetTracerManager/AddTracer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetTracerManagerClient) RemoveTracer(ctx context.Context, in *TracerID, opts ...grpc.CallOption) (*RemoveTracerResponse, error) {
	out := new(RemoveTracerResponse)
	err := grpc.Invoke(ctx, "/gadgettracermanager.GadgetTracerManager/RemoveTracer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetTracerManagerClient) AddContainer(ctx context.Context, in *ContainerDefinition, opts ...grpc.CallOption) (*AddContainerResponse, error) {
	out := new(AddContainerResponse)
	err := grpc.Invoke(ctx, "/gadgettracermanager.GadgetTracerManager/AddContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetTracerManagerClient) RemoveContainer(ctx context.Context, in *ContainerDefinition, opts ...grpc.CallOption) (*RemoveContainerResponse, error) {
	out := new(RemoveContainerResponse)
	err := grpc.Invoke(ctx, "/gadgettracermanager.GadgetTracerManager/RemoveContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetTracerManagerClient) DumpState(ctx context.Context, in *DumpStateRequest, opts ...grpc.CallOption) (*Dump, error) {
	out := new(Dump)
	err := grpc.Invoke(ctx, "/gadgettracermanager.GadgetTracerManager/DumpState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GadgetTracerManager service

type GadgetTracerManagerServer interface {
	AddTracer(context.Context, *AddTracerRequest) (*TracerID, error)
	RemoveTracer(context.Context, *TracerID) (*RemoveTracerResponse, error)
	AddContainer(context.Context, *ContainerDefinition) (*AddContainerResponse, error)
	RemoveContainer(context.Context, *ContainerDefinition) (*RemoveContainerResponse, error)
	DumpState(context.Context, *DumpStateRequest) (*Dump, error)
}

func RegisterGadgetTracerManagerServer(s *grpc.Server, srv GadgetTracerManagerServer) {
	s.RegisterService(&_GadgetTracerManager_serviceDesc, srv)
}

func _GadgetTracerManager_AddTracer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTracerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetTracerManagerServer).AddTracer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gadgettracermanager.GadgetTracerManager/AddTracer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetTracerManagerServer).AddTracer(ctx, req.(*AddTracerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetTracerManager_RemoveTracer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TracerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetTracerManagerServer).RemoveTracer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gadgettracermanager.GadgetTracerManager/RemoveTracer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetTracerManagerServer).RemoveTracer(ctx, req.(*TracerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetTracerManager_AddContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetTracerManagerServer).AddContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gadgettracermanager.GadgetTracerManager/AddContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetTracerManagerServer).AddContainer(ctx, req.(*ContainerDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetTracerManager_RemoveContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetTracerManagerServer).RemoveContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gadgettracermanager.GadgetTracerManager/RemoveContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetTracerManagerServer).RemoveContainer(ctx, req.(*ContainerDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetTracerManager_DumpState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetTracerManagerServer).DumpState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gadgettracermanager.GadgetTracerManager/DumpState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetTracerManagerServer).DumpState(ctx, req.(*DumpStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GadgetTracerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gadgettracermanager.GadgetTracerManager",
	HandlerType: (*GadgetTracerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTracer",
			Handler:    _GadgetTracerManager_AddTracer_Handler,
		},
		{
			MethodName: "RemoveTracer",
			Handler:    _GadgetTracerManager_RemoveTracer_Handler,
		},
		{
			MethodName: "AddContainer",
			Handler:    _GadgetTracerManager_AddContainer_Handler,
		},
		{
			MethodName: "RemoveContainer",
			Handler:    _GadgetTracerManager_RemoveContainer_Handler,
		},
		{
			MethodName: "DumpState",
			Handler:    _GadgetTracerManager_DumpState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gadgettracermanager.proto",
}

func init() { proto.RegisterFile("gadgettracermanager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0xbd, 0xc9, 0x24, 0x6a, 0xc3, 0xa4, 0x12, 0x6e, 0x28, 0x22, 0xac, 0x04, 0x18, 0xa9,
	0x6a, 0xa5, 0xf0, 0x05, 0x85, 0x48, 0x28, 0x12, 0x08, 0xe4, 0xf0, 0xc4, 0x4b, 0xb5, 0xf1, 0x4e,
	0x5d, 0xab, 0xf1, 0xda, 0xd8, 0xeb, 0x0a, 0x3e, 0x87, 0x37, 0xfe, 0x8d, 0x9f, 0x40, 0xeb, 0x5b,
	0xd2, 0xb0, 0x09, 0xe5, 0x2d, 0x73, 0x7c, 0xe6, 0x76, 0x66, 0x4f, 0xe0, 0xc4, 0xe3, 0xc2, 0x23,
	0xa5, 0x62, 0xee, 0x52, 0x1c, 0x70, 0xc9, 0x3d, 0x8a, 0xcf, 0xa3, 0x38, 0x54, 0x21, 0x8e, 0x0c,
	0x9f, 0xd8, 0x05, 0xb4, 0x3f, 0xf0, 0x25, 0xad, 0x70, 0x08, 0xcd, 0x5b, 0xfa, 0x61, 0xd5, 0x27,
	0x75, 0xbb, 0xe7, 0xe8, 0x9f, 0x78, 0x0c, 0xed, 0x3b, 0xbe, 0x4a, 0xc9, 0x6a, 0x64, 0x58, 0x1e,
	0xb0, 0x6b, 0x18, 0x5e, 0x0a, 0xf1, 0x25, 0x2b, 0xe2, 0xd0, 0xb7, 0x94, 0x12, 0x85, 0x87, 0xd0,
	0xf0, 0x45, 0x91, 0xda, 0xf0, 0x05, 0xbe, 0x85, 0x6e, 0x42, 0x2b, 0x72, 0x55, 0x18, 0x67, 0xc9,
	0xfd, 0xe9, 0xcb, 0x73, 0xd3, 0x5c, 0xef, 0x42, 0xa9, 0xb8, 0x2f, 0x29, 0x5e, 0x14, 0x6c, 0xa7,
	0xca, 0x63, 0x67, 0x70, 0xec, 0x50, 0x10, 0xde, 0x51, 0xd9, 0x2a, 0x89, 0x42, 0x99, 0x90, 0x9e,
	0x4a, 0xd0, 0x32, 0xf5, 0x8a, 0x76, 0x79, 0xa0, 0xd9, 0x97, 0x42, 0x54, 0xf5, 0xfe, 0xc1, 0xbe,
	0x80, 0xc7, 0x79, 0xed, 0x87, 0x26, 0xfc, 0xaa, 0xc3, 0xa3, 0xbf, 0x86, 0xc5, 0x53, 0xe8, 0x49,
	0x1e, 0x50, 0x12, 0x71, 0x97, 0x0a, 0xfe, 0x1a, 0x40, 0x0b, 0x0e, 0xa2, 0x50, 0xe8, 0xb8, 0x10,
	0xb0, 0x0c, 0x71, 0x0a, 0x9d, 0x95, 0xd6, 0x3c, 0xb1, 0x9a, 0x93, 0xa6, 0xdd, 0x9f, 0x8e, 0x8d,
	0xe2, 0x64, 0x67, 0x71, 0x0a, 0x26, 0xbe, 0x82, 0x23, 0xb7, 0x1c, 0xe0, 0xca, 0x97, 0x82, 0xbe,
	0x5b, 0xad, 0x49, 0xdd, 0x6e, 0x3b, 0x87, 0x15, 0x3c, 0xd7, 0x28, 0x1b, 0x43, 0x37, 0x57, 0x6c,
	0x3e, 0xdb, 0xbe, 0x0b, 0xfb, 0xd9, 0x80, 0x51, 0xb5, 0xc6, 0x8c, 0xae, 0x7d, 0xe9, 0x2b, 0x3f,
	0x94, 0xf8, 0x1c, 0x06, 0x1b, 0xc5, 0xcb, 0x8c, 0xfe, 0xba, 0xb2, 0xc0, 0x67, 0xd0, 0x77, 0xbd,
	0x38, 0x4c, 0xa3, 0xab, 0x88, 0xab, 0x9b, 0x62, 0x23, 0xc8, 0xa1, 0xcf, 0x5c, 0xdd, 0xe0, 0x13,
	0xe8, 0x15, 0x04, 0x5f, 0x58, 0xcd, 0x49, 0xdd, 0x6e, 0x39, 0xdd, 0x1c, 0x98, 0x0b, 0xad, 0x6a,
	0x20, 0x95, 0x4c, 0xb2, 0x99, 0x5b, 0x4e, 0x1e, 0xdc, 0xd7, 0xaf, 0xbd, 0x47, 0xbf, 0xce, 0x7d,
	0xfd, 0x0c, 0x5a, 0x1c, 0x98, 0xb4, 0xd8, 0x10, 0xba, 0xfb, 0x50, 0xa1, 0x19, 0xc2, 0x70, 0x96,
	0x06, 0xd1, 0x42, 0x71, 0x45, 0xc5, 0xfb, 0x66, 0xa7, 0xd0, 0xd2, 0x98, 0x5e, 0x23, 0xd1, 0x78,
	0xf9, 0x38, 0xb2, 0x60, 0xfa, 0xbb, 0x09, 0xa3, 0xf7, 0x59, 0xdd, 0x5c, 0xf8, 0x8f, 0x79, 0x5d,
	0x5c, 0x40, 0xaf, 0x72, 0x0a, 0xbe, 0x30, 0xb6, 0xde, 0x76, 0xd2, 0xf8, 0xa9, 0x91, 0x56, 0x1e,
	0x94, 0xd5, 0xf0, 0x2b, 0x0c, 0x36, 0x6d, 0x81, 0xfb, 0x13, 0xc6, 0xaf, 0x8d, 0x9f, 0x4d, 0xc6,
	0x62, 0x35, 0x24, 0x18, 0x6c, 0x9a, 0x08, 0xed, 0xfd, 0xa6, 0x5d, 0x3f, 0xa0, 0x1d, 0x6d, 0x4c,
	0x8e, 0x64, 0x35, 0xbc, 0x85, 0xa3, 0x2d, 0xf7, 0xfd, 0x47, 0xa7, 0xb3, 0x3d, 0x0b, 0x99, 0x9a,
	0x7d, 0x82, 0x5e, 0x75, 0xce, 0x1d, 0x47, 0xd8, 0x3e, 0xf7, 0xf8, 0x64, 0x27, 0x8d, 0xd5, 0x96,
	0x9d, 0xec, 0xcf, 0xf4, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xee, 0x4b, 0xf9, 0x69,
	0x05, 0x00, 0x00,
}
