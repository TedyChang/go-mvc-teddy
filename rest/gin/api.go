package gin

import (
	"codetest/board"
	"codetest/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RestApi() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	boardController := board.InitializeBoardController()
	ReplyController := board.InitializeReplyController()
	UserController := user.InitializeUserController()

	r.GET("/boards", boardController.GetAll())

	r.GET("/replies", func(c *gin.Context) {
		ReplyController.GetAllReply()
	})

	r.GET("/users", func(c *gin.Context) {
		UserController.GetAll()
	})

	err := r.Run("localhost:1234")
	if err != nil {
		return
	}
}
