// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.5.0
//  VPP:              22.02-release
// source: /usr/share/vpp/api/core/flow.api.json

// Package flow contains generated bindings for API file flow.api.
//
// Contents:
//   8 messages
//
package flow

import (
	api "git.fd.io/govpp.git/api"
	_ "git.fd.io/govpp.git/binapi/ethernet_types"
	flow_types "git.fd.io/govpp.git/binapi/flow_types"
	_ "git.fd.io/govpp.git/binapi/interface_types"
	_ "git.fd.io/govpp.git/binapi/ip_types"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "flow"
	APIVersion = "0.0.2"
	VersionCrc = 0x140d3585
)

// FlowAdd defines message 'flow_add'.
// InProgress: the message form may change in the future versions
type FlowAdd struct {
	Flow flow_types.FlowRule `binapi:"flow_rule,name=flow" json:"flow,omitempty"`
}

func (m *FlowAdd) Reset()               { *m = FlowAdd{} }
func (*FlowAdd) GetMessageName() string { return "flow_add" }
func (*FlowAdd) GetCrcString() string   { return "f946ed84" }
func (*FlowAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Flow.Type
	size += 4      // m.Flow.Index
	size += 4      // m.Flow.Actions
	size += 4      // m.Flow.MarkFlowID
	size += 4      // m.Flow.RedirectNodeIndex
	size += 4      // m.Flow.RedirectDeviceInputNextIndex
	size += 4      // m.Flow.RedirectQueue
	size += 4      // m.Flow.BufferAdvance
	size += 1 * 82 // m.Flow.Flow
	return size
}
func (m *FlowAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Flow.Type))
	buf.EncodeUint32(m.Flow.Index)
	buf.EncodeUint32(uint32(m.Flow.Actions))
	buf.EncodeUint32(m.Flow.MarkFlowID)
	buf.EncodeUint32(m.Flow.RedirectNodeIndex)
	buf.EncodeUint32(m.Flow.RedirectDeviceInputNextIndex)
	buf.EncodeUint32(m.Flow.RedirectQueue)
	buf.EncodeInt32(m.Flow.BufferAdvance)
	buf.EncodeBytes(m.Flow.Flow.XXX_UnionData[:], 82)
	return buf.Bytes(), nil
}
func (m *FlowAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Flow.Type = flow_types.FlowType(buf.DecodeUint32())
	m.Flow.Index = buf.DecodeUint32()
	m.Flow.Actions = flow_types.FlowAction(buf.DecodeUint32())
	m.Flow.MarkFlowID = buf.DecodeUint32()
	m.Flow.RedirectNodeIndex = buf.DecodeUint32()
	m.Flow.RedirectDeviceInputNextIndex = buf.DecodeUint32()
	m.Flow.RedirectQueue = buf.DecodeUint32()
	m.Flow.BufferAdvance = buf.DecodeInt32()
	copy(m.Flow.Flow.XXX_UnionData[:], buf.DecodeBytes(82))
	return nil
}

// FlowAddReply defines message 'flow_add_reply'.
// InProgress: the message form may change in the future versions
type FlowAddReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	FlowIndex uint32 `binapi:"u32,name=flow_index" json:"flow_index,omitempty"`
}

func (m *FlowAddReply) Reset()               { *m = FlowAddReply{} }
func (*FlowAddReply) GetMessageName() string { return "flow_add_reply" }
func (*FlowAddReply) GetCrcString() string   { return "8587dc85" }
func (*FlowAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.FlowIndex
	return size
}
func (m *FlowAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.FlowIndex)
	return buf.Bytes(), nil
}
func (m *FlowAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.FlowIndex = buf.DecodeUint32()
	return nil
}

// FlowDel defines message 'flow_del'.
// InProgress: the message form may change in the future versions
type FlowDel struct {
	FlowIndex uint32 `binapi:"u32,name=flow_index" json:"flow_index,omitempty"`
}

func (m *FlowDel) Reset()               { *m = FlowDel{} }
func (*FlowDel) GetMessageName() string { return "flow_del" }
func (*FlowDel) GetCrcString() string   { return "b6b9b02c" }
func (*FlowDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.FlowIndex
	return size
}
func (m *FlowDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.FlowIndex)
	return buf.Bytes(), nil
}
func (m *FlowDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.FlowIndex = buf.DecodeUint32()
	return nil
}

// FlowDelReply defines message 'flow_del_reply'.
// InProgress: the message form may change in the future versions
type FlowDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowDelReply) Reset()               { *m = FlowDelReply{} }
func (*FlowDelReply) GetMessageName() string { return "flow_del_reply" }
func (*FlowDelReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// FlowDisable defines message 'flow_disable'.
// InProgress: the message form may change in the future versions
type FlowDisable struct {
	FlowIndex uint32 `binapi:"u32,name=flow_index" json:"flow_index,omitempty"`
	HwIfIndex uint32 `binapi:"u32,name=hw_if_index" json:"hw_if_index,omitempty"`
}

func (m *FlowDisable) Reset()               { *m = FlowDisable{} }
func (*FlowDisable) GetMessageName() string { return "flow_disable" }
func (*FlowDisable) GetCrcString() string   { return "2024be69" }
func (*FlowDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.FlowIndex
	size += 4 // m.HwIfIndex
	return size
}
func (m *FlowDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.FlowIndex)
	buf.EncodeUint32(m.HwIfIndex)
	return buf.Bytes(), nil
}
func (m *FlowDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.FlowIndex = buf.DecodeUint32()
	m.HwIfIndex = buf.DecodeUint32()
	return nil
}

// FlowDisableReply defines message 'flow_disable_reply'.
// InProgress: the message form may change in the future versions
type FlowDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowDisableReply) Reset()               { *m = FlowDisableReply{} }
func (*FlowDisableReply) GetMessageName() string { return "flow_disable_reply" }
func (*FlowDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// FlowEnable defines message 'flow_enable'.
// InProgress: the message form may change in the future versions
type FlowEnable struct {
	FlowIndex uint32 `binapi:"u32,name=flow_index" json:"flow_index,omitempty"`
	HwIfIndex uint32 `binapi:"u32,name=hw_if_index" json:"hw_if_index,omitempty"`
}

func (m *FlowEnable) Reset()               { *m = FlowEnable{} }
func (*FlowEnable) GetMessageName() string { return "flow_enable" }
func (*FlowEnable) GetCrcString() string   { return "2024be69" }
func (*FlowEnable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FlowEnable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.FlowIndex
	size += 4 // m.HwIfIndex
	return size
}
func (m *FlowEnable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.FlowIndex)
	buf.EncodeUint32(m.HwIfIndex)
	return buf.Bytes(), nil
}
func (m *FlowEnable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.FlowIndex = buf.DecodeUint32()
	m.HwIfIndex = buf.DecodeUint32()
	return nil
}

// FlowEnableReply defines message 'flow_enable_reply'.
// InProgress: the message form may change in the future versions
type FlowEnableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FlowEnableReply) Reset()               { *m = FlowEnableReply{} }
func (*FlowEnableReply) GetMessageName() string { return "flow_enable_reply" }
func (*FlowEnableReply) GetCrcString() string   { return "e8d4e804" }
func (*FlowEnableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FlowEnableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FlowEnableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FlowEnableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_flow_binapi_init() }
func file_flow_binapi_init() {
	api.RegisterMessage((*FlowAdd)(nil), "flow_add_f946ed84")
	api.RegisterMessage((*FlowAddReply)(nil), "flow_add_reply_8587dc85")
	api.RegisterMessage((*FlowDel)(nil), "flow_del_b6b9b02c")
	api.RegisterMessage((*FlowDelReply)(nil), "flow_del_reply_e8d4e804")
	api.RegisterMessage((*FlowDisable)(nil), "flow_disable_2024be69")
	api.RegisterMessage((*FlowDisableReply)(nil), "flow_disable_reply_e8d4e804")
	api.RegisterMessage((*FlowEnable)(nil), "flow_enable_2024be69")
	api.RegisterMessage((*FlowEnableReply)(nil), "flow_enable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FlowAdd)(nil),
		(*FlowAddReply)(nil),
		(*FlowDel)(nil),
		(*FlowDelReply)(nil),
		(*FlowDisable)(nil),
		(*FlowDisableReply)(nil),
		(*FlowEnable)(nil),
		(*FlowEnableReply)(nil),
	}
}
