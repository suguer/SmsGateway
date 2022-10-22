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
