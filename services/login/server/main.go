package main

import (
	loginPb "etby/services/login"
	in "etby/services/login/server/internal"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
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
	loginPb.RegisterAuthenticationServer(s, &server{})
	log.Println("server listenning at ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve ", err)
	}
}
