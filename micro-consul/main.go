package main

import (
	"context"
	"log"

	hello "github.com/ilovelili/micro-consul/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	// k8s "github.com/micro/kubernetes/go/micro"
)

// Say say
type Say struct{}

// Hello test
func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	opts := registry.Option(func(opts *registry.Options) {
		// opts.Addrs = []string{"10.27.253.238:8500"}   35.234.41.252:8500
		opts.Addrs = []string{"35.234.41.252:8500"}
	})

	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.Registry(registry.NewRegistry(opts)),
	)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
