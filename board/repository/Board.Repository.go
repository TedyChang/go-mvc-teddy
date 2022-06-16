package repository

import (
	"codetest/board/entity"
	"codetest/database"
	"fmt"
)

type BoardRepositoryImpl struct {
	Model *entity.Rows
}

func NewRepository() BoardRepository {
	return BoardRepositoryImpl{&database.Db.TBoard}
}

type BoardRepository interface {
	FindById(id int64) entity.Board
	FindAll() []entity.Board
	Save(board entity.Board) int64
}

func (r BoardRepositoryImpl) FindById(id int64) entity.Board {
	result := entity.Board{}
	for _, board := range r.Model.BoardRows {
		if isBoardId(board, id) {
			result = board
			break
		}
	}
	return result
}

func (r BoardRepositoryImpl) FindAll() []entity.Board {
	arr := make([]entity.Board, len(r.Model.BoardRows))

	for i, board := range r.Model.BoardRows {
		arr[i] = board
	}
	return arr
}

func (r BoardRepositoryImpl) Save(board entity.Board) int64 {
	r.Model.BoardRows = append(r.Model.BoardRows, board)
	fmt.Println("save : ", board.Title)
	return board.Id
}

func isBoardId(board entity.Board, id int64) bool {
	return board.Id == id
}
