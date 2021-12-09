package internal

import (
	"context"
	loginPb "etby/services/login"
	reg "etby/services/register"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	loginPb.AuthenticationServer
}

func (s *Server) Login(ctx context.Context, in *loginPb.AuthData) (*loginPb.Tokens, error) {
	hash := reg.Hash(in.Login, in.Password)
	dbHash, id, err := GetUser(in.Login)
	if err != nil {
		log.WithFields(log.Fields{"id": id, "err": err}).Warn("Failed login")
		return &loginPb.Tokens{}, err
	}
	if hash != dbHash {
		log.WithField("id", id).Warn("Login with wrong pass")
		return &loginPb.Tokens{}, fmt.Errorf("wrong pass")
	}
	auth, ref, err := generateToken(id)
	log.WithField("id", id).Info("Successful login")
	return &loginPb.Tokens{
		AuthToken:    auth,
		RefreshToken: ref,
	}, err

}
func (s *Server) Refresh(ctx context.Context, in *loginPb.RefreshToken) (*loginPb.Tokens, error) {
	id, err := parseToken(in.RefreshToken)
	if err != nil {
		log.WithField("err", err).Warn("Failed refresh")
		return &loginPb.Tokens{}, err
	}
	auth, ref, err := generateToken(id)
	log.WithField("id", id).Info("Success refresh")
	return &loginPb.Tokens{
		AuthToken:    auth,
		RefreshToken: ref,
	}, err
}
