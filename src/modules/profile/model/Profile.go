package model

import (
	"time"
)

type Profile struct {
	ID           string    `bson="id"`
	FirstName    string    `bson="first_name"`
	LastName     string    `bson="last_name"`
	Email        string    `bson="email"`
	Password     string    `bson="password"`
	ImageProfile string    `bson="image_profile"`
	CreatedAt    time.Time `bson="created_at"`
	UpdatedAt    time.Time `bson="updated_at"`
}

type Profiles []Profile

func (p *Profile) IsValidPassword(password string) bool {
	return p.Password == password
}
