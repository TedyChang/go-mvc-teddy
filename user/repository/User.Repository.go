package repository

import (
	"codetest/database"
	"codetest/user/entity"
	"fmt"
)

type UserRepository struct {
	Model *entity.Rows
}

func NewUserRepository() UserRepository {
	return UserRepository{&database.Db.TUser}
}

func (r UserRepository) FindById(id int64) entity.User {
	result := entity.User{}
	for _, user := range r.Model.Rows {
		if isId(user, id) {
			result = user
			break
		}
	}
	return result
}

func (r UserRepository) FindAll() []entity.User {
	arr := make([]entity.User, len(r.Model.Rows))

	for i, user := range r.Model.Rows {
		arr[i] = user
	}
	return arr
}

func (r *UserRepository) Save(user entity.User) int64 {
	r.Model.Rows = append(r.Model.Rows, user)
	fmt.Println("save : ", user.Name)
	return user.Id
}

func isId(user entity.User, id int64) bool {
	return user.Id == id
}
