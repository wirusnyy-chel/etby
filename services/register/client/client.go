package client

import (
	"context"
	rebPB "etby/services/register"
	"log"

	"google.golang.org/grpc"
)

const address = "localhost:50052"

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

type regClient struct{}

func GetRegClient() regClient {
	return regClient{}
}
func (cli *regClient) Register(ctx context.Context, login, pass string) error {
	conn := getConn()
	defer conn.Close()
	c := rebPB.NewRegisterClient(conn)

	_, err := c.Register(ctx, &rebPB.RegData{
		Login:    login,
		Password: pass,
	})
	return err
}
func (cli *regClient) ChangePass(ctx context.Context, login, oldPass, newPass string) error {
	conn := getConn()
	defer conn.Close()
	c := rebPB.NewRegisterClient(conn)

	_, err := c.ChangePass(ctx, &rebPB.ChangeData{
		Login:   login,
		OldPass: oldPass,
		NewPass: newPass,
	})
	return err
}
