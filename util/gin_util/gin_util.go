package gin_util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadJson(c *gin.Context, s map[string]any) {
	c.JSON(http.StatusBadRequest, s)
}

func OkJson(c *gin.Context, s map[string]any) {
	c.JSON(http.StatusOK, s)
}
