package server

import (
	"context"
	"expression-agent/internal/service"
	pb "expression-agent/proto"
)

type Server struct {
	pb.MiniTaskServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Solve(ctx context.Context, in *pb.MiniTaskRequest) (*pb.MiniTaskResponse, error) {
	calc := service.NewCalc(in.Task)
	res, err := calc.Calculate()

	return &pb.MiniTaskResponse{
		ExpressionId: in.ExpressionId,
		Task:         in.Task,
		Result:       int64(res),
		IsValid:      err == nil,
	}, nil
}
