package main

import (
	"log"

	hello "github.com/ilovelili/micro-k8s/proto"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"context"
)

// Say say
type Say struct{}

// Hello hello
func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := k8s.NewService(
		micro.Name("greeter"),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
