// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calculatorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	PrimeNumber(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/caculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeNumber(ctx context.Context, in *PrimeNumberRequest, opts ...grpc.CallOption) (CalculatorService_PrimeNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/caculator.CalculatorService/PrimeNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeNumberClient interface {
	Recv() (*PrimeNumberResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeNumberClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeNumberClient) Recv() (*PrimeNumberResponse, error) {
	m := new(PrimeNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations should embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	PrimeNumber(*PrimeNumberRequest, CalculatorService_PrimeNumberServer) error
}

// UnimplementedCalculatorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) PrimeNumber(*PrimeNumberRequest, CalculatorService_PrimeNumberServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeNumber not implemented")
}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/caculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeNumber(m, &calculatorServicePrimeNumberServer{stream})
}

type CalculatorService_PrimeNumberServer interface {
	Send(*PrimeNumberResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeNumberServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeNumberServer) Send(m *PrimeNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "caculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumber",
			Handler:       _CalculatorService_PrimeNumber_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
