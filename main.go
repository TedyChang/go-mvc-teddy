package main

import (
	"codetest/board"
	"codetest/database"
)

func main() {
	database.Db.DbSetting()

	// mvc struct, interface & DI
	boardControllerDI := board.InitializeController()

	// test boardService
	// session 001 board service di

	boardControllerDI.SaveBoard(7, "컨트롤러에서 올린 글입니다.", "내용은 없습니다")
	boardControllerDI.GetById(3)

}
