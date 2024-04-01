package main

import (
	server "expression-agent/internal"
	pb "expression-agent/proto"
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	pb.MiniTaskServiceServer
}

func init() {
	viper.SetConfigName(".env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {

	host := viper.GetString("server.host")
	port := viper.GetString("server.port")

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
