package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	pb "github.com/micro-community/x-email/pb"
)

type EmailService struct{}

func (e *EmailService) Handle(ctx context.Context, msg *pb.Empty) error {
	log.Log("Handler Received message: ")
	return nil
}

func Handler(ctx context.Context, msg *pb.Empty) error {
	log.Log("Function Received message: ")
	return nil
}
