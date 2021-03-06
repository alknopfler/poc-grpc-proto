// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template.proto

package template

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GenerateTemplateRequest struct {
	TemplateId string                              `protobuf:"bytes,1,opt,name=templateId" json:"templateId,omitempty"`
	Chunk      *GenerateTemplateRequest_Chunk      `protobuf:"bytes,3,opt,name=chunk" json:"chunk,omitempty"`
	Params     []*GenerateTemplateRequest_KeyValue `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
}

func (m *GenerateTemplateRequest) Reset()                    { *m = GenerateTemplateRequest{} }
func (m *GenerateTemplateRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateTemplateRequest) ProtoMessage()               {}
func (*GenerateTemplateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GenerateTemplateRequest) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *GenerateTemplateRequest) GetChunk() *GenerateTemplateRequest_Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *GenerateTemplateRequest) GetParams() []*GenerateTemplateRequest_KeyValue {
	if m != nil {
		return m.Params
	}
	return nil
}

type GenerateTemplateRequest_Chunk struct {
	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Position int64  `protobuf:"varint,2,opt,name=position" json:"position,omitempty"`
}

func (m *GenerateTemplateRequest_Chunk) Reset()         { *m = GenerateTemplateRequest_Chunk{} }
func (m *GenerateTemplateRequest_Chunk) String() string { return proto.CompactTextString(m) }
func (*GenerateTemplateRequest_Chunk) ProtoMessage()    {}
func (*GenerateTemplateRequest_Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0}
}

func (m *GenerateTemplateRequest_Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GenerateTemplateRequest_Chunk) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type GenerateTemplateRequest_KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *GenerateTemplateRequest_KeyValue) Reset()         { *m = GenerateTemplateRequest_KeyValue{} }
func (m *GenerateTemplateRequest_KeyValue) String() string { return proto.CompactTextString(m) }
func (*GenerateTemplateRequest_KeyValue) ProtoMessage()    {}
func (*GenerateTemplateRequest_KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 1}
}

func (m *GenerateTemplateRequest_KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GenerateTemplateRequest_KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GenerateTemplateResponse struct {
	// Types that are valid to be assigned to Content:
	//	*GenerateTemplateResponse_Error
	//	*GenerateTemplateResponse_Result
	Content isGenerateTemplateResponse_Content `protobuf_oneof:"content"`
}

func (m *GenerateTemplateResponse) Reset()                    { *m = GenerateTemplateResponse{} }
func (m *GenerateTemplateResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateTemplateResponse) ProtoMessage()               {}
func (*GenerateTemplateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isGenerateTemplateResponse_Content interface {
	isGenerateTemplateResponse_Content()
}

type GenerateTemplateResponse_Error struct {
	Error *Error `protobuf:"bytes,1,opt,name=error,oneof"`
}
type GenerateTemplateResponse_Result struct {
	Result string `protobuf:"bytes,2,opt,name=result,oneof"`
}

func (*GenerateTemplateResponse_Error) isGenerateTemplateResponse_Content()  {}
func (*GenerateTemplateResponse_Result) isGenerateTemplateResponse_Content() {}

func (m *GenerateTemplateResponse) GetContent() isGenerateTemplateResponse_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GenerateTemplateResponse) GetError() *Error {
	if x, ok := m.GetContent().(*GenerateTemplateResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (m *GenerateTemplateResponse) GetResult() string {
	if x, ok := m.GetContent().(*GenerateTemplateResponse_Result); ok {
		return x.Result
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GenerateTemplateResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GenerateTemplateResponse_OneofMarshaler, _GenerateTemplateResponse_OneofUnmarshaler, _GenerateTemplateResponse_OneofSizer, []interface{}{
		(*GenerateTemplateResponse_Error)(nil),
		(*GenerateTemplateResponse_Result)(nil),
	}
}

func _GenerateTemplateResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GenerateTemplateResponse)
	// content
	switch x := m.Content.(type) {
	case *GenerateTemplateResponse_Error:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *GenerateTemplateResponse_Result:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Result)
	case nil:
	default:
		return fmt.Errorf("GenerateTemplateResponse.Content has unexpected type %T", x)
	}
	return nil
}

func _GenerateTemplateResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GenerateTemplateResponse)
	switch tag {
	case 1: // content.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Error)
		err := b.DecodeMessage(msg)
		m.Content = &GenerateTemplateResponse_Error{msg}
		return true, err
	case 2: // content.result
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Content = &GenerateTemplateResponse_Result{x}
		return true, err
	default:
		return false, nil
	}
}

func _GenerateTemplateResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GenerateTemplateResponse)
	// content
	switch x := m.Content.(type) {
	case *GenerateTemplateResponse_Error:
		s := proto.Size(x.Error)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GenerateTemplateResponse_Result:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Result)))
		n += len(x.Result)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*GenerateTemplateRequest)(nil), "template.GenerateTemplateRequest")
	proto.RegisterType((*GenerateTemplateRequest_Chunk)(nil), "template.GenerateTemplateRequest.Chunk")
	proto.RegisterType((*GenerateTemplateRequest_KeyValue)(nil), "template.GenerateTemplateRequest.KeyValue")
	proto.RegisterType((*GenerateTemplateResponse)(nil), "template.GenerateTemplateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Template service

type TemplateClient interface {
	GenerateTemplate(ctx context.Context, in *GenerateTemplateRequest, opts ...grpc.CallOption) (*GenerateTemplateResponse, error)
}

type templateClient struct {
	cc *grpc.ClientConn
}

func NewTemplateClient(cc *grpc.ClientConn) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) GenerateTemplate(ctx context.Context, in *GenerateTemplateRequest, opts ...grpc.CallOption) (*GenerateTemplateResponse, error) {
	out := new(GenerateTemplateResponse)
	err := grpc.Invoke(ctx, "/template.Template/GenerateTemplate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Template service

type TemplateServer interface {
	GenerateTemplate(context.Context, *GenerateTemplateRequest) (*GenerateTemplateResponse, error)
}

func RegisterTemplateServer(s *grpc.Server, srv TemplateServer) {
	s.RegisterService(&_Template_serviceDesc, srv)
}

func _Template_GenerateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GenerateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.Template/GenerateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GenerateTemplate(ctx, req.(*GenerateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Template_serviceDesc = grpc.ServiceDesc{
	ServiceName: "template.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTemplate",
			Handler:    _Template_GenerateTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template.proto",
}

func init() { proto.RegisterFile("template.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x6e, 0x1a, 0x53, 0xd3, 0x89, 0x68, 0x19, 0x04, 0x97, 0x1c, 0x24, 0xf6, 0xd2, 0xe0, 0xa1,
	0x87, 0x78, 0xf0, 0xe4, 0xa5, 0x22, 0x56, 0xbc, 0x2d, 0x22, 0x78, 0x5c, 0xdb, 0x81, 0x96, 0xb6,
	0xbb, 0x71, 0x77, 0x23, 0xf4, 0x91, 0x7c, 0x4b, 0xc9, 0x6e, 0x52, 0x45, 0xd1, 0xde, 0xf6, 0x9b,
	0x99, 0xef, 0x67, 0x86, 0x85, 0x63, 0x4b, 0x9b, 0x72, 0x2d, 0x2c, 0x8d, 0x4b, 0xad, 0xac, 0xc2,
	0xb8, 0xc5, 0x69, 0x42, 0x5a, 0x2b, 0xed, 0xcb, 0xc3, 0x8f, 0x2e, 0x9c, 0xdd, 0x93, 0x24, 0x2d,
	0x2c, 0x3d, 0x35, 0x13, 0x9c, 0xde, 0x2a, 0x32, 0x16, 0xcf, 0x01, 0x5a, 0xd2, 0xc3, 0x9c, 0x05,
	0x59, 0x90, 0xf7, 0xf9, 0xb7, 0x0a, 0xde, 0x40, 0x34, 0x5b, 0x54, 0x72, 0xc5, 0xc2, 0x2c, 0xc8,
	0x93, 0x62, 0x34, 0xde, 0x59, 0xfe, 0xa1, 0x38, 0xbe, 0xad, 0xc7, 0xb9, 0x67, 0xe1, 0x04, 0x7a,
	0xa5, 0xd0, 0x62, 0x63, 0x58, 0x37, 0x0b, 0xf3, 0xa4, 0xb8, 0xdc, 0xcf, 0x7f, 0xa4, 0xed, 0xb3,
	0x58, 0x57, 0xc4, 0x1b, 0x66, 0x7a, 0x0d, 0x91, 0xd3, 0x44, 0x84, 0x83, 0xb9, 0xb0, 0xc2, 0xa5,
	0x3c, 0xe2, 0xee, 0x8d, 0x29, 0xc4, 0xa5, 0x32, 0x4b, 0xbb, 0x54, 0x92, 0x75, 0xb3, 0x20, 0x0f,
	0xf9, 0x0e, 0xa7, 0x05, 0xc4, 0xad, 0x18, 0x0e, 0x20, 0x5c, 0xd1, 0xb6, 0x59, 0xb0, 0x7e, 0xe2,
	0x29, 0x44, 0xef, 0x75, 0xcb, 0xd1, 0xfa, 0xdc, 0x83, 0xe1, 0x02, 0xd8, 0xef, 0x60, 0xa6, 0x54,
	0xd2, 0x10, 0x8e, 0x20, 0x72, 0x67, 0x75, 0x2a, 0x49, 0x71, 0xf2, 0xb5, 0xcb, 0x5d, 0x5d, 0x9e,
	0x76, 0xb8, 0xef, 0x23, 0x83, 0x9e, 0x26, 0x53, 0xad, 0xad, 0xd7, 0x9e, 0x76, 0x78, 0x83, 0x27,
	0x7d, 0x38, 0x9c, 0x29, 0x69, 0x49, 0xda, 0x82, 0x20, 0x6e, 0x1d, 0xf0, 0x05, 0x06, 0x3f, 0x5d,
	0xf1, 0x62, 0xef, 0xa9, 0xd2, 0xe1, 0x7f, 0x23, 0x3e, 0xf4, 0x6b, 0xcf, 0xfd, 0x81, 0xab, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0xb9, 0x32, 0x4e, 0x2c, 0x02, 0x00, 0x00,
}
