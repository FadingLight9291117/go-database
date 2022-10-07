package main

import (
	pb "com.fadinglight/db/api/sql"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	port := 50051
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewSqlClient(conn)

	ctx, cancelFun := context.WithTimeout(context.Background(), time.Second)
	defer cancelFun()
	resp, err := client.Select(ctx, &pb.SelectRequest{Sql: "Test"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Results)
}
