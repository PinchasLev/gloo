// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//A gateway describes the routes to upstreams that are reachable via a specific port on the Gateway Proxy itself.
//
//Deprecated: see gateway.solo.io.v2.Gateway
type Gateway struct {
	// if set to false, only use virtual services with no ssl configured.
	// if set to true, only use virtual services with ssl configured.
	Ssl bool `protobuf:"varint,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// names of the the virtual services, which contain the actual routes for the gateway
	// if the list is empty, all virtual services will apply to this gateway (with accordance to tls flag above).
	VirtualServices []core.ResourceRef `protobuf:"bytes,2,rep,name=virtual_services,json=virtualServices,proto3" json:"virtual_services"`
	// the bind address the gateway should serve traffic on
	BindAddress string `protobuf:"bytes,3,opt,name=bind_address,json=bindAddress,proto3" json:"bind_address,omitempty"`
	// bind ports must not conflict across gateways in a namespace
	BindPort uint32 `protobuf:"varint,4,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	// top level plugin configuration for all routes on the gateway
	Plugins *v1.HttpListenerPlugins `protobuf:"bytes,5,opt,name=plugins,proto3" json:"plugins,omitempty"`
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,6,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Enable ProxyProtocol support for this listener
	UseProxyProto        *types.BoolValue `protobuf:"bytes,8,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Gateway) Reset()         { *m = Gateway{} }
func (m *Gateway) String() string { return proto.CompactTextString(m) }
func (*Gateway) ProtoMessage()    {}
func (*Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_30f7529f6633771c, []int{0}
}
func (m *Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gateway.Unmarshal(m, b)
}
func (m *Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gateway.Marshal(b, m, deterministic)
}
func (m *Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gateway.Merge(m, src)
}
func (m *Gateway) XXX_Size() int {
	return xxx_messageInfo_Gateway.Size(m)
}
func (m *Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Gateway proto.InternalMessageInfo

func (m *Gateway) GetSsl() bool {
	if m != nil {
		return m.Ssl
	}
	return false
}

func (m *Gateway) GetVirtualServices() []core.ResourceRef {
	if m != nil {
		return m.VirtualServices
	}
	return nil
}

func (m *Gateway) GetBindAddress() string {
	if m != nil {
		return m.BindAddress
	}
	return ""
}

func (m *Gateway) GetBindPort() uint32 {
	if m != nil {
		return m.BindPort
	}
	return 0
}

func (m *Gateway) GetPlugins() *v1.HttpListenerPlugins {
	if m != nil {
		return m.Plugins
	}
	return nil
}

func (m *Gateway) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Gateway) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Gateway) GetUseProxyProto() *types.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func init() {
	proto.RegisterType((*Gateway)(nil), "gateway.solo.io.Gateway")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gateway/api/v1/gateway.proto", fileDescriptor_30f7529f6633771c)
}

var fileDescriptor_30f7529f6633771c = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0x3f, 0x37, 0xf9, 0x92, 0x74, 0x43, 0x95, 0xb2, 0xaa, 0x2a, 0x37, 0x48, 0x6d, 0x9a,
	0x53, 0x0e, 0xe0, 0x55, 0xdb, 0x4b, 0x09, 0xe2, 0x80, 0x2f, 0x45, 0x08, 0xa4, 0xc8, 0x95, 0x38,
	0x70, 0x89, 0x36, 0xf6, 0x78, 0x59, 0xea, 0x66, 0xac, 0xdd, 0x75, 0x02, 0xd7, 0x3e, 0x0d, 0xef,
	0xc1, 0x85, 0xa7, 0xe8, 0x81, 0x37, 0x28, 0x4f, 0x80, 0x76, 0xbd, 0xae, 0x54, 0x09, 0x89, 0xf4,
	0xe4, 0x9d, 0x99, 0xfd, 0xcd, 0xfc, 0x3d, 0x33, 0x4b, 0x5e, 0x0b, 0x69, 0x3e, 0x57, 0x8b, 0x28,
	0xc5, 0x6b, 0xa6, 0xb1, 0xc0, 0x17, 0x12, 0x99, 0x28, 0x10, 0x59, 0xa9, 0xf0, 0x0b, 0xa4, 0x46,
	0x33, 0xc1, 0x0d, 0xac, 0xf9, 0x37, 0xc6, 0x4b, 0xc9, 0x56, 0x27, 0x8d, 0x19, 0x95, 0x0a, 0x0d,
	0xd2, 0x41, 0x63, 0x5a, 0x36, 0x92, 0x38, 0x3c, 0x14, 0x88, 0xa2, 0x00, 0xe6, 0xc2, 0x8b, 0x2a,
	0x67, 0x6b, 0xc5, 0xcb, 0x12, 0x94, 0xae, 0x81, 0xe1, 0x9e, 0x40, 0x81, 0xee, 0xc8, 0xec, 0xc9,
	0x7b, 0x4f, 0xfe, 0xa2, 0xc2, 0x7d, 0xaf, 0xa4, 0x69, 0x0a, 0x5f, 0x83, 0xe1, 0x19, 0x37, 0xdc,
	0x23, 0x6c, 0x03, 0x44, 0x1b, 0x6e, 0xaa, 0xa6, 0xf2, 0xf3, 0x0d, 0x00, 0x05, 0xf9, 0x23, 0x14,
	0x35, 0xb6, 0x47, 0xce, 0xff, 0xdd, 0x4a, 0x6b, 0x79, 0xb8, 0x54, 0xf8, 0xd5, 0x77, 0x71, 0x38,
	0x7d, 0x1c, 0x59, 0x54, 0x42, 0x2e, 0xfd, 0x6f, 0x8d, 0x7f, 0xb4, 0x48, 0xf7, 0xa2, 0x1e, 0x02,
	0xdd, 0x25, 0x2d, 0xad, 0x8b, 0x30, 0x18, 0x05, 0x93, 0x5e, 0x62, 0x8f, 0xf4, 0x1d, 0xd9, 0x5d,
	0x49, 0x65, 0x2a, 0x5e, 0xcc, 0x35, 0xa8, 0x95, 0x4c, 0x41, 0x87, 0x5b, 0xa3, 0xd6, 0xa4, 0x7f,
	0x7a, 0x10, 0xa5, 0xa8, 0xa0, 0x99, 0x5b, 0x94, 0x80, 0xc6, 0x4a, 0xa5, 0x90, 0x40, 0x1e, 0xb7,
	0x7f, 0xde, 0x1e, 0xfd, 0x97, 0x0c, 0x3c, 0x78, 0xe9, 0x39, 0x7a, 0x4c, 0x9e, 0x2c, 0xe4, 0x32,
	0x9b, 0xf3, 0x2c, 0x53, 0xa0, 0x75, 0xd8, 0x1a, 0x05, 0x93, 0xed, 0xa4, 0x6f, 0x7d, 0x6f, 0x6a,
	0x17, 0x7d, 0x46, 0xb6, 0xdd, 0x95, 0x12, 0x95, 0x09, 0xdb, 0xa3, 0x60, 0xb2, 0x93, 0xf4, 0xac,
	0x63, 0x86, 0xca, 0xd0, 0x57, 0xa4, 0xeb, 0xa5, 0x87, 0xff, 0x8f, 0x82, 0x49, 0xff, 0xf4, 0x38,
	0xb2, 0xbf, 0x75, 0x2f, 0xe1, 0xad, 0x31, 0xe5, 0x7b, 0xa9, 0x0d, 0x2c, 0x41, 0xcd, 0xea, 0x8b,
	0x49, 0x43, 0xd0, 0x0b, 0xd2, 0xa9, 0xa7, 0x19, 0x76, 0x1c, 0xbb, 0xf7, 0x50, 0xfe, 0xa5, 0x8b,
	0xc5, 0x07, 0x56, 0xf9, 0xef, 0xdb, 0xa3, 0xa7, 0x06, 0xb4, 0xc9, 0x64, 0x9e, 0x4f, 0xc7, 0x52,
	0x2c, 0x51, 0xc1, 0x38, 0xf1, 0x38, 0x3d, 0x27, 0xbd, 0x66, 0x93, 0xc2, 0xae, 0x4b, 0xb5, 0xff,
	0x30, 0xd5, 0x07, 0x1f, 0xf5, 0x6d, 0xb8, 0xbf, 0x4d, 0x63, 0x32, 0xa8, 0x34, 0xcc, 0xdd, 0xe0,
	0xe6, 0xae, 0xf9, 0x61, 0xcf, 0x25, 0x18, 0x46, 0xf5, 0xd2, 0x47, 0xcd, 0xd2, 0x47, 0x31, 0x62,
	0xf1, 0x91, 0x17, 0x15, 0x24, 0x3b, 0x95, 0x86, 0x99, 0x25, 0x66, 0x36, 0x36, 0xdd, 0xbf, 0xb9,
	0x6b, 0xb7, 0xc9, 0x96, 0x58, 0xdf, 0xdc, 0xb5, 0x09, 0xed, 0xf9, 0xd7, 0xa3, 0xe3, 0x97, 0xdf,
	0x7f, 0x1d, 0x06, 0x9f, 0xce, 0x36, 0x7e, 0x8c, 0xe5, 0x95, 0xf0, 0xeb, 0xb0, 0xe8, 0xb8, 0xaa,
	0x67, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x7a, 0xd7, 0x8a, 0xca, 0x03, 0x00, 0x00,
}

func (this *Gateway) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Gateway)
	if !ok {
		that2, ok := that.(Gateway)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Ssl != that1.Ssl {
		return false
	}
	if len(this.VirtualServices) != len(that1.VirtualServices) {
		return false
	}
	for i := range this.VirtualServices {
		if !this.VirtualServices[i].Equal(&that1.VirtualServices[i]) {
			return false
		}
	}
	if this.BindAddress != that1.BindAddress {
		return false
	}
	if this.BindPort != that1.BindPort {
		return false
	}
	if !this.Plugins.Equal(that1.Plugins) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.UseProxyProto.Equal(that1.UseProxyProto) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
