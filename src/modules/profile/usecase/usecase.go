package usecase

import (
	"go_mongo_iris/src/modules/profile/model"
	storyModel "go_mongo_iris/src/modules/story/model"
)

type ProfileUsecase interface {
	SaveProfile(*model.Profile) (*model.Profile, error)
	UpdateProfile(string, *model.Profile) (*model.Profile, error)
	GetByID(string) (*model.Profile, error)
	GetByEmail(string) (*model.Profile, error)

	// Story
	CreateStory(*storyModel.Story) (*storyModel.Story, error)
}
