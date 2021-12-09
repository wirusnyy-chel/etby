package client

import (
	"context"
	userPB "etby/services/user"
	"log"

	"google.golang.org/grpc"
)

const address = "localhost:50053"

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

type userClient struct{}

func GetUserClient() userClient {
	return userClient{}
}

//Curent user requered in cxt
func (cli *userClient) GetUser(ctx context.Context, id int64) (*userPB.UserInfo, error) {
	conn := getConn()
	defer conn.Close()
	c := userPB.NewUserClient(conn)
	return c.GetUser(ctx, &userPB.UserID{Id: id})
}
func (cli *userClient) AddUser(ctx context.Context, user *userPB.UserData) error {
	conn := getConn()
	defer conn.Close()
	c := userPB.NewUserClient(conn)
	_, err := c.AddUser(ctx, user)
	return err
}

func (cli *userClient) EditUser(ctx context.Context, user *userPB.UserData) error {
	conn := getConn()
	defer conn.Close()
	c := userPB.NewUserClient(conn)
	_, err := c.EditUser(ctx, user)
	return err
}
