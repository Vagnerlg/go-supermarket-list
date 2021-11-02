package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	httpgin "github.com/vagnerlg/supermaketlist/src/adapter/http/gin"
)

func main() {

	r := gin.Default()

	itemRoute := r.Group("/item")

	itemRoute.GET("/", httpgin.All)
	itemRoute.GET("/:id", httpgin.FindById)
	itemRoute.POST("/", httpgin.Insert)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
