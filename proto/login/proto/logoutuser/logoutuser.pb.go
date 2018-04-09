// Code generated by protoc-gen-go.
// source: github.com/hailocab/h2/proto/login/proto/logoutuser/logoutuser.proto
// DO NOT EDIT!

/*
Package com_hailocab_service_login_logoutuser is a generated protocol buffer package.

It is generated from these files:
	github.com/hailocab/h2/proto/login/proto/logoutuser/logoutuser.proto

It has these top-level messages:
	Request
	Response
*/
package com_hailocab_service_login_logoutuser

import proto "github.com/hailocab/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	Mech             *string `protobuf:"bytes,1,req,name=mech" json:"mech,omitempty"`
	DeviceType       *string `protobuf:"bytes,2,req,name=deviceType" json:"deviceType,omitempty"`
	Uid              *string `protobuf:"bytes,3,req,name=uid" json:"uid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetMech() string {
	if m != nil && m.Mech != nil {
		return *m.Mech
	}
	return ""
}

func (m *Request) GetDeviceType() string {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return ""
}

func (m *Request) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

type Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func init() {
}
