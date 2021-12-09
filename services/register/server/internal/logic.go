package internal

import (
	"context"
	regPb "etby/services/register"
	"fmt"
)

type Server struct {
	regPb.RegisterServer
}

func (s *Server) Register(ctx context.Context, in *regPb.RegData) (*regPb.Response, error) {
	_, _, err := GetUser(in.Login)
	if err == nil {
		return &regPb.Response{}, fmt.Errorf("login already used")
	}
	hash := regPb.Hash(in.Login, in.Password)
	err = AddUser(in.Login, hash)
	return &regPb.Response{}, err
}
func (s *Server) ChangePass(ctx context.Context, in *regPb.ChangeData) (*regPb.Response, error) {
	dbHash, id, err := GetUser(in.Login)
	if err != nil {
		return &regPb.Response{}, err
	}
	hash := regPb.Hash(in.Login, in.OldPass)
	if dbHash != hash {
		return &regPb.Response{}, fmt.Errorf("not auth to change pass")
	}
	hash = regPb.Hash(in.Login, in.NewPass)
	err = UpdatePass(hash, id)
	return &regPb.Response{}, err
}
