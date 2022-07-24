package chat

import (
	"context"
	"fmt"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, e *Empty) (*Response, error) {
	fmt.Println("server-sayhello called")
	return &Response{Msg: "Hi"}, nil
}

func (s *Server) mustEmbedUnimplementedChatServiceServer() {}
