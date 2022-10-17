// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.6.0
//  VPP:              22.06-release
// source: /usr/share/vpp/api/plugins/mdata.api.json

// Package mdata contains generated bindings for API file mdata.api.
//
// Contents:
// -  2 messages
package mdata

import (
	api "go.fd.io/govpp/api"
	_ "go.fd.io/govpp/binapi/interface_types"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "mdata"
	APIVersion = "0.1.0"
	VersionCrc = 0x5bd69477
)

// MdataEnableDisable defines message 'mdata_enable_disable'.
// InProgress: the message form may change in the future versions
type MdataEnableDisable struct {
	EnableDisable bool `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
}

func (m *MdataEnableDisable) Reset()               { *m = MdataEnableDisable{} }
func (*MdataEnableDisable) GetMessageName() string { return "mdata_enable_disable" }
func (*MdataEnableDisable) GetCrcString() string   { return "2e7b47df" }
func (*MdataEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessageType
}

func (m *MdataEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	return size
}
func (m *MdataEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	return buf.Bytes(), nil
}
func (m *MdataEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	return nil
}

// MdataEnableDisableReply defines message 'mdata_enable_disable_reply'.
// InProgress: the message form may change in the future versions
type MdataEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MdataEnableDisableReply) Reset()               { *m = MdataEnableDisableReply{} }
func (*MdataEnableDisableReply) GetMessageName() string { return "mdata_enable_disable_reply" }
func (*MdataEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*MdataEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessageType
}

func (m *MdataEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MdataEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MdataEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_mdata_binapi_init() }
func file_mdata_binapi_init() {
	api.RegisterMessage((*MdataEnableDisable)(nil), "mdata_enable_disable_2e7b47df")
	api.RegisterMessage((*MdataEnableDisableReply)(nil), "mdata_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MdataEnableDisable)(nil),
		(*MdataEnableDisableReply)(nil),
	}
}
