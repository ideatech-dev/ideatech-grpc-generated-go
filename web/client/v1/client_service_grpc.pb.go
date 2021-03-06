// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.2
// source: web/client/v1/client_service.proto

package grpcweb_client_v1

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

// SurveyClientServiceClient is the client API for SurveyClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SurveyClientServiceClient interface {
	StartSurvey(ctx context.Context, in *StartSurveyRequest, opts ...grpc.CallOption) (*StartSurveyResponse, error)
	GetQuestion(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionResponse, error)
	SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error)
}

type surveyClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSurveyClientServiceClient(cc grpc.ClientConnInterface) SurveyClientServiceClient {
	return &surveyClientServiceClient{cc}
}

func (c *surveyClientServiceClient) StartSurvey(ctx context.Context, in *StartSurveyRequest, opts ...grpc.CallOption) (*StartSurveyResponse, error) {
	out := new(StartSurveyResponse)
	err := c.cc.Invoke(ctx, "/web.client.v1.SurveyClientService/StartSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClientServiceClient) GetQuestion(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionResponse, error) {
	out := new(GetQuestionResponse)
	err := c.cc.Invoke(ctx, "/web.client.v1.SurveyClientService/GetQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *surveyClientServiceClient) SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error) {
	out := new(SubmitAnswerResponse)
	err := c.cc.Invoke(ctx, "/web.client.v1.SurveyClientService/SubmitAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SurveyClientServiceServer is the server API for SurveyClientService service.
// All implementations must embed UnimplementedSurveyClientServiceServer
// for forward compatibility
type SurveyClientServiceServer interface {
	StartSurvey(context.Context, *StartSurveyRequest) (*StartSurveyResponse, error)
	GetQuestion(context.Context, *GetQuestionRequest) (*GetQuestionResponse, error)
	SubmitAnswer(context.Context, *SubmitAnswerRequest) (*SubmitAnswerResponse, error)
	mustEmbedUnimplementedSurveyClientServiceServer()
}

// UnimplementedSurveyClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSurveyClientServiceServer struct {
}

func (UnimplementedSurveyClientServiceServer) StartSurvey(context.Context, *StartSurveyRequest) (*StartSurveyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSurvey not implemented")
}
func (UnimplementedSurveyClientServiceServer) GetQuestion(context.Context, *GetQuestionRequest) (*GetQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (UnimplementedSurveyClientServiceServer) SubmitAnswer(context.Context, *SubmitAnswerRequest) (*SubmitAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswer not implemented")
}
func (UnimplementedSurveyClientServiceServer) mustEmbedUnimplementedSurveyClientServiceServer() {}

// UnsafeSurveyClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SurveyClientServiceServer will
// result in compilation errors.
type UnsafeSurveyClientServiceServer interface {
	mustEmbedUnimplementedSurveyClientServiceServer()
}

func RegisterSurveyClientServiceServer(s grpc.ServiceRegistrar, srv SurveyClientServiceServer) {
	s.RegisterService(&SurveyClientService_ServiceDesc, srv)
}

func _SurveyClientService_StartSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSurveyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyClientServiceServer).StartSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.client.v1.SurveyClientService/StartSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyClientServiceServer).StartSurvey(ctx, req.(*StartSurveyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SurveyClientService_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyClientServiceServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.client.v1.SurveyClientService/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyClientServiceServer).GetQuestion(ctx, req.(*GetQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SurveyClientService_SubmitAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SurveyClientServiceServer).SubmitAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.client.v1.SurveyClientService/SubmitAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SurveyClientServiceServer).SubmitAnswer(ctx, req.(*SubmitAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SurveyClientService_ServiceDesc is the grpc.ServiceDesc for SurveyClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SurveyClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "web.client.v1.SurveyClientService",
	HandlerType: (*SurveyClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartSurvey",
			Handler:    _SurveyClientService_StartSurvey_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _SurveyClientService_GetQuestion_Handler,
		},
		{
			MethodName: "SubmitAnswer",
			Handler:    _SurveyClientService_SubmitAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web/client/v1/client_service.proto",
}
