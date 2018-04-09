// Code generated by protoc-gen-go.
// source: github.com/hailocab/h2/proto/discovery/proto/register/register.proto
// DO NOT EDIT!

/*
Package com_hailocab_kernel_discovery_register is a generated protocol buffer package.

It is generated from these files:
	github.com/hailocab/h2/proto/discovery/proto/register/register.proto

It has these top-level messages:
	Request
	MultiRequest
	Response
*/
package com_hailocab_kernel_discovery_register

import proto "github.com/hailocab/protobuf/proto"
import json "encoding/json"
import math "math"
import com_hailocab_kernel_discovery "github.com/hailocab/h2/proto/discovery/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Request struct {
	InstanceId       *string                                `protobuf:"bytes,1,req,name=instanceId" json:"instanceId,omitempty"`
	Hostname         *string                                `protobuf:"bytes,2,req,name=hostname" json:"hostname,omitempty"`
	AzName           *string                                `protobuf:"bytes,3,req,name=azName" json:"azName,omitempty"`
	Service          *com_hailocab_kernel_discovery.Service `protobuf:"bytes,4,req,name=service" json:"service,omitempty"`
	Endpoint         *string                                `protobuf:"bytes,5,req,name=endpoint" json:"endpoint,omitempty"`
	Mean             *int32                                 `protobuf:"varint,6,req,name=mean" json:"mean,omitempty"`
	Upper95          *int32                                 `protobuf:"varint,7,req,name=upper95" json:"upper95,omitempty"`
	Subscribe        *string                                `protobuf:"bytes,8,opt,name=subscribe" json:"subscribe,omitempty"`
	MachineClass     *string                                `protobuf:"bytes,9,opt,name=machineClass" json:"machineClass,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *Request) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Request) GetAzName() string {
	if m != nil && m.AzName != nil {
		return *m.AzName
	}
	return ""
}

func (m *Request) GetService() *com_hailocab_kernel_discovery.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Request) GetEndpoint() string {
	if m != nil && m.Endpoint != nil {
		return *m.Endpoint
	}
	return ""
}

func (m *Request) GetMean() int32 {
	if m != nil && m.Mean != nil {
		return *m.Mean
	}
	return 0
}

func (m *Request) GetUpper95() int32 {
	if m != nil && m.Upper95 != nil {
		return *m.Upper95
	}
	return 0
}

func (m *Request) GetSubscribe() string {
	if m != nil && m.Subscribe != nil {
		return *m.Subscribe
	}
	return ""
}

func (m *Request) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

type MultiRequest struct {
	InstanceId       *string                                `protobuf:"bytes,1,req,name=instanceId" json:"instanceId,omitempty"`
	Hostname         *string                                `protobuf:"bytes,2,req,name=hostname" json:"hostname,omitempty"`
	AzName           *string                                `protobuf:"bytes,3,req,name=azName" json:"azName,omitempty"`
	Service          *com_hailocab_kernel_discovery.Service `protobuf:"bytes,4,req,name=service" json:"service,omitempty"`
	Endpoints        []*MultiRequest_Endpoint               `protobuf:"bytes,5,rep,name=endpoints" json:"endpoints,omitempty"`
	MachineClass     *string                                `protobuf:"bytes,6,opt,name=machineClass" json:"machineClass,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *MultiRequest) Reset()         { *m = MultiRequest{} }
func (m *MultiRequest) String() string { return proto.CompactTextString(m) }
func (*MultiRequest) ProtoMessage()    {}

func (m *MultiRequest) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *MultiRequest) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *MultiRequest) GetAzName() string {
	if m != nil && m.AzName != nil {
		return *m.AzName
	}
	return ""
}

func (m *MultiRequest) GetService() *com_hailocab_kernel_discovery.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *MultiRequest) GetEndpoints() []*MultiRequest_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *MultiRequest) GetMachineClass() string {
	if m != nil && m.MachineClass != nil {
		return *m.MachineClass
	}
	return ""
}

type MultiRequest_Endpoint struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Mean             *int32  `protobuf:"varint,2,req,name=mean" json:"mean,omitempty"`
	Upper95          *int32  `protobuf:"varint,3,req,name=upper95" json:"upper95,omitempty"`
	Subscribe        *string `protobuf:"bytes,4,opt,name=subscribe" json:"subscribe,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MultiRequest_Endpoint) Reset()         { *m = MultiRequest_Endpoint{} }
func (m *MultiRequest_Endpoint) String() string { return proto.CompactTextString(m) }
func (*MultiRequest_Endpoint) ProtoMessage()    {}

func (m *MultiRequest_Endpoint) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MultiRequest_Endpoint) GetMean() int32 {
	if m != nil && m.Mean != nil {
		return *m.Mean
	}
	return 0
}

func (m *MultiRequest_Endpoint) GetUpper95() int32 {
	if m != nil && m.Upper95 != nil {
		return *m.Upper95
	}
	return 0
}

func (m *MultiRequest_Endpoint) GetSubscribe() string {
	if m != nil && m.Subscribe != nil {
		return *m.Subscribe
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
