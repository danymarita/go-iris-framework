package model

import (
	"time"
	// Menggunakan alias karena nama packagenya sama
	profileModel "go_mongo_iris/src/modules/profile/model"
)

type Story struct {
	ID        string                `bson:"id"`
	Profile   *profileModel.Profile `bson:"profile"`
	Title     string                `bson:"title"`
	Content   string                `bson:"content"`
	CreatedAt time.Time             `bson:"created_at"`
	UpdatedAt time.Time             `bson:"updated_at"`
}

type Stories []Story
