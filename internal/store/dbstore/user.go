package dbstore

import (
	"github.com/go-rendezvous/rendezvous/internal/model"
	"github.com/go-rendezvous/rendezvous/internal/store"
	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func (s *userStore) Insert(user *model.User) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.db.Model(&user).Create(user).Error
	})
}

func (s *userStore) Delete(userId int) error {
	return s.db.Where("user_id = ?", userId).Delete(&model.User{}).Error
}

func (s *userStore) Update(user *model.User) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return s.db.Save(user).Error
	})
}

func (s *userStore) List(userIds []int) ([]model.User, error) {
	var err error
	users := []model.User{}

	err = s.db.Model(&users).Where("user_id in ?", userIds).Find(&users).Error

	return users, err
}

func (s *userStore) GetUser(username string) (model.User, error) {
	var err error
	user := model.User{}

	err = s.db.Model(&user).Where("username = ?", username).Find(&user).Error
	return user, err
}

func newUserStore(s *datastore) store.UserStore {
	return &userStore{db: s.db}
}
