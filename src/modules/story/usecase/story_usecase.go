package usecase

import (
	"go_mongo_iris/src/modules/story/model"
)

type StoryUsecase interface {
	GetAll() (model.Stories, error)
}
