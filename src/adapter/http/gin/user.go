package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/supermaketlist/src/domain"
	"github.com/vagnerlg/supermaketlist/src/port/repository"
)

var repositoryUser repository.RepositoryUser

type GinUser struct{}

func NewGinUser(repository repository.RepositoryUser) GinUser {
	repositoryUser = repository

	return GinUser{}
}

func (u GinUser) Insert(route string) {
	r.POST(route, func(c *gin.Context) {
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

		user = repositoryUser.Insert(user)

		c.JSON(http.StatusOK, user)
	})
}

func (c GinUser) Login(route string) {
	r.POST(route, func(c *gin.Context) {
		user := domain.User{}
		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "request com erro",
			})
			return
		}

		userInternal, erro := repositoryUser.FindByEmail(user.Email)
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
			"token": userInternal.GenerateJWT(),
		})
	})
}
