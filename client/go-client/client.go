package main

import (
	"context"
	"fmt"

	"github.com/Ashilesh/chat-client/chat"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	res, err := c.SayHello(context.Background(), &chat.Empty{})

	if err != nil {
		panic(err)
	}

	fmt.Println("response from server", res)
}
