package internal

import (
	"context"
	friendsClient "etby/services/friends/client"
	userPB "etby/services/user"
	"fmt"
)

type Server struct {
	userPB.UserServer
}

func (s *Server) GetUser(ctx context.Context, in *userPB.UserID) (*userPB.UserInfo, error) {
	//curentID := ctx.Value(userPB.IdKey)
	id, name, privacy, err := GetUser(in.Id)
	if err != nil {
		return &userPB.UserInfo{}, err
	}
	if privacy == userPB.Privacy_all {
		return &userPB.UserInfo{
			Id:             id,
			Name:           name,
			ProfileVisible: privacy,
		}, nil
	}
	ran, err := friendsClient.RangeToFriend(ctx, in.Id)
	if err != nil {
		return &userPB.UserInfo{}, err
	}
	if ran <= int64(privacy) {
		return &userPB.UserInfo{
			Id:             id,
			Name:           name,
			ProfileVisible: privacy,
		}, nil
	}
	return &userPB.UserInfo{}, fmt.Errorf("not permited")
}
func (s *Server) AddUser(ctx context.Context, in *userPB.UserData) (*userPB.UserInfo, error) {
	return &userPB.UserInfo{}, fmt.Errorf("not implemented")
}
func (s *Server) EditUser(ctx context.Context, in *userPB.UserData) (*userPB.UserInfo, error) {
	return &userPB.UserInfo{}, fmt.Errorf("not implemented")
}
