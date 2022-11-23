package main

import (
	"GMGgRPCServer/pkg/adder"
	"GMGgRPCServer/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		log.Fatal(err)
	}

	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	api.RegisterScoreServer(s, &adder.GRPCServer{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}
