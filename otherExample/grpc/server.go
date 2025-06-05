package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/grpc/proto"
	"google.golang.org/grpc"
)

type UserInfoService struct{
	pb.UnimplementedUserInfoServiceServer
}

var u = UserInfoService{}

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name

	if name == "edward" {
		resp = &pb.UserResponse{
			Id:    123,
			Name:  name,
			Age:   48,
			Title: []string{"engineer", "programmer"},
		}
	}

	err = nil
	return
}

func main() {
	port := ":12345"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	fmt.Println("server is running on port: ", port)
	s := grpc.NewServer()

	pb.RegisterUserInfoServiceServer(s, &u)
	s.Serve(l)
}
