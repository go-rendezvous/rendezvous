package dbstore

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func (u *userStore) Insert(user *model.User) error {
	return nil
}

func (u *userStore) Delete(userId int) error {
	return nil
}

func (u *userStore) Update(user *model.User) error {
	return nil
}

func (u *userStore) List(userIds []int) ([]*model.User, error) {
	var err error
	users := []*model.User{}
	err = u.db.Where("user_id in ?", userIds).Find(&users).Error

	return users, err
}

func newUserStore(s *datastore) store.UserStore {
	return &userStore{db: s.db}
}
