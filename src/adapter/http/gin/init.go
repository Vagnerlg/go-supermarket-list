package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine
var itemRoute *gin.RouterGroup

type GinHttp struct{}

func init() {
	r = gin.New()
	itemRoute = r.Group("/item", authorizationJwt())
}

func (gh GinHttp) Run() {
	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
