package main

import "github.com/suguer/SmsGateway/server"

func main() {
	s := server.NewGrpcServer()
	s.Start(50051)
}
