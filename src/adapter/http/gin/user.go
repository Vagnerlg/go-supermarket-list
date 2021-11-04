package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/supermaketlist/src/adapter/repository/mongo"
	"github.com/vagnerlg/supermaketlist/src/domain"
)

var userRepository = mongo.UserMongoRepository{}

func UserInsert(c *gin.Context) {
	user := domain.User{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request com erro",
		})
		return
	}

	err = domain.NewUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user = userRepository.Insert(user)

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	user := domain.User{}
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request com erro",
		})
		return
	}

	userInternal, erro := userRepository.FindByEmail(user.Email)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and/or password invalid",
		})
		return
	}

	if userInternal.ComparePassword(user.Password) == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and/or password invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": "token do jwt",
	})
}
