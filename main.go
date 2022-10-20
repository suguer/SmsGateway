package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/suguer/SmsGateway/gateway"
	"github.com/suguer/SmsGateway/model"
	pb "github.com/suguer/SmsGateway/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedSmsGatewayServer
}

// CreateMessage implements helloworld.GreeterServer
func (s *server) CreateMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
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

func (s *server) CreateSmsTemplate(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateReply, error) {
	gateway, err := gateway.NewGatewayInterface(in.GetPlatform(), &model.Config{
		AppID:     in.GetConfig().GetAppID(),
		AppSecret: in.GetConfig().GetAppSecret(),
	})
	if err != nil {
		return &pb.TemplateReply{}, err
	}
	template := model.NewTemplate(in.GetName(), in.GetContent(), in.GetTemplateType(), in.GetRemark(), in.GetInternational())
	fmt.Printf("template: %v\n", template)
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

func (s *server) QuerySmsTemplate(ctx context.Context, in *pb.QuerySmsTemplateRequest) (*pb.CommonReply, error) {
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

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSmsGatewayServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
