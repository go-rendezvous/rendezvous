package dto

import (
	"time"

	"github.com/go-rendezvous/rendezvous/internal/model"
)

type UserInsertRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (r *UserInsertRequest) GenModel() *model.User {
	return &model.User{
		Username:  r.Username,
		Password:  r.Password,
		Phone:     r.Phone,
		Email:     r.Email,
		CreatedAt: time.Now(),
	}
}

type UserDeleteRequest struct {
	UserId int `param:"user_id"`
}

type UserUpdateRequest struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (r *UserUpdateRequest) GenModel() *model.User {
	return &model.User{
		UserId:    r.UserId,
		Username:  r.Username,
		Phone:     r.Phone,
		Email:     r.Email,
		CreatedAt: time.Now(),
	}
}

type UserListRequest struct {
	Ids []int `query:"ids"`
}

type UserListResponse struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func GenUserListResponse(user *model.User) UserListResponse {
	return UserListResponse{
		UserId:   user.UserId,
		Username: user.Username,
		Phone:    user.Phone,
		Email:    user.Email,
	}
}
