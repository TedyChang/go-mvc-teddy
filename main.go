package main

import (
	"codetest/board"
	"codetest/database"
	"codetest/user"
)

func main() {
	database.Db.DbSetting()

	// mvc struct, interface & DI
	boardControllerDI := board.InitializeBoardController()
	ReplyControllerDI := board.InitializeReplyController()
	UserControllerDI := user.InitializeUserController()

	// test boardService
	// session 001 board service di

	boardControllerDI.SaveBoard("컨트롤러에서 올린 글입니다.", "내용은 없습니다")
	boardControllerDI.GetById(3)

	ReplyControllerDI.SaveReply(1, "댓글입니다.")
	ReplyControllerDI.SaveReply(2, "두번째 댓글입니다.")
	ReplyControllerDI.GetAllReply()

	UserControllerDI.GetAll()
}
