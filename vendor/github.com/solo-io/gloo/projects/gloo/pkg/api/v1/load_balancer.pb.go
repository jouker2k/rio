// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/lbhash"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// LoadBalancerConfig is the settings for the load balancer used to send request to the Upstream
// endpoints.
type LoadBalancerConfig struct {
	// Configures envoy's panic threshold Percent between 0-100. Once the number of non health hosts
	// reaches this percentage, envoy disregards health information.
	// see more info [here](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/load_balancing/panic_threshold#arch-overview-load-balancing-panic-threshold).
	HealthyPanicThreshold *types.DoubleValue `protobuf:"bytes,1,opt,name=healthy_panic_threshold,json=healthyPanicThreshold,proto3" json:"healthy_panic_threshold,omitempty"`
	// This allows batch updates of endpoints health/weight/metadata that happen during a time window.
	// this help lower cpu usage when endpoint change rate is high. defaults to 1 second.
	// Set to 0 to disable and have changes applied immediately.
	UpdateMergeWindow *time.Duration `protobuf:"bytes,2,opt,name=update_merge_window,json=updateMergeWindow,proto3,stdduration" json:"update_merge_window,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*LoadBalancerConfig_RoundRobin_
	//	*LoadBalancerConfig_LeastRequest_
	//	*LoadBalancerConfig_Random_
	//	*LoadBalancerConfig_RingHash_
	//	*LoadBalancerConfig_Maglev_
	Type                 isLoadBalancerConfig_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LoadBalancerConfig) Reset()         { *m = LoadBalancerConfig{} }
func (m *LoadBalancerConfig) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig) ProtoMessage()    {}
func (*LoadBalancerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0}
}
func (m *LoadBalancerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig.Unmarshal(m, b)
}
func (m *LoadBalancerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig.Merge(m, src)
}
func (m *LoadBalancerConfig) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig.Size(m)
}
func (m *LoadBalancerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig proto.InternalMessageInfo

type isLoadBalancerConfig_Type interface {
	isLoadBalancerConfig_Type()
	Equal(interface{}) bool
}

type LoadBalancerConfig_RoundRobin_ struct {
	RoundRobin *LoadBalancerConfig_RoundRobin `protobuf:"bytes,3,opt,name=round_robin,json=roundRobin,proto3,oneof" json:"round_robin,omitempty"`
}
type LoadBalancerConfig_LeastRequest_ struct {
	LeastRequest *LoadBalancerConfig_LeastRequest `protobuf:"bytes,4,opt,name=least_request,json=leastRequest,proto3,oneof" json:"least_request,omitempty"`
}
type LoadBalancerConfig_Random_ struct {
	Random *LoadBalancerConfig_Random `protobuf:"bytes,5,opt,name=random,proto3,oneof" json:"random,omitempty"`
}
type LoadBalancerConfig_RingHash_ struct {
	RingHash *LoadBalancerConfig_RingHash `protobuf:"bytes,6,opt,name=ring_hash,json=ringHash,proto3,oneof" json:"ring_hash,omitempty"`
}
type LoadBalancerConfig_Maglev_ struct {
	Maglev *LoadBalancerConfig_Maglev `protobuf:"bytes,7,opt,name=maglev,proto3,oneof" json:"maglev,omitempty"`
}

func (*LoadBalancerConfig_RoundRobin_) isLoadBalancerConfig_Type()   {}
func (*LoadBalancerConfig_LeastRequest_) isLoadBalancerConfig_Type() {}
func (*LoadBalancerConfig_Random_) isLoadBalancerConfig_Type()       {}
func (*LoadBalancerConfig_RingHash_) isLoadBalancerConfig_Type()     {}
func (*LoadBalancerConfig_Maglev_) isLoadBalancerConfig_Type()       {}

func (m *LoadBalancerConfig) GetType() isLoadBalancerConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *LoadBalancerConfig) GetHealthyPanicThreshold() *types.DoubleValue {
	if m != nil {
		return m.HealthyPanicThreshold
	}
	return nil
}

func (m *LoadBalancerConfig) GetUpdateMergeWindow() *time.Duration {
	if m != nil {
		return m.UpdateMergeWindow
	}
	return nil
}

func (m *LoadBalancerConfig) GetRoundRobin() *LoadBalancerConfig_RoundRobin {
	if x, ok := m.GetType().(*LoadBalancerConfig_RoundRobin_); ok {
		return x.RoundRobin
	}
	return nil
}

func (m *LoadBalancerConfig) GetLeastRequest() *LoadBalancerConfig_LeastRequest {
	if x, ok := m.GetType().(*LoadBalancerConfig_LeastRequest_); ok {
		return x.LeastRequest
	}
	return nil
}

func (m *LoadBalancerConfig) GetRandom() *LoadBalancerConfig_Random {
	if x, ok := m.GetType().(*LoadBalancerConfig_Random_); ok {
		return x.Random
	}
	return nil
}

func (m *LoadBalancerConfig) GetRingHash() *LoadBalancerConfig_RingHash {
	if x, ok := m.GetType().(*LoadBalancerConfig_RingHash_); ok {
		return x.RingHash
	}
	return nil
}

func (m *LoadBalancerConfig) GetMaglev() *LoadBalancerConfig_Maglev {
	if x, ok := m.GetType().(*LoadBalancerConfig_Maglev_); ok {
		return x.Maglev
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoadBalancerConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoadBalancerConfig_RoundRobin_)(nil),
		(*LoadBalancerConfig_LeastRequest_)(nil),
		(*LoadBalancerConfig_Random_)(nil),
		(*LoadBalancerConfig_RingHash_)(nil),
		(*LoadBalancerConfig_Maglev_)(nil),
	}
}

type LoadBalancerConfig_RoundRobin struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_RoundRobin) Reset()         { *m = LoadBalancerConfig_RoundRobin{} }
func (m *LoadBalancerConfig_RoundRobin) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RoundRobin) ProtoMessage()    {}
func (*LoadBalancerConfig_RoundRobin) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 0}
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RoundRobin.Merge(m, src)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RoundRobin.Size(m)
}
func (m *LoadBalancerConfig_RoundRobin) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RoundRobin.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RoundRobin proto.InternalMessageInfo

type LoadBalancerConfig_LeastRequest struct {
	// How many choices to take into account. defaults to 2.
	ChoiceCount          uint32   `protobuf:"varint,1,opt,name=choice_count,json=choiceCount,proto3" json:"choice_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_LeastRequest) Reset()         { *m = LoadBalancerConfig_LeastRequest{} }
func (m *LoadBalancerConfig_LeastRequest) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_LeastRequest) ProtoMessage()    {}
func (*LoadBalancerConfig_LeastRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 1}
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_LeastRequest.Merge(m, src)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_LeastRequest.Size(m)
}
func (m *LoadBalancerConfig_LeastRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_LeastRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_LeastRequest proto.InternalMessageInfo

func (m *LoadBalancerConfig_LeastRequest) GetChoiceCount() uint32 {
	if m != nil {
		return m.ChoiceCount
	}
	return 0
}

type LoadBalancerConfig_Random struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_Random) Reset()         { *m = LoadBalancerConfig_Random{} }
func (m *LoadBalancerConfig_Random) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_Random) ProtoMessage()    {}
func (*LoadBalancerConfig_Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 2}
}
func (m *LoadBalancerConfig_Random) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_Random.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_Random.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_Random.Merge(m, src)
}
func (m *LoadBalancerConfig_Random) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_Random.Size(m)
}
func (m *LoadBalancerConfig_Random) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_Random.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_Random proto.InternalMessageInfo

// Customizes the parameters used in the hashing algorithm to refine performance or resource usage.
type LoadBalancerConfig_RingHashConfig struct {
	// Minimum hash ring size. The larger the ring is (that is, the more hashes there are for each provided host)
	// the better the request distribution will reflect the desired weights. Defaults to 1024 entries, and limited
	// to 8M entries.
	MinimumRingSize uint64 `protobuf:"varint,1,opt,name=minimum_ring_size,json=minimumRingSize,proto3" json:"minimum_ring_size,omitempty"`
	// Maximum hash ring size. Defaults to 8M entries, and limited to 8M entries, but can be lowered to further
	// constrain resource use.
	MaximumRingSize      uint64   `protobuf:"varint,2,opt,name=maximum_ring_size,json=maximumRingSize,proto3" json:"maximum_ring_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_RingHashConfig) Reset()         { *m = LoadBalancerConfig_RingHashConfig{} }
func (m *LoadBalancerConfig_RingHashConfig) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RingHashConfig) ProtoMessage()    {}
func (*LoadBalancerConfig_RingHashConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 3}
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Merge(m, src)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RingHashConfig.Size(m)
}
func (m *LoadBalancerConfig_RingHashConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RingHashConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RingHashConfig proto.InternalMessageInfo

func (m *LoadBalancerConfig_RingHashConfig) GetMinimumRingSize() uint64 {
	if m != nil {
		return m.MinimumRingSize
	}
	return 0
}

func (m *LoadBalancerConfig_RingHashConfig) GetMaximumRingSize() uint64 {
	if m != nil {
		return m.MaximumRingSize
	}
	return 0
}

type LoadBalancerConfig_RingHash struct {
	// Optional, customizes the parameters used in the hashing algorithm
	RingHashConfig       *LoadBalancerConfig_RingHashConfig `protobuf:"bytes,1,opt,name=ring_hash_config,json=ringHashConfig,proto3" json:"ring_hash_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *LoadBalancerConfig_RingHash) Reset()         { *m = LoadBalancerConfig_RingHash{} }
func (m *LoadBalancerConfig_RingHash) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_RingHash) ProtoMessage()    {}
func (*LoadBalancerConfig_RingHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 4}
}
func (m *LoadBalancerConfig_RingHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_RingHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_RingHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_RingHash.Merge(m, src)
}
func (m *LoadBalancerConfig_RingHash) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_RingHash.Size(m)
}
func (m *LoadBalancerConfig_RingHash) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_RingHash.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_RingHash proto.InternalMessageInfo

func (m *LoadBalancerConfig_RingHash) GetRingHashConfig() *LoadBalancerConfig_RingHashConfig {
	if m != nil {
		return m.RingHashConfig
	}
	return nil
}

type LoadBalancerConfig_Maglev struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadBalancerConfig_Maglev) Reset()         { *m = LoadBalancerConfig_Maglev{} }
func (m *LoadBalancerConfig_Maglev) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerConfig_Maglev) ProtoMessage()    {}
func (*LoadBalancerConfig_Maglev) Descriptor() ([]byte, []int) {
	return fileDescriptor_aaa1c019b03e4b0f, []int{0, 5}
}
func (m *LoadBalancerConfig_Maglev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Unmarshal(m, b)
}
func (m *LoadBalancerConfig_Maglev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Marshal(b, m, deterministic)
}
func (m *LoadBalancerConfig_Maglev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadBalancerConfig_Maglev.Merge(m, src)
}
func (m *LoadBalancerConfig_Maglev) XXX_Size() int {
	return xxx_messageInfo_LoadBalancerConfig_Maglev.Size(m)
}
func (m *LoadBalancerConfig_Maglev) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadBalancerConfig_Maglev.DiscardUnknown(m)
}

var xxx_messageInfo_LoadBalancerConfig_Maglev proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoadBalancerConfig)(nil), "gloo.solo.io.LoadBalancerConfig")
	proto.RegisterType((*LoadBalancerConfig_RoundRobin)(nil), "gloo.solo.io.LoadBalancerConfig.RoundRobin")
	proto.RegisterType((*LoadBalancerConfig_LeastRequest)(nil), "gloo.solo.io.LoadBalancerConfig.LeastRequest")
	proto.RegisterType((*LoadBalancerConfig_Random)(nil), "gloo.solo.io.LoadBalancerConfig.Random")
	proto.RegisterType((*LoadBalancerConfig_RingHashConfig)(nil), "gloo.solo.io.LoadBalancerConfig.RingHashConfig")
	proto.RegisterType((*LoadBalancerConfig_RingHash)(nil), "gloo.solo.io.LoadBalancerConfig.RingHash")
	proto.RegisterType((*LoadBalancerConfig_Maglev)(nil), "gloo.solo.io.LoadBalancerConfig.Maglev")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/load_balancer.proto", fileDescriptor_aaa1c019b03e4b0f)
}

var fileDescriptor_aaa1c019b03e4b0f = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x93, 0x12, 0x4c, 0xd8, 0xa4, 0x85, 0x1a, 0x10, 0xc6, 0x42, 0xe5, 0xe3, 0xc2, 0x97,
	0x6a, 0x53, 0x90, 0x38, 0x43, 0xca, 0xc1, 0x87, 0x16, 0xd0, 0x12, 0x81, 0xe0, 0x62, 0xad, 0xed,
	0xed, 0x7a, 0x61, 0xbd, 0x63, 0xd6, 0x76, 0x43, 0xfb, 0x24, 0x3c, 0x02, 0x6f, 0x55, 0x89, 0x27,
	0x41, 0xfb, 0x91, 0x12, 0xa8, 0x50, 0xd2, 0x53, 0x76, 0x76, 0xe6, 0xff, 0xf3, 0xfc, 0x27, 0x63,
	0xa3, 0x97, 0x8c, 0xb7, 0x65, 0x97, 0x45, 0x39, 0x54, 0x71, 0x03, 0x02, 0xb6, 0x39, 0xc4, 0x4c,
	0x00, 0xc4, 0xb5, 0x82, 0x2f, 0x34, 0x6f, 0x1b, 0x1b, 0x91, 0x9a, 0xc7, 0x87, 0x3b, 0xb1, 0x00,
	0x52, 0xa4, 0x19, 0x11, 0x44, 0xe6, 0x54, 0x45, 0xb5, 0x82, 0x16, 0xfc, 0xb1, 0x2e, 0x88, 0xb4,
	0x36, 0xe2, 0x10, 0x26, 0xe7, 0xe2, 0x41, 0xdd, 0x72, 0x90, 0x4d, 0x2c, 0xb2, 0x92, 0x34, 0xa5,
	0xfb, 0xb1, 0xdc, 0xf0, 0x3a, 0x03, 0x06, 0xe6, 0x18, 0xeb, 0x93, 0xbb, 0xdd, 0x62, 0x00, 0x4c,
	0xd0, 0xd8, 0x44, 0x59, 0x77, 0x10, 0x17, 0x9d, 0x22, 0x1a, 0xf2, 0xbf, 0xfc, 0x4c, 0x91, 0xba,
	0xa6, 0xaa, 0xb1, 0xf9, 0xfb, 0x27, 0x1e, 0xf2, 0xf7, 0x80, 0x14, 0x13, 0x67, 0x62, 0x17, 0xe4,
	0x01, 0x67, 0xfe, 0x14, 0xdd, 0x2c, 0x29, 0x11, 0x6d, 0x79, 0x94, 0xd6, 0x44, 0xf2, 0x3c, 0x6d,
	0x4b, 0x45, 0x9b, 0x12, 0x44, 0x11, 0xf4, 0xef, 0xf6, 0x1f, 0x8e, 0x9e, 0xdd, 0x8e, 0x2c, 0x38,
	0x9a, 0x83, 0xa3, 0xd7, 0xd0, 0x65, 0x82, 0x7e, 0x20, 0xa2, 0xa3, 0xf8, 0x86, 0x13, 0xbf, 0xd3,
	0xda, 0xe9, 0x5c, 0xea, 0xbf, 0x45, 0xd7, 0xba, 0xba, 0x20, 0x2d, 0x4d, 0x2b, 0xaa, 0x18, 0x4d,
	0x67, 0x5c, 0x16, 0x30, 0x0b, 0xd6, 0x0c, 0xf1, 0xd6, 0x59, 0xa2, 0xb3, 0x32, 0x19, 0xfc, 0x38,
	0xb9, 0xd3, 0xc7, 0x9b, 0x56, 0xbb, 0xaf, 0xa5, 0x1f, 0x8d, 0xd2, 0x7f, 0x83, 0x46, 0x0a, 0x3a,
	0x59, 0xa4, 0x0a, 0x32, 0x2e, 0x83, 0x0b, 0x06, 0xf4, 0x24, 0x5a, 0xfc, 0x07, 0xa2, 0xb3, 0xee,
	0x22, 0xac, 0x35, 0x58, 0x4b, 0x92, 0x1e, 0x46, 0xea, 0x34, 0xf2, 0xa7, 0x68, 0x5d, 0x50, 0xd2,
	0xb4, 0xa9, 0xa2, 0xdf, 0x3a, 0xda, 0xb4, 0xc1, 0xc0, 0x10, 0xb7, 0x97, 0x12, 0xf7, 0xb4, 0x0a,
	0x5b, 0x51, 0xd2, 0xc3, 0x63, 0xb1, 0x10, 0xfb, 0xaf, 0x90, 0xa7, 0x88, 0x2c, 0xa0, 0x0a, 0x2e,
	0x1a, 0xdc, 0x83, 0xe5, 0x0d, 0x9a, 0xf2, 0xa4, 0x87, 0x9d, 0xd0, 0x4f, 0xd0, 0x65, 0xc5, 0x25,
	0x4b, 0xf5, 0x3e, 0x04, 0x9e, 0xa1, 0x3c, 0x5a, 0x4e, 0xe1, 0x92, 0x25, 0xa4, 0x29, 0x93, 0x1e,
	0x1e, 0x2a, 0x77, 0xd6, 0xcd, 0x54, 0x84, 0x09, 0x7a, 0x18, 0x5c, 0x5a, 0xb1, 0x99, 0x7d, 0x53,
	0xae, 0x9b, 0xb1, 0xc2, 0x70, 0x8c, 0xd0, 0x9f, 0x09, 0x86, 0x3b, 0x68, 0xbc, 0xe8, 0xde, 0xbf,
	0x87, 0xc6, 0x79, 0x09, 0x3c, 0xa7, 0x69, 0x0e, 0x9d, 0x6c, 0xcd, 0xbe, 0xac, 0xe3, 0x91, 0xbd,
	0xdb, 0xd5, 0x57, 0xe1, 0x10, 0x79, 0xd6, 0x61, 0x58, 0xa2, 0x8d, 0x79, 0x97, 0x6e, 0xf3, 0x1e,
	0xa3, 0xcd, 0x8a, 0x4b, 0x5e, 0x75, 0x55, 0x6a, 0x1c, 0x37, 0xfc, 0x98, 0x1a, 0xc6, 0x00, 0x5f,
	0x71, 0x09, 0xad, 0x78, 0xcf, 0x8f, 0xa9, 0xa9, 0x25, 0xdf, 0xff, 0xa9, 0x5d, 0x73, 0xb5, 0x36,
	0x31, 0xaf, 0x0d, 0x29, 0x1a, 0xce, 0x9f, 0xe4, 0x7f, 0x42, 0x57, 0x4f, 0xa7, 0x99, 0xe6, 0xe6,
	0xb9, 0x6e, 0xad, 0xe3, 0x95, 0x87, 0x6a, 0x43, 0xbc, 0xa1, 0xfe, 0x8a, 0xb5, 0x35, 0x3b, 0xaf,
	0x89, 0x87, 0x06, 0xed, 0x51, 0x4d, 0x27, 0x2f, 0x7e, 0xfe, 0xda, 0xea, 0x7f, 0x7e, 0xba, 0xda,
	0x77, 0xa0, 0xfe, 0xca, 0xdc, 0xb7, 0x20, 0xf3, 0xcc, 0x7b, 0xf0, 0xfc, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x47, 0x80, 0x86, 0x11, 0x92, 0x04, 0x00, 0x00,
}

func (this *LoadBalancerConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig)
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
	if !this.HealthyPanicThreshold.Equal(that1.HealthyPanicThreshold) {
		return false
	}
	if this.UpdateMergeWindow != nil && that1.UpdateMergeWindow != nil {
		if *this.UpdateMergeWindow != *that1.UpdateMergeWindow {
			return false
		}
	} else if this.UpdateMergeWindow != nil {
		return false
	} else if that1.UpdateMergeWindow != nil {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RoundRobin_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RoundRobin_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin_)
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
	if !this.RoundRobin.Equal(that1.RoundRobin) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_LeastRequest_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_LeastRequest_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest_)
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
	if !this.LeastRequest.Equal(that1.LeastRequest) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Random_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Random_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random_)
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
	if !this.Random.Equal(that1.Random) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHash_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHash_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash_)
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
	if !this.RingHash.Equal(that1.RingHash) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Maglev_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Maglev_)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev_)
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
	if !this.Maglev.Equal(that1.Maglev) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RoundRobin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RoundRobin)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RoundRobin)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_LeastRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_LeastRequest)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_LeastRequest)
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
	if this.ChoiceCount != that1.ChoiceCount {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Random) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Random)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Random)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHashConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHashConfig)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHashConfig)
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
	if this.MinimumRingSize != that1.MinimumRingSize {
		return false
	}
	if this.MaximumRingSize != that1.MaximumRingSize {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_RingHash) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_RingHash)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_RingHash)
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
	if !this.RingHashConfig.Equal(that1.RingHashConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LoadBalancerConfig_Maglev) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LoadBalancerConfig_Maglev)
	if !ok {
		that2, ok := that.(LoadBalancerConfig_Maglev)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
