// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order/order.proto

/*
Package go_micro_service_order is a generated protocol buffer package.

It is generated from these files:
	proto/order/order.proto

It has these top-level messages:
	AllOrderRequest
	AllOrder
	OrderID
	OrderInfo
	OrderDetail
	Response
	PayStatus
	ShipStatus
*/
package go_micro_service_order

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

type AllOrderRequest struct {
}

func (m *AllOrderRequest) Reset()                    { *m = AllOrderRequest{} }
func (m *AllOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*AllOrderRequest) ProtoMessage()               {}
func (*AllOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllOrder struct {
	OrderInfo []*OrderInfo `protobuf:"bytes,1,rep,name=order_info,json=orderInfo" json:"order_info,omitempty"`
}

func (m *AllOrder) Reset()                    { *m = AllOrder{} }
func (m *AllOrder) String() string            { return proto.CompactTextString(m) }
func (*AllOrder) ProtoMessage()               {}
func (*AllOrder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllOrder) GetOrderInfo() []*OrderInfo {
	if m != nil {
		return m.OrderInfo
	}
	return nil
}

type OrderID struct {
	OrderId int64 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *OrderID) Reset()                    { *m = OrderID{} }
func (m *OrderID) String() string            { return proto.CompactTextString(m) }
func (*OrderID) ProtoMessage()               {}
func (*OrderID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OrderID) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type OrderInfo struct {
	Id          int64          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	PayStatus   int32          `protobuf:"varint,2,opt,name=pay_status,json=payStatus" json:"pay_status,omitempty"`
	ShipStatus  int32          `protobuf:"varint,3,opt,name=ship_status,json=shipStatus" json:"ship_status,omitempty"`
	Price       float64        `protobuf:"fixed64,4,opt,name=price" json:"price,omitempty"`
	OrderDetail []*OrderDetail `protobuf:"bytes,5,rep,name=order_detail,json=orderDetail" json:"order_detail,omitempty"`
}

func (m *OrderInfo) Reset()                    { *m = OrderInfo{} }
func (m *OrderInfo) String() string            { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()               {}
func (*OrderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OrderInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderInfo) GetPayStatus() int32 {
	if m != nil {
		return m.PayStatus
	}
	return 0
}

func (m *OrderInfo) GetShipStatus() int32 {
	if m != nil {
		return m.ShipStatus
	}
	return 0
}

func (m *OrderInfo) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderInfo) GetOrderDetail() []*OrderDetail {
	if m != nil {
		return m.OrderDetail
	}
	return nil
}

type OrderDetail struct {
	Id           int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ProductId    int64 `protobuf:"varint,2,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	ProductNum   int64 `protobuf:"varint,3,opt,name=product_num,json=productNum" json:"product_num,omitempty"`
	ProductSize  int64 `protobuf:"varint,4,opt,name=product_size,json=productSize" json:"product_size,omitempty"`
	ProductPrice int64 `protobuf:"varint,5,opt,name=product_price,json=productPrice" json:"product_price,omitempty"`
	OrderId      int64 `protobuf:"varint,6,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
}

func (m *OrderDetail) Reset()                    { *m = OrderDetail{} }
func (m *OrderDetail) String() string            { return proto.CompactTextString(m) }
func (*OrderDetail) ProtoMessage()               {}
func (*OrderDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OrderDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OrderDetail) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *OrderDetail) GetProductNum() int64 {
	if m != nil {
		return m.ProductNum
	}
	return 0
}

func (m *OrderDetail) GetProductSize() int64 {
	if m != nil {
		return m.ProductSize
	}
	return 0
}

func (m *OrderDetail) GetProductPrice() int64 {
	if m != nil {
		return m.ProductPrice
	}
	return 0
}

func (m *OrderDetail) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PayStatus struct {
	OrderId   int64 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	PayStatus int32 `protobuf:"varint,2,opt,name=pay_status,json=payStatus" json:"pay_status,omitempty"`
}

func (m *PayStatus) Reset()                    { *m = PayStatus{} }
func (m *PayStatus) String() string            { return proto.CompactTextString(m) }
func (*PayStatus) ProtoMessage()               {}
func (*PayStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PayStatus) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *PayStatus) GetPayStatus() int32 {
	if m != nil {
		return m.PayStatus
	}
	return 0
}

type ShipStatus struct {
	OrderId    int64 `protobuf:"varint,1,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	ShipStatus int32 `protobuf:"varint,2,opt,name=ship_status,json=shipStatus" json:"ship_status,omitempty"`
}

func (m *ShipStatus) Reset()                    { *m = ShipStatus{} }
func (m *ShipStatus) String() string            { return proto.CompactTextString(m) }
func (*ShipStatus) ProtoMessage()               {}
func (*ShipStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ShipStatus) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *ShipStatus) GetShipStatus() int32 {
	if m != nil {
		return m.ShipStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*AllOrderRequest)(nil), "go.micro.service.order.AllOrderRequest")
	proto.RegisterType((*AllOrder)(nil), "go.micro.service.order.AllOrder")
	proto.RegisterType((*OrderID)(nil), "go.micro.service.order.OrderID")
	proto.RegisterType((*OrderInfo)(nil), "go.micro.service.order.OrderInfo")
	proto.RegisterType((*OrderDetail)(nil), "go.micro.service.order.OrderDetail")
	proto.RegisterType((*Response)(nil), "go.micro.service.order.Response")
	proto.RegisterType((*PayStatus)(nil), "go.micro.service.order.PayStatus")
	proto.RegisterType((*ShipStatus)(nil), "go.micro.service.order.ShipStatus")
}

func init() { proto.RegisterFile("proto/order/order.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0xed, 0xba, 0x89, 0xc7, 0x81, 0x8a, 0x55, 0x01, 0x13, 0x09, 0xd5, 0xdd, 0x56, 0x22,
	0x27, 0x23, 0x95, 0x3f, 0xc0, 0x87, 0xa1, 0x44, 0x42, 0x50, 0x6d, 0x0a, 0x07, 0x24, 0x14, 0x99,
	0x78, 0x5a, 0x2c, 0x25, 0x59, 0xe3, 0xb5, 0x91, 0xd2, 0x7f, 0xc6, 0x8d, 0x1f, 0xc6, 0x01, 0x65,
	0x37, 0xeb, 0x98, 0x28, 0x8e, 0x7d, 0x89, 0x76, 0x67, 0xde, 0x3c, 0xbf, 0x37, 0x33, 0x1b, 0x78,
	0x9c, 0x66, 0x3c, 0xe7, 0xcf, 0x79, 0x16, 0x63, 0xa6, 0x7e, 0x03, 0x19, 0x21, 0x8f, 0x6e, 0x79,
	0x30, 0x4f, 0xa6, 0x19, 0x0f, 0x04, 0x66, 0xbf, 0x92, 0x29, 0x06, 0x32, 0x4b, 0x1f, 0xc0, 0xd1,
	0xab, 0xd9, 0xec, 0xd3, 0xea, 0xcc, 0xf0, 0x67, 0x81, 0x22, 0xa7, 0x1f, 0xa0, 0xa7, 0x43, 0xe4,
	0x25, 0x80, 0xc4, 0x4d, 0x92, 0xc5, 0x0d, 0xf7, 0x0c, 0xdf, 0x1a, 0xba, 0x17, 0xa7, 0xc1, 0x6e,
	0xae, 0x40, 0x96, 0x8c, 0x16, 0x37, 0x9c, 0x39, 0x5c, 0x1f, 0xe9, 0x39, 0x74, 0x55, 0x3c, 0x24,
	0x4f, 0xa0, 0xb7, 0x26, 0x8b, 0x3d, 0xc3, 0x37, 0x86, 0x16, 0xeb, 0x2a, 0x5c, 0x4c, 0x7f, 0x1b,
	0xe0, 0x94, 0xe5, 0xe4, 0x3e, 0x98, 0x25, 0xc4, 0x4c, 0x62, 0xf2, 0x14, 0x20, 0x8d, 0x96, 0x13,
	0x91, 0x47, 0x79, 0x21, 0x3c, 0xd3, 0x37, 0x86, 0x36, 0x73, 0xd2, 0x68, 0x39, 0x96, 0x01, 0x72,
	0x02, 0xae, 0xf8, 0x91, 0xa4, 0x3a, 0x6f, 0xc9, 0x3c, 0xac, 0x42, 0x6b, 0xc0, 0x31, 0xd8, 0x69,
	0x96, 0x4c, 0xd1, 0x3b, 0xf0, 0x8d, 0xa1, 0xc1, 0xd4, 0x85, 0xbc, 0x83, 0xbe, 0x92, 0x13, 0x63,
	0x1e, 0x25, 0x33, 0xcf, 0x96, 0xee, 0xce, 0xf6, 0xba, 0x0b, 0x25, 0x94, 0xb9, 0x7c, 0x73, 0xa1,
	0x7f, 0x0c, 0x70, 0x2b, 0xc9, 0x9d, 0xea, 0x33, 0x1e, 0x17, 0xd3, 0x7c, 0x65, 0xdc, 0x94, 0x71,
	0x67, 0x1d, 0x19, 0xc5, 0x2b, 0xf5, 0x3a, 0xbd, 0x28, 0xe6, 0x52, 0xbd, 0xc5, 0x74, 0xc5, 0xc7,
	0x62, 0x4e, 0x4e, 0xa1, 0xaf, 0x01, 0x22, 0xb9, 0x53, 0x26, 0x2c, 0xa6, 0x8b, 0xc6, 0xc9, 0x1d,
	0x92, 0x33, 0xb8, 0xa7, 0x21, 0xca, 0xa8, 0x2d, 0x31, 0xba, 0xee, 0x4a, 0xfa, 0xad, 0xb6, 0xff,
	0xf0, 0xff, 0xf6, 0x9f, 0x43, 0x8f, 0xa1, 0x48, 0xf9, 0x42, 0x20, 0xf1, 0xa0, 0x3b, 0x47, 0x21,
	0xa2, 0x5b, 0x94, 0x1e, 0x1c, 0xa6, 0xaf, 0xf4, 0x2d, 0x38, 0x57, 0x65, 0xd3, 0xeb, 0x87, 0xd9,
	0x30, 0x2e, 0xfa, 0x1e, 0x60, 0xbc, 0x99, 0xcd, 0x1e, 0x9e, 0xad, 0xb9, 0x9a, 0xdb, 0x73, 0xbd,
	0xf8, 0x7b, 0x00, 0xb6, 0xda, 0xd3, 0x6b, 0xe8, 0x5f, 0x62, 0x2e, 0xcf, 0xaf, 0x97, 0xa3, 0x90,
	0x9c, 0xec, 0xdf, 0xd1, 0x70, 0xd0, 0xbc, 0xc4, 0xb4, 0x43, 0xbe, 0x82, 0x7b, 0x89, 0x79, 0xf9,
	0x18, 0x9e, 0xd5, 0xd5, 0x6c, 0xbd, 0xa0, 0x81, 0xdf, 0x04, 0xa4, 0x1d, 0x32, 0x06, 0xf7, 0x4d,
	0x86, 0x51, 0x8e, 0x8a, 0xbb, 0x59, 0xcf, 0xa0, 0xc9, 0x13, 0xed, 0x90, 0x2f, 0x70, 0x14, 0xe2,
	0x0c, 0xd7, 0xa4, 0xed, 0x3a, 0x51, 0x2b, 0x56, 0x6f, 0x04, 0xed, 0x90, 0x6f, 0x70, 0xfc, 0x39,
	0x8d, 0xb5, 0xd8, 0xcd, 0x12, 0xd4, 0xaa, 0x2e, 0x21, 0xad, 0xe8, 0x27, 0xf0, 0xb0, 0x42, 0x5f,
	0x59, 0x0e, 0x5a, 0x57, 0xbc, 0xc1, 0xb4, 0xfa, 0xc0, 0x35, 0xb8, 0x95, 0x0f, 0xb4, 0x69, 0x76,
	0x0b, 0xd6, 0xef, 0x87, 0xf2, 0xaf, 0xf5, 0xc5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x36,
	0xaf, 0xcc, 0x75, 0x05, 0x00, 0x00,
}
