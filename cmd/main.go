package main

import (
	server "expression-agent/internal"
	pb "expression-agent/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	pb.MiniTaskServiceServer
}

func main() {
	host := "0.0.0.0"
	port := "2552"

	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Println("could not start server")
		os.Exit(1)
	}

	log.Println("server started at port: ", port)

	grpcServe := grpc.NewServer()
	taskServer := server.NewServer()

	pb.RegisterMiniTaskServiceServer(grpcServe, taskServer)

	if err := grpcServe.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}
}
