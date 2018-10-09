/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pod_affinity.proto

package firmament

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchLabels struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchLabels) Reset()         { *m = MatchLabels{} }
func (m *MatchLabels) String() string { return proto.CompactTextString(m) }
func (*MatchLabels) ProtoMessage()    {}
func (*MatchLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_f260e0ec83421d4c, []int{0}
}

func (m *MatchLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchLabels.Unmarshal(m, b)
}
func (m *MatchLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchLabels.Marshal(b, m, deterministic)
}
func (m *MatchLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchLabels.Merge(m, src)
}
func (m *MatchLabels) XXX_Size() int {
	return xxx_messageInfo_MatchLabels.Size(m)
}
func (m *MatchLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchLabels.DiscardUnknown(m)
}

var xxx_messageInfo_MatchLabels proto.InternalMessageInfo

func (m *MatchLabels) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MatchLabels) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type LabelSelectorRequirement struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Values               []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelSelectorRequirement) Reset()         { *m = LabelSelectorRequirement{} }
func (m *LabelSelectorRequirement) String() string { return proto.CompactTextString(m) }
func (*LabelSelectorRequirement) ProtoMessage()    {}
func (*LabelSelectorRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_f260e0ec83421d4c, []int{1}
}

func (m *LabelSelectorRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelSelectorRequirement.Unmarshal(m, b)
}
func (m *LabelSelectorRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelSelectorRequirement.Marshal(b, m, deterministic)
}
func (m *LabelSelectorRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelectorRequirement.Merge(m, src)
}
func (m *LabelSelectorRequirement) XXX_Size() int {
	return xxx_messageInfo_LabelSelectorRequirement.Size(m)
}
func (m *LabelSelectorRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelectorRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelectorRequirement proto.InternalMessageInfo

func (m *LabelSelectorRequirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelSelectorRequirement) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *LabelSelectorRequirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type PodLabelSelector struct {
	MatchLabels          *MatchLabels                `protobuf:"bytes,1,opt,name=matchLabels,proto3" json:"matchLabels,omitempty"`
	MatchExpressions     []*LabelSelectorRequirement `protobuf:"bytes,2,rep,name=matchExpressions,proto3" json:"matchExpressions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PodLabelSelector) Reset()         { *m = PodLabelSelector{} }
func (m *PodLabelSelector) String() string { return proto.CompactTextString(m) }
func (*PodLabelSelector) ProtoMessage()    {}
func (*PodLabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_f260e0ec83421d4c, []int{2}
}

func (m *PodLabelSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodLabelSelector.Unmarshal(m, b)
}
func (m *PodLabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodLabelSelector.Marshal(b, m, deterministic)
}
func (m *PodLabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodLabelSelector.Merge(m, src)
}
func (m *PodLabelSelector) XXX_Size() int {
	return xxx_messageInfo_PodLabelSelector.Size(m)
}
func (m *PodLabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_PodLabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_PodLabelSelector proto.InternalMessageInfo

func (m *PodLabelSelector) GetMatchLabels() *MatchLabels {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

func (m *PodLabelSelector) GetMatchExpressions() []*LabelSelectorRequirement {
	if m != nil {
		return m.MatchExpressions
	}
	return nil
}

type PodAffinityTerm struct {
	LabelSelector *PodLabelSelector `protobuf:"bytes,1,opt,name=labelSelector,proto3" json:"labelSelector,omitempty"`
	Namespaces    []string          `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// Empty topologyKey is not allowed.
	TopologyKey          string   `protobuf:"bytes,3,opt,name=topologyKey,proto3" json:"topologyKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PodAffinityTerm) Reset()         { *m = PodAffinityTerm{} }
func (m *PodAffinityTerm) String() string { return proto.CompactTextString(m) }
func (*PodAffinityTerm) ProtoMessage()    {}
func (*PodAffinityTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_f260e0ec83421d4c, []int{3}
}

func (m *PodAffinityTerm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodAffinityTerm.Unmarshal(m, b)
}
func (m *PodAffinityTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodAffinityTerm.Marshal(b, m, deterministic)
}
func (m *PodAffinityTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodAffinityTerm.Merge(m, src)
}
func (m *PodAffinityTerm) XXX_Size() int {
	return xxx_messageInfo_PodAffinityTerm.Size(m)
}
func (m *PodAffinityTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_PodAffinityTerm.DiscardUnknown(m)
}

var xxx_messageInfo_PodAffinityTerm proto.InternalMessageInfo

func (m *PodAffinityTerm) GetLabelSelector() *PodLabelSelector {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

func (m *PodAffinityTerm) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *PodAffinityTerm) GetTopologyKey() string {
	if m != nil {
		return m.TopologyKey
	}
	return ""
}

type WeightedPodAffinityTerm struct {
	// weight associated with matching the corresponding podAffinityTerm,
	// in the range 1-100.
	Weight int32 `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	// A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm      *PodAffinityTerm `protobuf:"bytes,2,opt,name=podAffinityTerm,proto3" json:"podAffinityTerm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WeightedPodAffinityTerm) Reset()         { *m = WeightedPodAffinityTerm{} }
func (m *WeightedPodAffinityTerm) String() string { return proto.CompactTextString(m) }
func (*WeightedPodAffinityTerm) ProtoMessage()    {}
func (*WeightedPodAffinityTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_f260e0ec83421d4c, []int{4}
}

func (m *WeightedPodAffinityTerm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedPodAffinityTerm.Unmarshal(m, b)
}
func (m *WeightedPodAffinityTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedPodAffinityTerm.Marshal(b, m, deterministic)
}
func (m *WeightedPodAffinityTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedPodAffinityTerm.Merge(m, src)
}
func (m *WeightedPodAffinityTerm) XXX_Size() int {
	return xxx_messageInfo_WeightedPodAffinityTerm.Size(m)
}
func (m *WeightedPodAffinityTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedPodAffinityTerm.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedPodAffinityTerm proto.InternalMessageInfo

func (m *WeightedPodAffinityTerm) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *WeightedPodAffinityTerm) GetPodAffinityTerm() *PodAffinityTerm {
	if m != nil {
		return m.PodAffinityTerm
	}
	return nil
}

type PodAffinity struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*PodAffinityTerm         `protobuf:"bytes,1,rep,name=requiredDuringSchedulingIgnoredDuringExecution,proto3" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []*WeightedPodAffinityTerm `protobuf:"bytes,2,rep,name=preferredDuringSchedulingIgnoredDuringExecution,proto3" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
	XXX_NoUnkeyedLiteral                            struct{}                   `json:"-"`
	XXX_unrecognized                                []byte                     `json:"-"`
	XXX_sizecache                                   int32                      `json:"-"`
}

func (m *PodAffinity) Reset()         { *m = PodAffinity{} }
func (m *PodAffinity) String() string { return proto.CompactTextString(m) }
func (*PodAffinity) ProtoMessage()    {}
func (*PodAffinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f260e0ec83421d4c, []int{5}
}

func (m *PodAffinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodAffinity.Unmarshal(m, b)
}
func (m *PodAffinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodAffinity.Marshal(b, m, deterministic)
}
func (m *PodAffinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodAffinity.Merge(m, src)
}
func (m *PodAffinity) XXX_Size() int {
	return xxx_messageInfo_PodAffinity.Size(m)
}
func (m *PodAffinity) XXX_DiscardUnknown() {
	xxx_messageInfo_PodAffinity.DiscardUnknown(m)
}

var xxx_messageInfo_PodAffinity proto.InternalMessageInfo

func (m *PodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() []*PodAffinityTerm {
	if m != nil {
		return m.RequiredDuringSchedulingIgnoredDuringExecution
	}
	return nil
}

func (m *PodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []*WeightedPodAffinityTerm {
	if m != nil {
		return m.PreferredDuringSchedulingIgnoredDuringExecution
	}
	return nil
}

func init() {
	proto.RegisterType((*MatchLabels)(nil), "firmament.MatchLabels")
	proto.RegisterType((*LabelSelectorRequirement)(nil), "firmament.LabelSelectorRequirement")
	proto.RegisterType((*PodLabelSelector)(nil), "firmament.PodLabelSelector")
	proto.RegisterType((*PodAffinityTerm)(nil), "firmament.PodAffinityTerm")
	proto.RegisterType((*WeightedPodAffinityTerm)(nil), "firmament.WeightedPodAffinityTerm")
	proto.RegisterType((*PodAffinity)(nil), "firmament.PodAffinity")
}

func init() { proto.RegisterFile("pod_affinity.proto", fileDescriptor_f260e0ec83421d4c) }

var fileDescriptor_f260e0ec83421d4c = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0x0d, 0x5b, 0xcc, 0x0b, 0xb2, 0x65, 0x90, 0x1a, 0x56, 0x90, 0x10, 0x2f, 0x3d, 0x45,
	0xa8, 0x08, 0x5e, 0x17, 0x76, 0x0f, 0xa2, 0xe2, 0x32, 0x15, 0x3c, 0xea, 0x34, 0x79, 0x49, 0x07,
	0x93, 0x99, 0x71, 0x32, 0xb1, 0xcd, 0xd5, 0xb3, 0x67, 0x6f, 0x1e, 0xfc, 0xa7, 0xd2, 0x49, 0xa8,
	0x93, 0x6a, 0xc5, 0xde, 0xf2, 0xbd, 0xbc, 0xef, 0x7b, 0xdf, 0xfb, 0x5e, 0x02, 0x44, 0xc9, 0xfc,
	0x03, 0x2b, 0x0a, 0x2e, 0xb8, 0xe9, 0x52, 0xa5, 0xa5, 0x91, 0x24, 0x28, 0xb8, 0xae, 0x59, 0x8d,
	0xc2, 0x24, 0xcf, 0x21, 0x7c, 0xc3, 0x4c, 0xb6, 0x79, 0xcd, 0xd6, 0x58, 0x35, 0x64, 0x06, 0xfe,
	0x27, 0xec, 0x22, 0x2f, 0xf6, 0x16, 0x01, 0xdd, 0x3f, 0x92, 0x07, 0x70, 0xf1, 0x85, 0x55, 0x2d,
	0x46, 0x13, 0x5b, 0xeb, 0x41, 0xf2, 0x11, 0x22, 0xcb, 0x58, 0x61, 0x85, 0x99, 0x91, 0x9a, 0xe2,
	0xe7, 0x96, 0x6b, 0xdc, 0x4b, 0xfe, 0x45, 0xe3, 0x0a, 0xee, 0x49, 0x85, 0x9a, 0x19, 0xa9, 0x07,
	0x99, 0x03, 0x26, 0x73, 0x98, 0x5a, 0xc9, 0x26, 0xf2, 0x63, 0x7f, 0x11, 0xd0, 0x01, 0x25, 0x3f,
	0x3c, 0x98, 0xdd, 0xc9, 0x7c, 0x34, 0x85, 0xbc, 0x80, 0xb0, 0xfe, 0xed, 0xd6, 0x8e, 0x08, 0x97,
	0xf3, 0xf4, 0xb0, 0x4e, 0xea, 0xec, 0x42, 0xdd, 0x56, 0xf2, 0x16, 0x66, 0x16, 0xde, 0xee, 0x94,
	0xc6, 0xa6, 0xe1, 0x52, 0x34, 0xd1, 0x24, 0xf6, 0x17, 0xe1, 0xf2, 0x89, 0x43, 0x3f, 0xb5, 0x13,
	0xfd, 0x83, 0x9c, 0x7c, 0xf7, 0xe0, 0xf2, 0x4e, 0xe6, 0xd7, 0x43, 0xb2, 0xef, 0x50, 0xd7, 0xe4,
	0x1a, 0xee, 0x57, 0xae, 0xc2, 0x60, 0xf0, 0x91, 0x33, 0xe1, 0x78, 0x25, 0x3a, 0x66, 0x90, 0xc7,
	0x00, 0x82, 0xd5, 0xd8, 0x28, 0x96, 0x61, 0xef, 0x30, 0xa0, 0x4e, 0x85, 0xc4, 0x10, 0x1a, 0xa9,
	0x64, 0x25, 0xcb, 0xee, 0x15, 0x76, 0x91, 0x6f, 0xd3, 0x74, 0x4b, 0xc9, 0x16, 0x1e, 0xbe, 0x47,
	0x5e, 0x6e, 0x0c, 0xe6, 0xc7, 0xfe, 0xe6, 0x30, 0xdd, 0xda, 0x57, 0xd6, 0xd8, 0x05, 0x1d, 0x10,
	0xb9, 0x81, 0x4b, 0x35, 0x6e, 0xb5, 0x67, 0x0a, 0x97, 0x57, 0x63, 0xe7, 0x6e, 0x07, 0x3d, 0xa6,
	0x24, 0x3f, 0x27, 0x10, 0x3a, 0x4d, 0xe4, 0xab, 0x07, 0xa9, 0xee, 0x33, 0xcc, 0x6f, 0x5a, 0xcd,
	0x45, 0xb9, 0xca, 0x36, 0x98, 0xb7, 0x15, 0x17, 0xe5, 0xcb, 0x52, 0xc8, 0x43, 0xf9, 0x76, 0x87,
	0x59, 0x6b, 0xb8, 0x14, 0x91, 0x67, 0x2f, 0xf2, 0xaf, 0xa9, 0x67, 0x2a, 0x92, 0x6f, 0x1e, 0x3c,
	0x55, 0x1a, 0x0b, 0xd4, 0xff, 0xef, 0xa2, 0xff, 0x2e, 0x12, 0xc7, 0xc5, 0x89, 0x40, 0xe9, 0xb9,
	0xd2, 0xeb, 0xa9, 0xfd, 0x01, 0x9f, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x1d, 0x62, 0x1e,
	0x96, 0x03, 0x00, 0x00,
}
