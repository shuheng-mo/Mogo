// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

/*
Package go_micro_service_product is a generated protocol buffer package.

It is generated from these files:
	product.proto

It has these top-level messages:
	ProductInfo
	ResponseProduct
*/
package go_micro_service_product

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductInfo struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ProductName string `protobuf:"bytes,2,opt,name=product_name,json=productName" json:"product_name,omitempty"`
}

func (m *ProductInfo) Reset()                    { *m = ProductInfo{} }
func (m *ProductInfo) String() string            { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()               {}
func (*ProductInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductInfo) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

type ResponseProduct struct {
	ProductId int64 `protobuf:"varint,1,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *ResponseProduct) Reset()                    { *m = ResponseProduct{} }
func (m *ResponseProduct) String() string            { return proto.CompactTextString(m) }
func (*ResponseProduct) ProtoMessage()               {}
func (*ResponseProduct) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResponseProduct) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "go.micro.service.product.ProductInfo")
	proto.RegisterType((*ResponseProduct)(nil), "go.micro.service.product.ResponseProduct")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xcf, 0xd7, 0xcb, 0xcd,
	0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0xca, 0x2b, 0x39,
	0x70, 0x71, 0x07, 0x40, 0x98, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x8a, 0x5c, 0x3c, 0x50, 0x95, 0xf1,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc, 0x50, 0x31, 0xbf, 0xc4,
	0xdc, 0x54, 0x25, 0x03, 0x2e, 0xfe, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa8, 0x49,
	0x42, 0xb2, 0x5c, 0x5c, 0x30, 0x5d, 0x70, 0xd3, 0x38, 0xa1, 0x22, 0x9e, 0x29, 0x46, 0x99, 0x5c,
	0xec, 0x30, 0x95, 0x71, 0x5c, 0x5c, 0x8e, 0x29, 0x29, 0x30, 0x9e, 0xaa, 0x1e, 0x2e, 0x77, 0xea,
	0x21, 0x39, 0x52, 0x4a, 0x13, 0xb7, 0x32, 0x34, 0x97, 0x24, 0xb1, 0x81, 0xfd, 0x6f, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x5c, 0xfe, 0xd2, 0xcc, 0x10, 0x01, 0x00, 0x00,
}
