package repository

import (
	"go_mongo_iris/src/modules/story/model"
)

type StoryRepository interface {
	Save(*model.Story) error
	FindByID(string) (*model.Story, error)
	FindByProfileID(string) (model.Stories, error)
	FindAll() (model.Stories, error)
}
