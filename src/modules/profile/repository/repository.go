package repository

import (
	"go_mongo_iris/src/modules/profile/model"
)

// Profile repository
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindByEmail(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
