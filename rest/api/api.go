package api

import (
	"codetest/board"
	"codetest/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gunc func(c *gin.Context)

func (r Gunc) toFunc() func(c *gin.Context) { return r }

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

	r.GET("/users", UserController.GetAll().toFunc())
	r.GET("/users/:user", UserController.GetById().toFunc())
	r.POST("/users", UserController.SaveUser().toFunc())

	err := r.Run(domain)
	if err != nil {
		return
	}
}
