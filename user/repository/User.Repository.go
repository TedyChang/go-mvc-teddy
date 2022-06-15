package repository

import (
	"codetest/user/entity"
	"fmt"
)

type UserRepository struct {
	Model *entity.Rows
}

func (r UserRepository) FindById(id int64) {
	for _, user := range r.Model.Rows {
		if isId(user, id) {
			fmt.Println(user)
		}
	}
}

func (r UserRepository) FindAll() {
	for _, user := range r.Model.Rows {
		fmt.Println(user)
	}
}

func (r *UserRepository) Save(user entity.User) {
	r.Model.Rows = append(r.Model.Rows, user)
	fmt.Println("save : ", user.Name)
}

func isId(user entity.User, id int64) bool {
	return user.Id == id
}
