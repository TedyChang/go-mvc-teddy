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

type MockBoardModel struct{}
type MockReplyModel struct{}

func (r MockBoardModel) First(dest interface{}, _ ...interface{}) repository.BoardModel {
	*dest.(*entity.Board) = entity.Board{ID: int64(10)}
	return r
}
func (r MockBoardModel) Find(interface{}, ...interface{}) repository.BoardModel  { return r }
func (r MockBoardModel) Create(interface{}) repository.BoardModel                { return r }
func (r MockBoardModel) Where(interface{}, ...interface{}) repository.BoardModel { return r }

func (m MockReplyModel) First(dest interface{}, _ ...interface{}) repository.ReplyModel {
	board := entity.Board{
		ID:       1,
		Title:    "테스트 제목 7",
		Contents: "테스트 내용 7",
	}
	*dest.(*entity.Reply) = entity.Reply{
		ID:        30,
		Board:     board,
		Content:   "안녕하세요 반갑습니다.",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	return m
}
func (m MockReplyModel) Find(dest interface{}, _ ...interface{}) repository.ReplyModel {
	board := entity.Board{
		ID:       1,
		Title:    "테스트 제목 7",
		Contents: "테스트 내용 7",
	}
	*dest.(*[]entity.Reply) = []entity.Reply{
		{
			ID:      1,
			Content: "안녕하세요 반갑습니다.",
			Board:   board,
		},
		{
			ID:      2,
			Content: "안녕하세요 반갑습니다.2",
			Board:   board,
		},
		{
			ID:      3,
			Content: "안녕하세요 반갑습니다.3",
			Board:   board,
		},
	}
	return m
}
func (m MockReplyModel) Create(value interface{}) repository.ReplyModel {
	*value.(*entity.Reply) = entity.Reply{
		ID:        10,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		BoardID:   0,
		Board:     entity.Board{},
		Content:   "",
	}
	return m
}
func (m MockReplyModel) Where(interface{}, ...interface{}) repository.ReplyModel {
	return m
}

var mockBoard = repository.BoardRepository{BoardModel: MockBoardModel{}}
var mockReply = repository.ReplyRepository{ReplyModel: MockReplyModel{}}

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
			fields: fields{mockReply, mockBoard},
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
	board := entity.Board{
		ID:       1,
		Title:    "테스트 제목 7",
		Contents: "테스트 내용 7",
	}

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
					ID:      1,
					Board:   board,
					Content: "안녕하세요 반갑습니다.",
				},
				{
					ID:      2,
					Board:   board,
					Content: "안녕하세요 반갑습니다.2",
				},
				{
					ID:      3,
					Board:   board,
					Content: "안녕하세요 반갑습니다.3",
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
				t.Errorf("GetAll() = %v\n, want %v", got, tt.want)
			}
		})
	}
}
