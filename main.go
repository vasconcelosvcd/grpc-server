package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb2 "github.com/vasconcelosvcd/grpc-server/api/bistream"
	pb "github.com/vasconcelosvcd/grpc-server/api/doubleNum"
	"github.com/vasconcelosvcd/grpc-server/pkg/servers"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDoubleNumServer(s, &servers.DoubleNumServer{})
	pb2.RegisterDoubleStreamServer(s, &servers.BistreamServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
