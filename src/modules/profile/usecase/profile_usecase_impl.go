package usecase

import (
	"go_mongo_iris/src/modules/profile/model"
	"go_mongo_iris/src/modules/profile/repository"
	storyModel "go_mongo_iris/src/modules/story/model"
	storyRepo "go_mongo_iris/src/modules/story/repository"
)

type profileUsecaseImpl struct {
	profileRepository repository.ProfileRepository
	storyRepository   storyRepo.StoryRepository
}

func NewProfileUsecase(profileRepository repository.ProfileRepository, storyRepository storyRepo.StoryRepository) *profileUsecaseImpl {
	return &profileUsecaseImpl{profileRepository: profileRepository, storyRepository: storyRepository}
}

func (p *profileUsecaseImpl) CreateStory(story *storyModel.Story) (*storyModel.Story, error) {
	err := p.storyRepository.Save(story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func (p *profileUsecaseImpl) SaveProfile(profile *model.Profile) (*model.Profile, error) {
	err := p.profileRepository.Save(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p *profileUsecaseImpl) UpdateProfile(id string, profile *model.Profile) (*model.Profile, error) {
	err := p.profileRepository.Update(id, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p *profileUsecaseImpl) GetByID(id string) (*model.Profile, error) {
	profile, err := p.profileRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p *profileUsecaseImpl) GetByEmail(email string) (*model.Profile, error) {
	profile, err := p.profileRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
