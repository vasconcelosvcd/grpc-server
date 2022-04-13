package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/vasconcelosvcd/grpc-server/api/doubleNum"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement pb.UnimplementedDoubleNumServer.
type server struct {
	pb.UnimplementedDoubleNumServer
}

// DoubleNum implements doubleNum.
func (s *server) DoubleNum(ctx context.Context, in *pb.DoubleNumRequest) (*pb.DoubleNumResponse, error) {
	log.Printf("Received: %v", in.GetNum())
	return &pb.DoubleNumResponse{Doubled: in.Num * 2}, nil
}

// DoubleNum implements doubleNum.
func (s *server) TripleNum(ctx context.Context, in *pb.TripleNumRequest) (*pb.TripleNumResponse, error) {
	log.Printf("Received: %v", in.GetNum())
	return &pb.TripleNumResponse{Tripled: in.Num * 3}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDoubleNumServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
