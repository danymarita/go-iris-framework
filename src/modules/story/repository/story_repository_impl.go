package repository

import (
	// "time"

	"go_mongo_iris/src/modules/story/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "gopkg.in/mgo.v2/bson"
)

type storyRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewStoryRepositoryMongo(db *mgo.Database, collection string) *storyRepositoryMongo {
	return &storyRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *storyRepositoryMongo) Save(story *model.Story) error {
	err := r.db.C(r.collection).Insert(story)
	return err
}

func (r *storyRepositoryMongo) FindByID(id string) (*model.Story, error) {
	var story model.Story
	err := r.db.C(r.collection).FindId(bson.M{`bson="id"`: id}).One(&story)
	if err != nil {
		return nil, err
	}
	return &story, nil
}

func (r *storyRepositoryMongo) FindByProfileID(profileID string) (model.Stories, error) {
	var stories model.Stories
	err := r.db.C(r.collection).Find(bson.M{`bson:"profile".id`: profileID}).All(&stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *storyRepositoryMongo) FindAll() (model.Stories, error) {
	var stories model.Stories
	err := r.db.C(r.collection).Find(bson.M{}).All(&stories)
	if err != nil {
		return nil, err
	}
	return stories, nil
}
