// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/microservice/v1alpha1/smart_limiter.proto

package v1alpha1

import (
	fmt "fmt"
	types "github.com/gogo/protobuf/types"
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

type SmartLimiterSpec struct {
	Domain               string                            `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Sets                 map[string]*SmartLimitDescriptors `protobuf:"bytes,2,rep,name=sets,proto3" json:"sets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SmartLimiterSpec) Reset()         { *m = SmartLimiterSpec{} }
func (m *SmartLimiterSpec) String() string { return proto.CompactTextString(m) }
func (*SmartLimiterSpec) ProtoMessage()    {}
func (*SmartLimiterSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{0}
}

func (m *SmartLimiterSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimiterSpec.Unmarshal(m, b)
}
func (m *SmartLimiterSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimiterSpec.Marshal(b, m, deterministic)
}
func (m *SmartLimiterSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimiterSpec.Merge(m, src)
}
func (m *SmartLimiterSpec) XXX_Size() int {
	return xxx_messageInfo_SmartLimiterSpec.Size(m)
}
func (m *SmartLimiterSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimiterSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimiterSpec proto.InternalMessageInfo

func (m *SmartLimiterSpec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *SmartLimiterSpec) GetSets() map[string]*SmartLimitDescriptors {
	if m != nil {
		return m.Sets
	}
	return nil
}

type SmartLimiterStatus struct {
	RatelimitStatus      map[string]*SmartLimitDescriptors `protobuf:"bytes,1,rep,name=ratelimitStatus,proto3" json:"ratelimitStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	EndPointStatus       map[string]string                 `protobuf:"bytes,2,rep,name=endPointStatus,proto3" json:"endPointStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ServiceStatus        *SmartLimiterStatus_ServiceStatus `protobuf:"bytes,3,opt,name=serviceStatus,proto3" json:"serviceStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SmartLimiterStatus) Reset()         { *m = SmartLimiterStatus{} }
func (m *SmartLimiterStatus) String() string { return proto.CompactTextString(m) }
func (*SmartLimiterStatus) ProtoMessage()    {}
func (*SmartLimiterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{1}
}

func (m *SmartLimiterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimiterStatus.Unmarshal(m, b)
}
func (m *SmartLimiterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimiterStatus.Marshal(b, m, deterministic)
}
func (m *SmartLimiterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimiterStatus.Merge(m, src)
}
func (m *SmartLimiterStatus) XXX_Size() int {
	return xxx_messageInfo_SmartLimiterStatus.Size(m)
}
func (m *SmartLimiterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimiterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimiterStatus proto.InternalMessageInfo

func (m *SmartLimiterStatus) GetRatelimitStatus() map[string]*SmartLimitDescriptors {
	if m != nil {
		return m.RatelimitStatus
	}
	return nil
}

func (m *SmartLimiterStatus) GetEndPointStatus() map[string]string {
	if m != nil {
		return m.EndPointStatus
	}
	return nil
}

func (m *SmartLimiterStatus) GetServiceStatus() *SmartLimiterStatus_ServiceStatus {
	if m != nil {
		return m.ServiceStatus
	}
	return nil
}

type SmartLimiterStatus_ServiceStatus struct {
	Selector             map[string]string                            `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Listener             []*SmartLimiterStatus_ServiceStatus_Listener `protobuf:"bytes,2,rep,name=listener,proto3" json:"listener,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *SmartLimiterStatus_ServiceStatus) Reset()         { *m = SmartLimiterStatus_ServiceStatus{} }
func (m *SmartLimiterStatus_ServiceStatus) String() string { return proto.CompactTextString(m) }
func (*SmartLimiterStatus_ServiceStatus) ProtoMessage()    {}
func (*SmartLimiterStatus_ServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{1, 2}
}

func (m *SmartLimiterStatus_ServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimiterStatus_ServiceStatus.Unmarshal(m, b)
}
func (m *SmartLimiterStatus_ServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimiterStatus_ServiceStatus.Marshal(b, m, deterministic)
}
func (m *SmartLimiterStatus_ServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimiterStatus_ServiceStatus.Merge(m, src)
}
func (m *SmartLimiterStatus_ServiceStatus) XXX_Size() int {
	return xxx_messageInfo_SmartLimiterStatus_ServiceStatus.Size(m)
}
func (m *SmartLimiterStatus_ServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimiterStatus_ServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimiterStatus_ServiceStatus proto.InternalMessageInfo

func (m *SmartLimiterStatus_ServiceStatus) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *SmartLimiterStatus_ServiceStatus) GetListener() []*SmartLimiterStatus_ServiceStatus_Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

type SmartLimiterStatus_ServiceStatus_Listener struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmartLimiterStatus_ServiceStatus_Listener) Reset() {
	*m = SmartLimiterStatus_ServiceStatus_Listener{}
}
func (m *SmartLimiterStatus_ServiceStatus_Listener) String() string { return proto.CompactTextString(m) }
func (*SmartLimiterStatus_ServiceStatus_Listener) ProtoMessage()    {}
func (*SmartLimiterStatus_ServiceStatus_Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{1, 2, 1}
}

func (m *SmartLimiterStatus_ServiceStatus_Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimiterStatus_ServiceStatus_Listener.Unmarshal(m, b)
}
func (m *SmartLimiterStatus_ServiceStatus_Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimiterStatus_ServiceStatus_Listener.Marshal(b, m, deterministic)
}
func (m *SmartLimiterStatus_ServiceStatus_Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimiterStatus_ServiceStatus_Listener.Merge(m, src)
}
func (m *SmartLimiterStatus_ServiceStatus_Listener) XXX_Size() int {
	return xxx_messageInfo_SmartLimiterStatus_ServiceStatus_Listener.Size(m)
}
func (m *SmartLimiterStatus_ServiceStatus_Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimiterStatus_ServiceStatus_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimiterStatus_ServiceStatus_Listener proto.InternalMessageInfo

func (m *SmartLimiterStatus_ServiceStatus_Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SmartLimiterStatus_ServiceStatus_Listener) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type SmartLimitDescriptor struct {
	Condition            string                                `protobuf:"bytes,2,opt,name=condition,proto3" json:"condition,omitempty"`
	Action               *SmartLimitDescriptor_Action          `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Match                []*SmartLimitDescriptor_HeaderMatcher `protobuf:"bytes,4,rep,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SmartLimitDescriptor) Reset()         { *m = SmartLimitDescriptor{} }
func (m *SmartLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*SmartLimitDescriptor) ProtoMessage()    {}
func (*SmartLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{2}
}

func (m *SmartLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimitDescriptor.Unmarshal(m, b)
}
func (m *SmartLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimitDescriptor.Marshal(b, m, deterministic)
}
func (m *SmartLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimitDescriptor.Merge(m, src)
}
func (m *SmartLimitDescriptor) XXX_Size() int {
	return xxx_messageInfo_SmartLimitDescriptor.Size(m)
}
func (m *SmartLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimitDescriptor proto.InternalMessageInfo

func (m *SmartLimitDescriptor) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *SmartLimitDescriptor) GetAction() *SmartLimitDescriptor_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *SmartLimitDescriptor) GetMatch() []*SmartLimitDescriptor_HeaderMatcher {
	if m != nil {
		return m.Match
	}
	return nil
}

type SmartLimitDescriptor_HeaderMatcher struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExactMatch           string   `protobuf:"bytes,2,opt,name=exact_match,json=exactMatch,proto3" json:"exact_match,omitempty"`
	RegexMatch           string   `protobuf:"bytes,3,opt,name=regex_match,json=regexMatch,proto3" json:"regex_match,omitempty"`
	PresentMatch         bool     `protobuf:"varint,7,opt,name=present_match,json=presentMatch,proto3" json:"present_match,omitempty"`
	PrefixMatch          string   `protobuf:"bytes,9,opt,name=prefix_match,json=prefixMatch,proto3" json:"prefix_match,omitempty"`
	SuffixMatch          string   `protobuf:"bytes,10,opt,name=suffix_match,json=suffixMatch,proto3" json:"suffix_match,omitempty"`
	InvertMatch          bool     `protobuf:"varint,8,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmartLimitDescriptor_HeaderMatcher) Reset()         { *m = SmartLimitDescriptor_HeaderMatcher{} }
func (m *SmartLimitDescriptor_HeaderMatcher) String() string { return proto.CompactTextString(m) }
func (*SmartLimitDescriptor_HeaderMatcher) ProtoMessage()    {}
func (*SmartLimitDescriptor_HeaderMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{2, 0}
}

func (m *SmartLimitDescriptor_HeaderMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimitDescriptor_HeaderMatcher.Unmarshal(m, b)
}
func (m *SmartLimitDescriptor_HeaderMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimitDescriptor_HeaderMatcher.Marshal(b, m, deterministic)
}
func (m *SmartLimitDescriptor_HeaderMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimitDescriptor_HeaderMatcher.Merge(m, src)
}
func (m *SmartLimitDescriptor_HeaderMatcher) XXX_Size() int {
	return xxx_messageInfo_SmartLimitDescriptor_HeaderMatcher.Size(m)
}
func (m *SmartLimitDescriptor_HeaderMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimitDescriptor_HeaderMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimitDescriptor_HeaderMatcher proto.InternalMessageInfo

func (m *SmartLimitDescriptor_HeaderMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SmartLimitDescriptor_HeaderMatcher) GetExactMatch() string {
	if m != nil {
		return m.ExactMatch
	}
	return ""
}

func (m *SmartLimitDescriptor_HeaderMatcher) GetRegexMatch() string {
	if m != nil {
		return m.RegexMatch
	}
	return ""
}

func (m *SmartLimitDescriptor_HeaderMatcher) GetPresentMatch() bool {
	if m != nil {
		return m.PresentMatch
	}
	return false
}

func (m *SmartLimitDescriptor_HeaderMatcher) GetPrefixMatch() string {
	if m != nil {
		return m.PrefixMatch
	}
	return ""
}

func (m *SmartLimitDescriptor_HeaderMatcher) GetSuffixMatch() string {
	if m != nil {
		return m.SuffixMatch
	}
	return ""
}

func (m *SmartLimitDescriptor_HeaderMatcher) GetInvertMatch() bool {
	if m != nil {
		return m.InvertMatch
	}
	return false
}

type SmartLimitDescriptor_Action struct {
	Quota                string          `protobuf:"bytes,1,opt,name=quota,proto3" json:"quota,omitempty"`
	FillInterval         *types.Duration `protobuf:"bytes,2,opt,name=fill_interval,json=fillInterval,proto3" json:"fill_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SmartLimitDescriptor_Action) Reset()         { *m = SmartLimitDescriptor_Action{} }
func (m *SmartLimitDescriptor_Action) String() string { return proto.CompactTextString(m) }
func (*SmartLimitDescriptor_Action) ProtoMessage()    {}
func (*SmartLimitDescriptor_Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{2, 1}
}

func (m *SmartLimitDescriptor_Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimitDescriptor_Action.Unmarshal(m, b)
}
func (m *SmartLimitDescriptor_Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimitDescriptor_Action.Marshal(b, m, deterministic)
}
func (m *SmartLimitDescriptor_Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimitDescriptor_Action.Merge(m, src)
}
func (m *SmartLimitDescriptor_Action) XXX_Size() int {
	return xxx_messageInfo_SmartLimitDescriptor_Action.Size(m)
}
func (m *SmartLimitDescriptor_Action) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimitDescriptor_Action.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimitDescriptor_Action proto.InternalMessageInfo

func (m *SmartLimitDescriptor_Action) GetQuota() string {
	if m != nil {
		return m.Quota
	}
	return ""
}

func (m *SmartLimitDescriptor_Action) GetFillInterval() *types.Duration {
	if m != nil {
		return m.FillInterval
	}
	return nil
}

type SmartLimitDescriptors struct {
	Descriptor_          []*SmartLimitDescriptor `protobuf:"bytes,4,rep,name=descriptor,proto3" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SmartLimitDescriptors) Reset()         { *m = SmartLimitDescriptors{} }
func (m *SmartLimitDescriptors) String() string { return proto.CompactTextString(m) }
func (*SmartLimitDescriptors) ProtoMessage()    {}
func (*SmartLimitDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_76bfc6c51416fcda, []int{3}
}

func (m *SmartLimitDescriptors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartLimitDescriptors.Unmarshal(m, b)
}
func (m *SmartLimitDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartLimitDescriptors.Marshal(b, m, deterministic)
}
func (m *SmartLimitDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartLimitDescriptors.Merge(m, src)
}
func (m *SmartLimitDescriptors) XXX_Size() int {
	return xxx_messageInfo_SmartLimitDescriptors.Size(m)
}
func (m *SmartLimitDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartLimitDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_SmartLimitDescriptors proto.InternalMessageInfo

func (m *SmartLimitDescriptors) GetDescriptor_() []*SmartLimitDescriptor {
	if m != nil {
		return m.Descriptor_
	}
	return nil
}

func init() {
	proto.RegisterType((*SmartLimiterSpec)(nil), "netease.microservice.v1alpha1.SmartLimiterSpec")
	proto.RegisterMapType((map[string]*SmartLimitDescriptors)(nil), "netease.microservice.v1alpha1.SmartLimiterSpec.SetsEntry")
	proto.RegisterType((*SmartLimiterStatus)(nil), "netease.microservice.v1alpha1.SmartLimiterStatus")
	proto.RegisterMapType((map[string]string)(nil), "netease.microservice.v1alpha1.SmartLimiterStatus.EndPointStatusEntry")
	proto.RegisterMapType((map[string]*SmartLimitDescriptors)(nil), "netease.microservice.v1alpha1.SmartLimiterStatus.RatelimitStatusEntry")
	proto.RegisterType((*SmartLimiterStatus_ServiceStatus)(nil), "netease.microservice.v1alpha1.SmartLimiterStatus.ServiceStatus")
	proto.RegisterMapType((map[string]string)(nil), "netease.microservice.v1alpha1.SmartLimiterStatus.ServiceStatus.SelectorEntry")
	proto.RegisterType((*SmartLimiterStatus_ServiceStatus_Listener)(nil), "netease.microservice.v1alpha1.SmartLimiterStatus.ServiceStatus.Listener")
	proto.RegisterType((*SmartLimitDescriptor)(nil), "netease.microservice.v1alpha1.SmartLimitDescriptor")
	proto.RegisterType((*SmartLimitDescriptor_HeaderMatcher)(nil), "netease.microservice.v1alpha1.SmartLimitDescriptor.HeaderMatcher")
	proto.RegisterType((*SmartLimitDescriptor_Action)(nil), "netease.microservice.v1alpha1.SmartLimitDescriptor.Action")
	proto.RegisterType((*SmartLimitDescriptors)(nil), "netease.microservice.v1alpha1.SmartLimitDescriptors")
}

func init() {
	proto.RegisterFile("pkg/apis/microservice/v1alpha1/smart_limiter.proto", fileDescriptor_76bfc6c51416fcda)
}

var fileDescriptor_76bfc6c51416fcda = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x95, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0xb5, 0x2b, 0xed, 0xe9, 0x0a, 0x93, 0x19, 0x68, 0x54, 0x7c, 0x8c, 0x72, 0xb3,
	0xab, 0x44, 0xfb, 0x10, 0x82, 0x21, 0x81, 0x86, 0x36, 0x34, 0xd0, 0x26, 0xa1, 0xf4, 0x02, 0x89,
	0x0b, 0x26, 0x2f, 0x39, 0xed, 0xac, 0x25, 0x71, 0xb0, 0x9d, 0x6a, 0x7b, 0x32, 0x9e, 0x83, 0xb7,
	0xe0, 0x86, 0x2b, 0x5e, 0x00, 0xc7, 0x76, 0xb2, 0x76, 0xaa, 0x26, 0xba, 0x89, 0x3b, 0xfb, 0xe4,
	0x7f, 0x7e, 0x7f, 0x9f, 0x63, 0x3b, 0x86, 0xcd, 0xec, 0x6c, 0xe4, 0xd3, 0x8c, 0x49, 0x3f, 0x61,
	0xa1, 0xe0, 0x12, 0xc5, 0x98, 0x85, 0xe8, 0x8f, 0x37, 0x68, 0x9c, 0x9d, 0xd2, 0x0d, 0x5f, 0x26,
	0x54, 0xa8, 0xe3, 0x98, 0x25, 0x4c, 0xa1, 0xf0, 0x32, 0xc1, 0x15, 0x27, 0x4f, 0x52, 0x54, 0x48,
	0x25, 0x7a, 0x93, 0x29, 0x5e, 0x99, 0xd2, 0x7b, 0x3a, 0xe2, 0x7c, 0x14, 0xa3, 0x6f, 0xc4, 0x27,
	0xf9, 0xd0, 0x8f, 0x72, 0x41, 0x15, 0xe3, 0xa9, 0x4d, 0xef, 0xff, 0xaa, 0xc1, 0xf2, 0xa0, 0xc0,
	0x1e, 0x5a, 0xea, 0x20, 0xc3, 0x90, 0x3c, 0x84, 0x66, 0xc4, 0x13, 0xca, 0xd2, 0xd5, 0xda, 0x5a,
	0x6d, 0xbd, 0x1d, 0xb8, 0x19, 0x39, 0x82, 0x86, 0x44, 0x25, 0x57, 0x17, 0xd6, 0xea, 0xeb, 0x9d,
	0xcd, 0xd7, 0xde, 0xb5, 0xd6, 0xde, 0x55, 0xac, 0x37, 0xd0, 0xb9, 0xfb, 0xa9, 0x12, 0x17, 0x81,
	0xc1, 0xf4, 0x12, 0x68, 0x57, 0x21, 0xb2, 0x0c, 0xf5, 0x33, 0xbc, 0x70, 0x86, 0xc5, 0x90, 0x7c,
	0x82, 0xc5, 0x31, 0x8d, 0x73, 0xd4, 0x76, 0x35, 0x6d, 0xb7, 0xfd, 0xcf, 0x76, 0x7b, 0x28, 0x43,
	0xc1, 0x32, 0xc5, 0x85, 0x0c, 0x2c, 0x62, 0x67, 0xe1, 0x55, 0xad, 0xff, 0xbb, 0x09, 0x64, 0x6a,
	0x4d, 0x8a, 0xaa, 0x5c, 0x92, 0x0c, 0xee, 0xe9, 0x8e, 0xa0, 0xe9, 0xaa, 0x0d, 0xe9, 0x45, 0x14,
	0xf5, 0x7d, 0x98, 0xa7, 0x3e, 0x93, 0xe8, 0x05, 0xd3, 0x20, 0x5b, 0xec, 0x55, 0x3c, 0x49, 0xe0,
	0x2e, 0xa6, 0xd1, 0x67, 0xce, 0xd2, 0xd2, 0xd0, 0x36, 0x74, 0x7f, 0x7e, 0xc3, 0xfd, 0x29, 0x8e,
	0xf5, 0xbb, 0x02, 0x27, 0x08, 0x5d, 0x87, 0x72, 0x6e, 0x75, 0xd3, 0xcf, 0x77, 0xf3, 0xbb, 0x0d,
	0x26, 0x31, 0xc1, 0x34, 0xb5, 0x77, 0x0e, 0x2b, 0xb3, 0xca, 0xff, 0xff, 0x1b, 0xdb, 0xdb, 0x85,
	0xfb, 0x33, 0xfa, 0x30, 0xc3, 0x78, 0x65, 0xd2, 0xb8, 0x3d, 0x89, 0xf8, 0xb9, 0x00, 0xdd, 0xa9,
	0xea, 0x08, 0x83, 0x96, 0xc4, 0x18, 0x43, 0x6d, 0xe6, 0xce, 0xc3, 0xd1, 0x2d, 0x1b, 0xa6, 0x67,
	0x96, 0x67, 0xb7, 0xa9, 0xc2, 0x93, 0x08, 0x5a, 0x31, 0x93, 0x0a, 0x53, 0x14, 0xee, 0x24, 0x1c,
	0xdc, 0xd6, 0xea, 0xd0, 0xf1, 0x82, 0x8a, 0xdc, 0x7b, 0x53, 0x54, 0x38, 0xb1, 0x80, 0xb9, 0xfa,
	0xb3, 0x09, 0xad, 0x12, 0x49, 0x08, 0x34, 0x52, 0x9a, 0xa0, 0x4b, 0x34, 0xe3, 0x22, 0x96, 0x71,
	0xa1, 0x4c, 0x62, 0x37, 0x30, 0xe3, 0xfe, 0x8f, 0x06, 0xac, 0xcc, 0xda, 0x3b, 0xf2, 0x18, 0xda,
	0x21, 0x4f, 0x23, 0x56, 0xfc, 0x86, 0x9c, 0xd5, 0x65, 0x80, 0x04, 0xd0, 0xa4, 0xa1, 0xf9, 0x64,
	0xcf, 0xe9, 0xce, 0x0d, 0x8e, 0x87, 0xb7, 0x6b, 0x08, 0x81, 0x23, 0x91, 0x2f, 0xb0, 0x98, 0x50,
	0x15, 0x9e, 0xae, 0x36, 0x4c, 0x7b, 0x77, 0x6f, 0x82, 0x3c, 0x40, 0x1a, 0xa1, 0x38, 0x2a, 0x30,
	0xba, 0xaf, 0x96, 0xd7, 0xfb, 0x53, 0x83, 0xee, 0xd4, 0x87, 0x99, 0xdd, 0x79, 0x06, 0x1d, 0x3c,
	0xd7, 0x4b, 0x39, 0xb6, 0x8b, 0xb0, 0x25, 0x83, 0x09, 0x99, 0xb4, 0x42, 0x20, 0x70, 0x84, 0xe7,
	0x4e, 0x50, 0xb7, 0x02, 0x13, 0xb2, 0x82, 0x17, 0xd0, 0xcd, 0x04, 0x4a, 0x4c, 0x4b, 0xc6, 0x1d,
	0x2d, 0x69, 0x05, 0x4b, 0x2e, 0x68, 0x45, 0xcf, 0xa1, 0x98, 0x0f, 0x59, 0x89, 0x69, 0x1b, 0x4c,
	0xc7, 0xc6, 0x2a, 0x89, 0xcc, 0x87, 0x97, 0x12, 0xb0, 0x12, 0x1b, 0xab, 0x24, 0x2c, 0x1d, 0xa3,
	0x28, 0x9d, 0x5a, 0xc6, 0xa9, 0x63, 0x63, 0x46, 0xd2, 0xfb, 0x06, 0x4d, 0xdb, 0xe0, 0xe2, 0xc4,
	0x7c, 0xcf, 0xb9, 0xa2, 0xae, 0x5c, 0x3b, 0x21, 0x6f, 0xa1, 0x3b, 0x64, 0x71, 0x7c, 0xac, 0x2f,
	0xa4, 0x6e, 0x2d, 0x8d, 0xdd, 0x45, 0x7f, 0xe4, 0xd9, 0xc7, 0xc8, 0x2b, 0x1f, 0x23, 0x6f, 0xcf,
	0x3d, 0x46, 0xc1, 0x52, 0xa1, 0xff, 0xe8, 0xe4, 0xfd, 0x18, 0x1e, 0xcc, 0xbc, 0xf4, 0x64, 0x00,
	0x10, 0x55, 0x53, 0xb7, 0x99, 0x5b, 0x37, 0xd8, 0xcc, 0x60, 0x02, 0xf3, 0xfe, 0xe5, 0xd7, 0xed,
	0x8b, 0x3c, 0xad, 0x28, 0x21, 0x4f, 0x7c, 0xa9, 0xff, 0x62, 0xfa, 0xcd, 0xbc, 0xf6, 0x35, 0x3e,
	0x69, 0x9a, 0x32, 0xb6, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x05, 0x66, 0x61, 0xe8, 0xb6, 0x07,
	0x00, 0x00,
}
