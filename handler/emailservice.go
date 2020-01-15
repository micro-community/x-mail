package handler

import (
	"context"

	pb "github.com/micro-community/x-email/pb"

)

type EmailService struct{}

func (e *EmailService) SendOrderConfirmation(context.Context, *pb.SendOrderConfirmationRequest, *pb.Empty) error {
	return nil
}
