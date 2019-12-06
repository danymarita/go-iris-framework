package controller

import (
	"go_mongo_iris/src/modules/story/usecase"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type StoryController struct {
	Ctx          iris.Context
	StoryUsecase usecase.StoryUsecase
}

// Function ini akan otomatis dimapping ke url localhost:3000/
func (c *StoryController) Get() mvc.Result {
	stories, err := c.StoryUsecase.GetAll()
	if err != nil {
		return mvc.View{
			Name: "index.html",
			Data: iris.Map{"Title": "Stories"},
		}
	}

	return mvc.View{
		Name: "index.html",
		Data: iris.Map{"Title": "Stories", "Stories": stories},
	}
}
