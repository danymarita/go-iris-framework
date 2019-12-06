package usecase

import (
	"go_mongo_iris/src/modules/story/model"
	"go_mongo_iris/src/modules/story/repository"
)

type storyUsecaseImpl struct {
	storyRepository repository.StoryRepository
}

func NewStoryUsecase(storyRepository repository.StoryRepository) *storyUsecaseImpl {
	return &storyUsecaseImpl{storyRepository}
}

func (r *storyUsecaseImpl) GetAll() (model.Stories, error) {
	var stories model.Stories
	stories, err := r.storyRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return stories, nil
}
