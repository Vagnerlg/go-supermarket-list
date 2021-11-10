package gin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/supermaketlist/src/domain"
)

func AuthorizationJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		auth = strings.ReplaceAll(auth, "Bearer ", "")
		auth = strings.Trim(auth, "")

		claim, err := domain.Validate(auth)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "email and/or password invalid",
			})
		}

		c.Set("user_id", claim.Id)
		c.Set("user_name", claim.Name)
	}
}
