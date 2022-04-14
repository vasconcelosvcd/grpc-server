package servers

import (
	"context"
	"log"

	pb "github.com/vasconcelosvcd/grpc-server/api/doubleNum"
)

// server is used to implement pb.UnimplementedDoubleNumServer.
type DoubleNumServer struct {
	pb.UnimplementedDoubleNumServer
}

// DoubleNum implements doubleNum.
func (s *DoubleNumServer) DoubleNum(ctx context.Context, in *pb.DoubleNumRequest) (*pb.DoubleNumResponse, error) {
	log.Printf("Received: %v", in.GetNum())
	return &pb.DoubleNumResponse{Doubled: in.Num * 2}, nil
}
