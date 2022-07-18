package main

import (
	"struct_demo/middle"
	"struct_demo/server"
)

func main() {

	s := server.NewHttpServer("test_server")
	s.Route("/getheader", middle.SayHello)

	s.Start(":9090")

}
