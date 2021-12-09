package main

import (
	regPb "etby/services/register"
	in "etby/services/register/server/internal"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	port = ":50051"
)

type server = in.Server

func main() {
	in.CheckConn()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle: 5 * time.Minute,
			}),
	)
	regPb.RegisterRegisterServer(s, &server{})
	log.Println("server listening at ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve ", err)
	}
}
