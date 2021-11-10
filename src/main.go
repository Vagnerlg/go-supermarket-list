package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	httpgin "github.com/vagnerlg/supermaketlist/src/adapter/http/gin"
)

func main() {

	godotenv.Load()

	r := gin.New()

	itemRoute := r.Group("/item", httpgin.AuthorizationJwt())

	itemRoute.GET("/", httpgin.All)
	itemRoute.GET("/:id", httpgin.FindById)
	itemRoute.POST("/", httpgin.Insert)

	userRoute := r.Group("/user")

	userRoute.POST("/", httpgin.UserInsert)
	userRoute.POST("/login", httpgin.Login)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
