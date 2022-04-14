package servers

import (
	"io"
	"log"
	"math/rand"

	pb "github.com/vasconcelosvcd/grpc-server/api/bistream"
)

type BistreamServer struct {
	pb.UnimplementedDoubleStreamServer
}

// DoubleStream implements biway stream of doubleNum.
func (s *BistreamServer) DoubleStream(stream pb.DoubleStream_DoubleStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// received
		log.Printf("received num: %d, ", in.GetNum())

		dsr := &pb.DoubleStreamResponse{}
		dsr.Num = in.GetNum()
		dsr.Id = int32(rand.Intn(999))
		dsr.Doubled = dsr.Num * 2

		if err := stream.Send(dsr); err != nil {
			return err
		}
	}
}
