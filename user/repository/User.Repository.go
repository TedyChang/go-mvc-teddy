package repository

import (
	"codetest/database"
	"codetest/user/entity"
	"gorm.io/gorm"
)

type Model interface {
	Model(interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
}

type UserRepository struct {
	Model
}

func NewUserRepository() UserRepository {
	return UserRepository{database.Db}
}

func (r UserRepository) FindById(id int64) entity.User {
	u := entity.User{}
	r.First(&u, "id = ?", id)
	return u
}

func (r UserRepository) FindAll() []entity.User {
	var users []entity.User
	r.Find(&users)
	return users
}

func (r *UserRepository) Save(user entity.User) int64 {
	_ = r.Create(&user)
	return user.ID
}
