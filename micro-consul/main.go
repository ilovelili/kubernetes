package main

import (
	"context"
	"log"

	hello "github.com/ilovelili/micro-k8s/proto"
	micro "github.com/micro/go-micro"
	// k8s "github.com/micro/kubernetes/go/micro"
)

// Say data model
type Say struct{}

// Hello test say hello
func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
	)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
