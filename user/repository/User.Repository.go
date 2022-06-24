package repository

import (
	"codetest/database"
	"codetest/user/entity"
	"gorm.io/gorm"
)

type UserModel interface {
	First(dest interface{}, conds ...interface{}) UserModel
	Find(dest interface{}, conds ...interface{}) UserModel
	Create(value interface{}) UserModel
}

type UserQuery struct {
	*gorm.DB
}

func (r UserQuery) First(dest interface{}, conds ...interface{}) UserModel {
	return UserModel(UserQuery{r.DB.First(dest, conds)})
}
func (r UserQuery) Find(dest interface{}, conds ...interface{}) UserModel {
	return UserModel(UserQuery{r.DB.Find(dest, conds)})
}
func (r UserQuery) Create(value interface{}) UserModel {
	return UserModel(UserQuery{r.DB.Create(value)})
}

type UserRepository struct {
	UserModel
}

func NewUserRepository() UserRepository {
	return UserRepository{UserQuery{database.Db}}
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
