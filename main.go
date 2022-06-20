package main

import (
	"codetest/database"
	"codetest/rest/gin"
	"fmt"
)

func main() {
	database.GormSetting()

	gin.RestApi()
	fmt.Println()
}
