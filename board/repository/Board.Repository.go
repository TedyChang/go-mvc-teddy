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
	FindAll()
	Save(board entity.Board) int64
}

func (r BoardRepositoryImpl) FindById(id int64) {
	for _, board := range r.Model.BoardRows {
		if isBoardId(board, id) {
			fmt.Println(board)
		}
	}
}

func (r BoardRepositoryImpl) FindAll() {
	for _, board := range r.Model.BoardRows {
		fmt.Println(board)
	}
}

func (r BoardRepositoryImpl) Save(board entity.Board) int64 {
	r.Model.BoardRows = append(r.Model.BoardRows, board)
	fmt.Println("save : ", board.Title)
	return board.Id
}

func isBoardId(board entity.Board, id int64) bool {
	return board.Id == id
}
