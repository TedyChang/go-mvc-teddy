package api

import (
	"codetest/board"
	"codetest/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestApi(domain string) {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	boardController := board.InitializeBoardController()
	ReplyController := board.InitializeReplyController()
	UserController := user.InitializeUserController()

	r.GET("/boards", boardController.GetAll())
	r.GET("/boards/:id", boardController.GetById())
	r.POST("/boards", boardController.SaveBoard())

	r.GET("/replies", ReplyController.GetAllReply())
	r.GET("/boards/replies/:reply", ReplyController.GetReplyById())
	r.POST("/boards/:board/replies", ReplyController.SaveReply())

	r.GET("/users", UserController.GetAll().ToFunc())
	r.GET("/users/:user", UserController.GetById().ToFunc())
	r.POST("/users", UserController.SaveUser().ToFunc())

	err := r.Run(domain)
	if err != nil {
		return
	}
}
