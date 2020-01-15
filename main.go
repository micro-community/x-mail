package main

import (
	"github.com/micro-community/x-email/handler"
	"github.com/micro-community/x-email/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	emailservice "github.com/micro-community/x-email/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hipstershop.email"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	emailservice.RegisterEmailServiceHandler(service.Server(), new(handler.EmailService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.email", service.Server(), new(subscriber.EmailService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.email", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
