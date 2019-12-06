package repository

import (
	"time"

	"go_mongo_iris/src/modules/profile/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type profileRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewProfileRepositoryMongo(db *mgo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (p *profileRepositoryMongo) Save(profile *model.Profile) error {
	err := p.db.C(p.collection).Insert(profile)
	return err
}
func (p *profileRepositoryMongo) Update(id string, profile *model.Profile) error {
	profile.UpdatedAt = time.Now()
	// Update data by ID
	err := p.db.C(p.collection).Update(bson.M{`bson="id"`: id}, profile)
	return err
}
func (p *profileRepositoryMongo) Delete(id string) error {
	err := p.db.C(p.collection).Remove(bson.M{`bson="id"`: id})
	return err
}

func (p *profileRepositoryMongo) FindByID(id string) (*model.Profile, error) {
	var profile model.Profile
	err := p.db.C(p.collection).Find(bson.M{`bson="id"`: id}).One(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (p *profileRepositoryMongo) FindByEmail(email string) (*model.Profile, error) {
	var profile model.Profile
	err := p.db.C(p.collection).Find(bson.M{`bson="email"`: email}).One(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (p *profileRepositoryMongo) FindAll() (model.Profiles, error) {
	var profiles model.Profiles
	err := p.db.C(p.collection).Find(bson.M{}).All(&profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}
