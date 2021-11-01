package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vagnerlg/supermaketlist/src/adapter/repository/mongo"
	"github.com/vagnerlg/supermaketlist/src/domain"
)

func All(c *gin.Context) {
	itens := mongo.All()
	if itens == nil {
		c.JSON(http.StatusOK, []string{})
		return
	}

	c.JSON(http.StatusOK, mongo.All())
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

	item = mongo.Insert(item)

	c.JSON(http.StatusOK, item)
}

func FindById(c *gin.Context) {
	id := c.Param("id")
	item, err := mongo.Fisrt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, item)
}
