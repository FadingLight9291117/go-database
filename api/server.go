package main

import (
	pb "com.fadinglight/db/api/sql"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SqlServer struct {
	pb.UnimplementedSqlServer
}

func (server *SqlServer) Select(ctx context.Context, request *pb.SelectRequest) (*pb.SelectResponse, error) {
	fmt.Println(request)
	return &pb.SelectResponse{
		Results: "Hello, world",
	}, nil
}

func ServerStart() {
	port := 50051
	server := grpc.NewServer()
	pb.RegisterSqlServer(server, &SqlServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	ServerStart()
}
