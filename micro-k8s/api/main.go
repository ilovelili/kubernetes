package main

import (
	"log"

	hello "github.com/ilovelili/micro-k8s/api/proto"
	"github.com/micro/go-micro/client"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	"github.com/emicklei/go-restful"
	"github.com/micro/go-web"
	"golang.org/x/net/context"
)

type Greeter struct{}

var (
	cl hello.SayService
)

func (g *Greeter) Greet(req *restful.Request, rsp *restful.Response) {
	log.Print("Received API request at /greet")
	name := req.PathParameter("name")
	response, err := cl.Hello(context.TODO(), &hello.Request{Name: name})
	if err != nil {
		rsp.WriteEntity(err)
	} else {
		rsp.WriteEntity(response)
	}
}

func main() {
	service := web.NewService(
		web.Name("client.greeter"),
	)

	service.Init()

	// DevSvc client
	cl = hello.NewSayService("server.greeter", client.DefaultClient)

	// handle restful
	greeter := new(Greeter)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path("/greeter")
	ws.Route(ws.GET("/{name}").To(greeter.Greet))
	wc.Add(ws)

	// register handler
	service.Handle("/", wc)

	// run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
