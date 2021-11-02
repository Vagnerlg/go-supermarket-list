package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/supermaketlist/src/adapter/repository/mongo"
	"github.com/vagnerlg/supermaketlist/src/domain"
)

var repository = mongo.MongoRepository{}

func All(c *gin.Context) {
	itens := repository.All()
	if itens == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, repository.All())
}

func Insert(c *gin.Context) {
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

	item = repository.Insert(item)

	c.JSON(http.StatusOK, item)
}

func FindById(c *gin.Context) {
	id := c.Param("id")
	item, err := repository.Fisrt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}
