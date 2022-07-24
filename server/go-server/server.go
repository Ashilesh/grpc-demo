package main

import (
	"net"

	"github.com/Ashilesh/grpc-demo-server/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	// register services
	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
