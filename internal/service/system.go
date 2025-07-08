package service

import (
	"time"

	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/go-rendezvous/rendezvous/internal/pkg/util"
	"github.com/go-rendezvous/rendezvous/internal/service/dto"
	"github.com/go-rendezvous/rendezvous/internal/store/dbstore"
	"github.com/go-rendezvous/rendezvous/pkg/service"
	"github.com/golang-jwt/jwt/v4"
)

type System struct {
	service.Service
}

func (s System) Login(req *dto.LoginRequest) (string, error) {
	user, err := dbstore.Factory().UserStore().GetUser(req.Username)
	if err != nil {
		return "", err
	}

	if err = util.ComparePassword(req.Password, user.Password); err != nil {
		return "", err
	}

	expirationDuration := config.GetConf().Expiration
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserId,
		"username": user.Username,
		"exp":      time.Now().Add(time.Duration(expirationDuration)).Unix(),
	})

	return token.SignedString([]byte(config.GetConf().Secret))
}

func (s System) Logout(req *dto.LogoutRequest) error {
	return nil
}
