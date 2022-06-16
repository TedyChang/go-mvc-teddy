package repository

import (
	"codetest/board/entity"
	"fmt"
)

type BoardRepositoryImpl struct {
	Model *entity.Rows
}

type BoardRepository interface {
	FindById(id int64)
	FindAll() []entity.Board
	Save(board entity.Board) int64
}

func (r BoardRepositoryImpl) FindById(id int64) {
	for _, board := range r.Model.BoardRows {
		if isBoardId(board, id) {
			fmt.Println(board)
		}
	}
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
