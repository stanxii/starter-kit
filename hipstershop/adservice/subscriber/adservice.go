package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	pb "github.com/micro-in-cn/starter-kit/hipstershop/pb"
)

type AdService struct{}

func (e *AdService) Handle(ctx context.Context, msg *pb.Empty) error {
	log.Log("Handler Received message: ")
	return nil
}

func Handler(ctx context.Context, msg *pb.Empty) error {
	log.Log("Function Received message: ")
	return nil
}
