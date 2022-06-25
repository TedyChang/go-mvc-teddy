package gu

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BadJson(c *gin.Context, s map[string]any) {
	c.JSON(http.StatusBadRequest, s)
}

func OkJson(c *gin.Context, s map[string]any) {
	c.JSON(http.StatusOK, s)
}

func GetID(key string, c *gin.Context) (int64, error) {
	key = c.Param(key)
	return strconv.ParseInt(key, 10, 64)
}

type Gunc func(c *gin.Context)

func (r Gunc) ToFunc() func(c *gin.Context) { return r }
