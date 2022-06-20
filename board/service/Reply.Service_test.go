package service

import (
	"codetest/board/dto"
	"codetest/board/entity"
	"codetest/board/repository"
	"reflect"
	"testing"
)

type mockReplyRepositoryImpl struct{}

func (m mockReplyRepositoryImpl) FindById(id int64) entity.Reply {
	return entity.Reply{
		ID: id,
		Board: entity.Board{
			ID:       1,
			Title:    "테스트 제목 7",
			Contents: "테스트 내용 7",
		},
		Content: "안녕하세요 반갑습니다.",
	}
}

func (m mockReplyRepositoryImpl) FindAll() []entity.Reply {
	return []entity.Reply{
		{
			ID: 1,
			Board: entity.Board{
				ID:       2,
				Title:    "테스트 제목 7",
				Contents: "테스트 내용 7",
			},
			Content: "안녕하세요 반갑습니다.",
		},
		{
			ID: 2,
			Board: entity.Board{
				ID:       2,
				Title:    "테스트 제목 7",
				Contents: "테스트 내용 7",
			},
			Content: "안녕하세요 반갑습니다.2",
		},
	}
}

func (m mockReplyRepositoryImpl) Save(reply entity.Reply) int64 {
	return reply.ID
}

var mockReply repository.ReplyRepository = repository.ReplyRepository{}

func TestReplyService_GetById(t *testing.T) {
	type fields struct {
		Repository      repository.ReplyRepository
		BoardRepository repository.BoardRepository
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   entity.Reply
	}{
		{
			name:   "reply > getById success",
			fields: fields{mockReply, mockBoardRepository},
			args:   args{int64(30)},
			want: entity.Reply{
				ID: 30,
				Board: entity.Board{
					ID:       1,
					Title:    "테스트 제목 7",
					Contents: "테스트 내용 7",
				},
				Content: "안녕하세요 반갑습니다.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReplyService{
				Repository:      tt.fields.Repository,
				BoardRepository: tt.fields.BoardRepository,
			}
			if got := r.GetById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplyService_Save(t *testing.T) {
	type fields struct {
		Repository      repository.ReplyRepository
		BoardRepository repository.BoardRepository
	}
	type args struct {
		boardId int64
		dto1    dto.SaveReplyDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name:   "reply > save success",
			fields: fields{mockReply, mockBoardRepository},
			args: args{
				boardId: int64(10),
			},
			want: int64(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReplyService{
				Repository:      tt.fields.Repository,
				BoardRepository: tt.fields.BoardRepository,
			}
			if got := r.Save(tt.args.boardId, tt.args.dto1); got != tt.want {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplyService_GetAll(t *testing.T) {
	type fields struct {
		Repository      repository.ReplyRepository
		BoardRepository repository.BoardRepository
	}
	tests := []struct {
		name   string
		fields fields
		want   []entity.Reply
	}{
		{
			name: "reply > getAll success",
			fields: fields{
				mockReply, mockBoardRepository,
			},
			want: []entity.Reply{
				{
					ID: 1,
					Board: entity.Board{
						ID:       2,
						Title:    "테스트 제목 7",
						Contents: "테스트 내용 7",
					},
					Content: "안녕하세요 반갑습니다.",
				},
				{
					ID: 2,
					Board: entity.Board{
						ID:       2,
						Title:    "테스트 제목 7",
						Contents: "테스트 내용 7",
					},
					Content: "안녕하세요 반갑습니다.2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ReplyService{
				Repository:      tt.fields.Repository,
				BoardRepository: tt.fields.BoardRepository,
			}
			if got := r.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
