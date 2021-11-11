package config

import (
	"github.com/vagnerlg/supermaketlist/src/adapter/http/gin"
	"github.com/vagnerlg/supermaketlist/src/adapter/repository/mongo"
	"github.com/vagnerlg/supermaketlist/src/port/http"
	"github.com/vagnerlg/supermaketlist/src/port/repository"
)

type Repository struct {
	Item repository.RepositoryItem
	User repository.RepositoryUser
}

type App struct {
	Http http.Http
}

func NewApp() App {
	repository := Repository{
		Item: mongo.MongoItem{},
		User: mongo.MongoUser{},
	}

	return App{
		Http: http.Http{
			Drive: gin.GinHttp{},
			Item:  gin.NewGinItem(repository.Item),
			User:  gin.NewGinUser(repository.User),
		},
	}
}
