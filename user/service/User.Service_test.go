package service

import (
	"codetest/user/dto"
	"codetest/user/entity"
	"codetest/user/repository"
	"reflect"
	"testing"
)

type MockUserModel struct{}

func (r MockUserModel) First(dest interface{}, _ ...interface{}) repository.UserModel {
	user := entity.User{
		ID:   20,
		Name: "teddyChang",
	}
	*dest.(*entity.User) = user
	return r
}
func (r MockUserModel) Find(dest interface{}, _ ...interface{}) repository.UserModel {
	user1 := entity.User{
		ID:   1,
		Name: "teddy",
	}
	user2 := entity.User{
		ID:   2,
		Name: "warmoil",
	}
	user3 := entity.User{
		ID:   3,
		Name: "unions",
	}
	*dest.(*[]entity.User) = []entity.User{user1, user2, user3}
	return r
}
func (r MockUserModel) Create(value interface{}) repository.UserModel {
	user := entity.User{
		ID:   20,
		Name: "teddyChang",
	}
	*value.(*entity.User) = user
	return r
}

var mockUser = repository.UserRepository{UserModel: MockUserModel{}}

func TestUserService_GetById(t *testing.T) {
	type fields struct {
		Repository repository.UserRepository
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   entity.User
	}{
		{
			name:   "user > getById success",
			fields: fields{mockUser},
			args:   args{int64(20)},
			want: entity.User{
				ID:   20,
				Name: "teddyChang",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserService{
				Repository: tt.fields.Repository,
			}
			if got := r.GetById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Save(t *testing.T) {
	type fields struct {
		Repository repository.UserRepository
	}
	type args struct {
		dto dto.SaveUserDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name:   "user > save success",
			fields: fields{mockUser},
			args:   args{},
			want:   int64(20),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserService{
				Repository: tt.fields.Repository,
			}
			if got := r.Save(tt.args.dto); got != tt.want {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetAll(t *testing.T) {
	user1 := entity.User{
		ID:   1,
		Name: "teddy",
	}
	user2 := entity.User{
		ID:   2,
		Name: "warmoil",
	}
	user3 := entity.User{
		ID:   3,
		Name: "unions",
	}
	type fields struct {
		Repository repository.UserRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []entity.User
	}{
		{
			name:   "user > getAll success",
			fields: fields{mockUser},
			want:   []entity.User{user1, user2, user3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UserService{
				Repository: tt.fields.Repository,
			}
			if got := r.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
