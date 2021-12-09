package client

import (
	"context"
	"log"
	"os"

	loginPb "etby/services/login"

	"google.golang.org/grpc"
)

var address string //= "localhost:50051"

func init() {
	host := os.Getenv("LOGIN_HOST")
	port := os.Getenv("LOGIN_PORT")
	address = host + ":" + port
}

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

type loginClient struct{}

func GetLoginClient() loginClient {
	return loginClient{}
}
func (cli *loginClient) Login(ctx context.Context, login, pass string) (authToken, refToken string, err error) {
	conn := getConn()
	defer conn.Close()
	c := loginPb.NewAuthenticationClient(conn)
	tok, err := c.Login(ctx, &loginPb.AuthData{
		Login:    login,
		Password: pass,
	})
	if err != nil {
		return "", "", err
	}
	return tok.AuthToken, tok.RefreshToken, err
}
func (cli *loginClient) RefreshToken(ctx context.Context, token string) (authToken, refToken string, err error) {
	conn := getConn()
	defer conn.Close()
	c := loginPb.NewAuthenticationClient(conn)
	tok, err := c.Refresh(ctx, &loginPb.RefreshToken{
		RefreshToken: token,
	})
	if err != nil {
		return "", "", err
	}
	return tok.AuthToken, tok.RefreshToken, err
}
