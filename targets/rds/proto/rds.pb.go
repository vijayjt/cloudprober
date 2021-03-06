// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/targets/rds/proto/rds.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/targets/rds/proto/rds.proto

It has these top-level messages:
	ListResourcesRequest
	Filter
	IPConfig
	Resource
	ListResourcesResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IPConfig_IPType int32

const (
	// Default IP of the resource.
	//  - Private IP for instance resource
	//  - Forwarding rule IP for forwarding rule.
	IPConfig_DEFAULT IPConfig_IPType = 0
	// Instance's external IP.
	IPConfig_PUBLIC IPConfig_IPType = 1
	// First IP address from the first Alias IP range. For example, for
	// alias IP range "192.168.12.0/24", 192.168.12.0 will be returned.
	// Supported only on GCE.
	IPConfig_ALIAS IPConfig_IPType = 2
)

var IPConfig_IPType_name = map[int32]string{
	0: "DEFAULT",
	1: "PUBLIC",
	2: "ALIAS",
}
var IPConfig_IPType_value = map[string]int32{
	"DEFAULT": 0,
	"PUBLIC":  1,
	"ALIAS":   2,
}

func (x IPConfig_IPType) Enum() *IPConfig_IPType {
	p := new(IPConfig_IPType)
	*p = x
	return p
}
func (x IPConfig_IPType) String() string {
	return proto1.EnumName(IPConfig_IPType_name, int32(x))
}
func (x *IPConfig_IPType) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(IPConfig_IPType_value, data, "IPConfig_IPType")
	if err != nil {
		return err
	}
	*x = IPConfig_IPType(value)
	return nil
}
func (IPConfig_IPType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type ListResourcesRequest struct {
	// Provider is the resource list provider, for example: "gcp", "aws", etc.
	Provider *string `protobuf:"bytes,1,req,name=provider" json:"provider,omitempty"`
	// Provider specific resource path. For example: for GCP, it could be
	// "gce_instances/<project>", "regional_forwarding_rules/<project>", etc.
	ResourcePath *string `protobuf:"bytes,2,opt,name=resource_path,json=resourcePath" json:"resource_path,omitempty"`
	// Filter for the resources list.
	Filter *Filter `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
	// Optional. If resource has an IP (and a NIC) address, following
	// fields determine which IP address will be included in the results.
	IpConfig         *IPConfig `protobuf:"bytes,4,opt,name=ip_config,json=ipConfig" json:"ip_config,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ListResourcesRequest) Reset()                    { *m = ListResourcesRequest{} }
func (m *ListResourcesRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListResourcesRequest) ProtoMessage()               {}
func (*ListResourcesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListResourcesRequest) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *ListResourcesRequest) GetResourcePath() string {
	if m != nil && m.ResourcePath != nil {
		return *m.ResourcePath
	}
	return ""
}

func (m *ListResourcesRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListResourcesRequest) GetIpConfig() *IPConfig {
	if m != nil {
		return m.IpConfig
	}
	return nil
}

type Filter struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Regex            *string `protobuf:"bytes,2,req,name=regex" json:"regex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto1.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Filter) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Filter) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

type IPConfig struct {
	// NIC index
	NicIndex         *int32           `protobuf:"varint,1,opt,name=nic_index,json=nicIndex,def=0" json:"nic_index,omitempty"`
	IpType           *IPConfig_IPType `protobuf:"varint,3,opt,name=ip_type,json=ipType,enum=cloudprober.targets.rds.IPConfig_IPType" json:"ip_type,omitempty"`
	IpVersion        *int32           `protobuf:"varint,4,opt,name=ip_version,json=ipVersion,def=4" json:"ip_version,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *IPConfig) Reset()                    { *m = IPConfig{} }
func (m *IPConfig) String() string            { return proto1.CompactTextString(m) }
func (*IPConfig) ProtoMessage()               {}
func (*IPConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_IPConfig_NicIndex int32 = 0
const Default_IPConfig_IpVersion int32 = 4

func (m *IPConfig) GetNicIndex() int32 {
	if m != nil && m.NicIndex != nil {
		return *m.NicIndex
	}
	return Default_IPConfig_NicIndex
}

func (m *IPConfig) GetIpType() IPConfig_IPType {
	if m != nil && m.IpType != nil {
		return *m.IpType
	}
	return IPConfig_DEFAULT
}

func (m *IPConfig) GetIpVersion() int32 {
	if m != nil && m.IpVersion != nil {
		return *m.IpVersion
	}
	return Default_IPConfig_IpVersion
}

type Resource struct {
	// Resource name.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// Resource's IP address, selected based on the request's ip_config.
	Ip *string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	// Id associated with the resource.
	Id               *string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto1.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Resource) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Resource) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Resource) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type ListResourcesResponse struct {
	Resources        []*Resource `protobuf:"bytes,1,rep,name=resources" json:"resources,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ListResourcesResponse) Reset()                    { *m = ListResourcesResponse{} }
func (m *ListResourcesResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListResourcesResponse) ProtoMessage()               {}
func (*ListResourcesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListResourcesResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto1.RegisterType((*ListResourcesRequest)(nil), "cloudprober.targets.rds.ListResourcesRequest")
	proto1.RegisterType((*Filter)(nil), "cloudprober.targets.rds.Filter")
	proto1.RegisterType((*IPConfig)(nil), "cloudprober.targets.rds.IPConfig")
	proto1.RegisterType((*Resource)(nil), "cloudprober.targets.rds.Resource")
	proto1.RegisterType((*ListResourcesResponse)(nil), "cloudprober.targets.rds.ListResourcesResponse")
	proto1.RegisterEnum("cloudprober.targets.rds.IPConfig_IPType", IPConfig_IPType_name, IPConfig_IPType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceDiscovery service

type ResourceDiscoveryClient interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error)
}

type resourceDiscoveryClient struct {
	cc *grpc.ClientConn
}

func NewResourceDiscoveryClient(cc *grpc.ClientConn) ResourceDiscoveryClient {
	return &resourceDiscoveryClient{cc}
}

func (c *resourceDiscoveryClient) ListResources(ctx context.Context, in *ListResourcesRequest, opts ...grpc.CallOption) (*ListResourcesResponse, error) {
	out := new(ListResourcesResponse)
	err := grpc.Invoke(ctx, "/cloudprober.targets.rds.ResourceDiscovery/ListResources", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceDiscovery service

type ResourceDiscoveryServer interface {
	// ListResources returns the list of resources matching the URI provided in
	// the request.
	ListResources(context.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
}

func RegisterResourceDiscoveryServer(s *grpc.Server, srv ResourceDiscoveryServer) {
	s.RegisterService(&_ResourceDiscovery_serviceDesc, srv)
}

func _ResourceDiscovery_ListResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.targets.rds.ResourceDiscovery/ListResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceDiscoveryServer).ListResources(ctx, req.(*ListResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceDiscovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprober.targets.rds.ResourceDiscovery",
	HandlerType: (*ResourceDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListResources",
			Handler:    _ResourceDiscovery_ListResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/google/cloudprober/targets/rds/proto/rds.proto",
}

func init() {
	proto1.RegisterFile("github.com/google/cloudprober/targets/rds/proto/rds.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xec, 0xba, 0x8d, 0x6b, 0xbf, 0xd2, 0x2a, 0xac, 0x8a, 0xb0, 0x7a, 0x00, 0x63, 0x2e, 0x3e,
	0x80, 0x5d, 0x45, 0x48, 0x88, 0x1e, 0x8a, 0x42, 0x4b, 0x25, 0x4b, 0x39, 0x44, 0x4b, 0x8b, 0xb8,
	0x59, 0xa9, 0xfd, 0xea, 0xac, 0x48, 0xbd, 0xcb, 0xee, 0x26, 0x6a, 0x7e, 0x00, 0x3f, 0x8b, 0xff,
	0xc0, 0x4f, 0x42, 0x5e, 0xdb, 0x7c, 0x89, 0x00, 0x27, 0xcf, 0x1b, 0xcd, 0x3c, 0xbf, 0xd1, 0x2c,
	0xbc, 0xaa, 0xb8, 0x99, 0x2f, 0xaf, 0x93, 0x42, 0xdc, 0xa6, 0x95, 0x10, 0xd5, 0x02, 0xd3, 0x62,
	0x21, 0x96, 0xa5, 0x54, 0xe2, 0x1a, 0x55, 0x6a, 0x66, 0xaa, 0x42, 0xa3, 0x53, 0x55, 0xea, 0x54,
	0x2a, 0x61, 0x44, 0x83, 0x12, 0x8b, 0xe8, 0xc3, 0x9f, 0x84, 0x49, 0x27, 0x4c, 0x54, 0xa9, 0xa3,
	0xaf, 0x04, 0x0e, 0x27, 0x5c, 0x1b, 0x86, 0x5a, 0x2c, 0x55, 0x81, 0x9a, 0xe1, 0xa7, 0x25, 0x6a,
	0x43, 0x8f, 0xc0, 0x93, 0x4a, 0xac, 0x78, 0x89, 0x2a, 0x20, 0xa1, 0x13, 0xfb, 0xec, 0xfb, 0x4c,
	0x9f, 0xc2, 0xbe, 0xea, 0xf4, 0xb9, 0x9c, 0x99, 0x79, 0xe0, 0x84, 0x24, 0xf6, 0xd9, 0xbd, 0x9e,
	0x9c, 0xce, 0xcc, 0x9c, 0xbe, 0x04, 0xf7, 0x86, 0x2f, 0x0c, 0xaa, 0x60, 0x3b, 0x24, 0xf1, 0xde,
	0xe8, 0x71, 0xb2, 0xe1, 0x86, 0xe4, 0xc2, 0xca, 0x58, 0x27, 0xa7, 0xa7, 0xe0, 0x73, 0x99, 0x17,
	0xa2, 0xbe, 0xe1, 0x55, 0xb0, 0x63, 0xbd, 0x4f, 0x36, 0x7a, 0xb3, 0xe9, 0x99, 0x15, 0x32, 0x8f,
	0xcb, 0x16, 0x45, 0xc7, 0xe0, 0xb6, 0x1b, 0xe9, 0x10, 0xb6, 0x3f, 0xe2, 0xba, 0x3b, 0xbf, 0x81,
	0xf4, 0x10, 0x06, 0x0a, 0x2b, 0xbc, 0x0b, 0x1c, 0xcb, 0xb5, 0x43, 0xf4, 0x85, 0x80, 0xd7, 0x2f,
	0xa2, 0x8f, 0xc0, 0xaf, 0x79, 0x91, 0xf3, 0xba, 0xc4, 0xbb, 0x80, 0x84, 0x24, 0x1e, 0x9c, 0x90,
	0x63, 0xe6, 0xd5, 0xbc, 0xc8, 0x1a, 0x8a, 0x8e, 0x61, 0x97, 0xcb, 0xdc, 0xac, 0x25, 0xda, 0x60,
	0x07, 0xa3, 0xf8, 0x9f, 0xc7, 0x25, 0xd9, 0xf4, 0x72, 0x2d, 0x91, 0xb9, 0x5c, 0x36, 0x5f, 0x1a,
	0x02, 0x70, 0x99, 0xaf, 0x50, 0x69, 0x2e, 0x6a, 0x1b, 0x71, 0x70, 0x42, 0x5e, 0x30, 0x9f, 0xcb,
	0xf7, 0x2d, 0x17, 0x3d, 0x03, 0xb7, 0xf5, 0xd0, 0x3d, 0xd8, 0x3d, 0x7f, 0x7b, 0x31, 0xbe, 0x9a,
	0x5c, 0x0e, 0xb7, 0x28, 0x80, 0x3b, 0xbd, 0x7a, 0x33, 0xc9, 0xce, 0x86, 0x84, 0xfa, 0x30, 0x18,
	0x4f, 0xb2, 0xf1, 0xbb, 0xa1, 0x13, 0x9d, 0x82, 0xd7, 0xf7, 0x47, 0x29, 0xec, 0xd4, 0xb3, 0x5b,
	0xec, 0x42, 0x5b, 0x4c, 0x0f, 0xc0, 0xe1, 0xb2, 0x2b, 0xc9, 0xe1, 0xd2, 0xce, 0xa5, 0xbd, 0xbe,
	0x99, 0xcb, 0xe8, 0x03, 0x3c, 0xf8, 0xed, 0x0d, 0x68, 0x29, 0x6a, 0x8d, 0xf4, 0x35, 0xf8, 0x7d,
	0xa7, 0x3a, 0x20, 0xe1, 0xf6, 0x5f, 0xab, 0xe8, 0xed, 0xec, 0x87, 0x67, 0xf4, 0x99, 0xc0, 0xfd,
	0x9e, 0x3f, 0xe7, 0xba, 0x10, 0x2b, 0x54, 0x6b, 0x2a, 0x61, 0xff, 0x97, 0xff, 0xd1, 0xe7, 0x1b,
	0x97, 0xfe, 0xe9, 0x6d, 0x1e, 0x25, 0xff, 0x2b, 0x6f, 0x63, 0x44, 0x5b, 0xdf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x97, 0x86, 0x0c, 0x6b, 0x3b, 0x03, 0x00, 0x00,
}
