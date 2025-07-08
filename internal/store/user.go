package store

import "github.com/go-rendezvous/rendezvous/internal/model"

// store for user
type UserStore interface {
	Insert(user *model.User) error
	Delete(userId int) error
	Update(user *model.User) error
	GetUser(username string) (model.User, error)
	List(userIds []int) ([]model.User, error)
}
