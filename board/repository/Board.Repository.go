package repository

import (
	"codetest/board/entity"
	"fmt"
)

type BoardRepository struct {
	Model *entity.Rows
}

func (r BoardRepository) FindById(id int64) {
	for _, board := range r.Model.Rows {
		if isId(board, id) {
			fmt.Println(board)
		}
	}
}

func (r BoardRepository) FindAll() {
	for _, board := range r.Model.Rows {
		fmt.Println(board)
	}
}

func (r *BoardRepository) Save(board entity.Board) {
	r.Model.Rows = append(r.Model.Rows, board)
	fmt.Println("save : ", board.Title)
}

func isId(board entity.Board, id int64) bool {
	return board.Id == id
}
