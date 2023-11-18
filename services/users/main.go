package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("%s", id),
		})
	})

	router.Run(":8002")
}
