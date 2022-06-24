package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

type MockModel struct{}

func (m MockModel) First(dest interface{}, _ ...interface{}) repository.BoardModel {
	*dest.(*entity.Board) = entity.Board{
		ID:        7,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Title:     "테스트 제목 7",
		Contents:  "테스트 내용 7",
	}
	return m
}
func (m MockModel) Find(dest interface{}, _ ...interface{}) repository.BoardModel {
	*dest.(*[]entity.Board) = []entity.Board{
		{
			ID:       1,
			Title:    "게시판 테스트 1",
			Contents: "게시판 내용 1",
		},
		{
			ID:       2,
			Title:    "게시판 테스트 2",
			Contents: "게시판 내용 2",
		},
		{
			ID:       3,
			Title:    "게시판 테스트 3",
			Contents: "게시판 내용 3",
		},
	}
	return m
}
func (m MockModel) Create(value interface{}) repository.BoardModel {
	*value.(*entity.Board) = entity.Board{
		ID:        4,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Title:     "테스트 제목 7",
		Contents:  "테스트 내용 7",
	}
	return m
}
func (m MockModel) Where(interface{}, ...interface{}) repository.BoardModel {
	return m
}

func (m MockModel) FindAll() []entity.Board {
	return []entity.Board{
		{
			ID:       1,
			Title:    "게시판 테스트 1",
			Contents: "게시판 내용 1",
		},
		{
			ID:       2,
			Title:    "게시판 테스트 2",
			Contents: "게시판 내용 2",
		},
		{
			ID:       3,
			Title:    "게시판 테스트 3",
			Contents: "게시판 내용 3",
		},
	}
}

func (m MockModel) Save(board entity.Board) int64 {
	return board.ID
}

var mockBoardRepository = repository.BoardRepository{BoardModel: MockModel{}}

func TestBoardServiceImpl_GetById(t *testing.T) {
	type fields struct {
		Repository repository.BoardRepository
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   entity.Board
	}{
		{
			name:   "board > getById > 성공",
			fields: fields{mockBoardRepository},
			args:   args{int64(7)},
			want: entity.Board{
				ID:       7,
				Title:    "테스트 제목 7",
				Contents: "테스트 내용 7",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BoardService{
				Repository: tt.fields.Repository,
			}
			if got := r.GetById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardServiceImpl_Save(t *testing.T) {
	type fields struct {
		Repository repository.BoardRepository
	}
	type args struct {
		dto1 dto.SaveBoardDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name:   "board > save > 성공",
			fields: fields{mockBoardRepository},
			args: args{dto1: dto.SaveBoardDto{
				Title:    "테스트 1",
				Contents: "테스트 내용",
			}},
			want: int64(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BoardService{
				Repository: tt.fields.Repository,
			}
			if got := r.Save(tt.args.dto1); got != tt.want {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardServiceImpl_GetAll(t *testing.T) {
	type fields struct {
		Repository repository.BoardRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []entity.Board
	}{
		{
			name:   "board > GetAll > 성공",
			fields: fields{mockBoardRepository},
			want: []entity.Board{
				{
					ID:       1,
					Title:    "게시판 테스트 1",
					Contents: "게시판 내용 1",
				},
				{
					ID:       2,
					Title:    "게시판 테스트 2",
					Contents: "게시판 내용 2",
				},
				{
					ID:       3,
					Title:    "게시판 테스트 3",
					Contents: "게시판 내용 3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BoardService{
				Repository: tt.fields.Repository,
			}
			if got := r.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
