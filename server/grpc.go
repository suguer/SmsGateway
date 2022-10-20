package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/suguer/SmsGateway/gateway"
	"github.com/suguer/SmsGateway/model"
	pb "github.com/suguer/SmsGateway/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	pb.UnimplementedSmsGatewayServer
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}
func (server *GrpcServer) Start(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSmsGatewayServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// CreateMessage implements helloworld.GreeterServer
func (server *GrpcServer) CreateMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	gateway, err := gateway.NewGatewayInterface(in.GetPlatform(), &model.Config{
		AppID:     in.GetConfig().GetAppID(),
		AppSecret: in.GetConfig().GetAppSecret(),
	})
	if err != nil {
		return &pb.MessageReply{}, err
	}
	phone := model.NewPhone(in.GetMobile())
	message := model.NewMessage(in.GetContent(), in.GetTemplateCode(), in.GetSignName(), in.GetParam())
	if len(in.GetSdkAppId()) > 0 {
		message.SetSdkAppId(in.GetSdkAppId())
	}
	response, err := gateway.SendMessage(phone, message)
	return &pb.MessageReply{
		Code:       response.Code.Val,
		Error:      err.Error(),
		RequestId:  response.RequestId,
		BizId:      response.BizId,
		RawContent: response.RawContent,
	}, nil
}

func (server *GrpcServer) CreateSmsTemplate(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateReply, error) {
	gateway, err := gateway.NewGatewayInterface(in.GetPlatform(), &model.Config{
		AppID:     in.GetConfig().GetAppID(),
		AppSecret: in.GetConfig().GetAppSecret(),
	})
	if err != nil {
		return &pb.TemplateReply{}, err
	}
	template := model.NewTemplate(in.GetName(), in.GetContent(), in.GetTemplateType(), in.GetRemark(), in.GetInternational())
	response, err := gateway.CreateSmsTemplate(template)
	Error := ""
	if err != nil {
		Error = err.Error()
	}
	return &pb.TemplateReply{
		Code:         response.Code.Val,
		Error:        Error,
		RequestId:    response.RequestId,
		TemplateCode: response.TemplateCode,
		RawContent:   response.RawContent,
	}, nil
}

func (server *GrpcServer) QuerySmsTemplate(ctx context.Context, in *pb.QuerySmsTemplateRequest) (*pb.CommonReply, error) {
	gateway, err := gateway.NewGatewayInterface(in.GetPlatform(), &model.Config{
		AppID:     in.GetConfig().GetAppID(),
		AppSecret: in.GetConfig().GetAppSecret(),
	})
	if err != nil {
		return &pb.CommonReply{}, err
	}
	response, err := gateway.QuerySmsTemplate(in.GetTemplateCode())
	return &pb.CommonReply{
		Code:       response.Code.Val,
		Error:      err.Error(),
		RequestId:  response.RequestId,
		RawContent: response.RawContent,
	}, nil
}
