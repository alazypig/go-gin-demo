package main

import (
	"context"
	"fmt"
	"log"

	pb "example.com/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to server: ", err)
	}
	defer conn.Close()

	client := pb.NewUserInfoServiceClient(conn)

	req := new(pb.UserRequest)
	req.Name = "edward"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatal("Failed to get user info: ", err)
	}

	fmt.Println("User Info: ", resp)

	req2 := new(pb.UserRequest)
	req2.Name = "another one"
	resp2, err2 := client.GetUserInfo(context.Background(), req2)
	if err2 != nil {
		log.Fatal("Failed to get user info: ", err2)
	}

	fmt.Println("User Info: ", resp2)
}
