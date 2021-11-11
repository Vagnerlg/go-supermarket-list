package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/supermaketlist/src/domain"
	"github.com/vagnerlg/supermaketlist/src/port/repository"
)

var repositoryItem repository.RepositoryItem

type GinItem struct{}

func NewGinItem(repository repository.RepositoryItem) GinItem {
	repositoryItem = repository

	return GinItem{}
}

func (g GinItem) All(route string) {
	itemRoute.GET(route, func(c *gin.Context) {
		itens := repositoryItem.All()
		if itens == nil {
			c.JSON(http.StatusOK, []string{})
			return
		}

		c.JSON(http.StatusOK, repositoryItem.All())
	})
}

func (g GinItem) Insert(route string) {
	itemRoute.POST(route, func(c *gin.Context) {
		item := domain.Item{}
		err := c.ShouldBindJSON(&item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "request com erro",
			})
			return
		}

		item, err = domain.New(item)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		item = repositoryItem.Insert(item)

		c.JSON(http.StatusOK, item)
	})
}

func (g GinItem) FindById(route string) {
	itemRoute.GET(route, func(c *gin.Context) {
		id := c.Param("id")
		item, err := repositoryItem.First(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, item)
	})
}
