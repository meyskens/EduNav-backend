package maps

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Map contains all info for a map of a building and a floor
type Map struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string        `json:"name,omitempty" bson:"name,omitempty"`
	Building      string        `json:"building,omitempty" bson:"building,omitempty"`
	Floor         int           `json:"floor" bson:"floor"` // no omit as 0 is empty for integers
	Campus        string        `json:"campus,omitempty" bson:"campus,omitempty"`
	ImageLocation string        `json:"imageLocation,omitempty" bson:"imageLocation,omitempty"`
}

// Maps returns the methods for the maps database
type Maps struct {
	database *mgo.Database
}

// New returns a new Maps interface for the database connection
func New(database *mgo.Database) Maps {
	return Maps{database: database}
}

// Get gets the Map for the given id
func (m *Maps) Get(id string) (Map, error) {
	c := m.database.C("maps").With(m.database.Session.Copy())

	result := Map{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetAll gets all maps in the database
func (m *Maps) GetAll() ([]Map, error) {
	c := m.database.C("maps").With(m.database.Session.Copy())

	result := []Map{}
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
