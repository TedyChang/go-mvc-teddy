package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
	"reflect"
	"testing"
)

type mockBoardRepo struct{}

func (m mockBoardRepo) FindById(id int64) entity.Board {
	return entity.Board{
		Id:       id,
		Title:    "테스트 제목 7",
		Contents: "테스트 내용 7",
	}
}
func (m mockBoardRepo) FindAll() []entity.Board {
	return []entity.Board{
		{
			Id:       1,
			Title:    "게시판 테스트 1",
			Contents: "게시판 내용 1",
		},
		{
			Id:       2,
			Title:    "게시판 테스트 2",
			Contents: "게시판 내용 2",
		},
		{
			Id:       3,
			Title:    "게시판 테스트 3",
			Contents: "게시판 내용 3",
		},
	}
}

func (m mockBoardRepo) Save(board entity.Board) int64 {
	return board.Id
}

var mockBoardRepository repository.BoardRepository = mockBoardRepo{}

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
				Id:       7,
				Title:    "테스트 제목 7",
				Contents: "테스트 내용 7",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BoardServiceImpl{
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
		dto1 dto.SaveDto
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
			args: args{dto1: dto.SaveDto{
				Title:    "테스트 1",
				Contents: "테스트 내용",
			}},
			want: int64(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BoardServiceImpl{
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
					Id:       1,
					Title:    "게시판 테스트 1",
					Contents: "게시판 내용 1",
				},
				{
					Id:       2,
					Title:    "게시판 테스트 2",
					Contents: "게시판 내용 2",
				},
				{
					Id:       3,
					Title:    "게시판 테스트 3",
					Contents: "게시판 내용 3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := BoardServiceImpl{
				Repository: tt.fields.Repository,
			}
			if got := r.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
