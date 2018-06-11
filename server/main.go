package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ryosan-470/go-grpc-study/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type aliveService struct{}

func (s *aliveService) GetStatus(ctx context.Context, in *pb.Empty) (*pb.AliveResponse, error) {
	return &pb.AliveResponse{Status: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAliveServiceServer(s, new(aliveService))
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
