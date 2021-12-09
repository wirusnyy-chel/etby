package client

import (
	"context"
	friendsPB "etby/services/friends"
	"log"

	"google.golang.org/grpc"
)

const address = "localhost:50054"

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

type friendsClient struct{}

func GetUserClient() friendsClient {
	return friendsClient{}
}

func GetFriendList(ctx context.Context, id int64) error {
	return nil
}

//Requered current user id in ctx
func AddFriend(ctx context.Context, id int64) error {
	return nil
}

//Requered current user id in ctx
func ConfirmFriend(ctx context.Context, id int64) error {
	return nil
}

//Requered current user id in ctx
func DeleteFriend(ctx context.Context, id int64) error {
	return nil
}

//Requered current user id in ctx
func RangeToFriend(ctx context.Context, id int64) (int64, error) {
	conn := getConn()
	defer conn.Close()
	c := friendsPB.NewFriendsClient(conn)
	i, err := c.RangeToFriend(ctx, &friendsPB.UserID{Id: id})
	return i.Range, err
}
