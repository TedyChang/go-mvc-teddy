// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package board

import (
	"codetest/board/controller"
	"codetest/board/repository"
	"codetest/board/service"
)

// Injectors from wire.go:

func InitializeBoardController() controller.Controller {
	boardRepository := repository.NewRepository()
	boardServiceImpl := service.NewService(boardRepository)
	controllerController := controller.NewController(boardServiceImpl)
	return controllerController
}

func InitializeReplyController() controller.ReplyController {
	replyRepository := repository.NewReplyRepository()
	boardRepository := repository.NewRepository()
	replyServiceImpl := service.NewReplyService(replyRepository, boardRepository)
	replyController := controller.NewReplyController(replyServiceImpl)
	return replyController
}
