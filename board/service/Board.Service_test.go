package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
	"testing"
)

type mockBoardRepo struct{}

func (m mockBoardRepo) FindById(id int64) {}
func (m mockBoardRepo) FindAll()          {}
func (m mockBoardRepo) Save(board entity.Board) int64 {
	return 99
}

var mock = mockBoardRepo{}

func TestBoardServiceImpl_Save(t *testing.T) {
	var mockRepo repository.BoardRepository = mock

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
			fields: fields{mockRepo},
			args: args{dto1: dto.SaveDto{
				Id:       99,
				Title:    "테스트 1",
				Contents: "테스트 내용",
			}},
			want: int64(99),
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
