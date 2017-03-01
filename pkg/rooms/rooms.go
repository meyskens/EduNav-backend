package maps

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Room contains all info for a room
type Room struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string        `json:"name,omitempty" bson:"name,omitempty"`
	MapID bson.ObjectId `json:"mapID,omitempty" bson:"mapID,omitempty"`
	X     int           `json:"x" bson:"x"`
	Y     int           `json:"y" bson:"y"`
}

// Rooms returns the methods for the maps database
type Rooms struct {
	database *mgo.Database
}

// New returns a new Maps interface for the database connection
func New(database *mgo.Database) Rooms {
	return Rooms{database: database}
}

// GetForName gets the Map for the given name
func (r *Rooms) GetForName(name string) (Room, error) {
	c := r.database.C("rooms").With(r.database.Session.Copy())

	result := Room{}
	err := c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetForMap gets all rooms in the database which are on a specific Map
func (r *Rooms) GetForMap(mapID string) ([]Room, error) {
	c := r.database.C("rooms").With(r.database.Session.Copy())

	result := []Room{}
	err := c.Find(bson.M{"mapID": mapID}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetAll gets all rooms in the database
func (r *Rooms) GetAll() ([]Room, error) {
	c := r.database.C("rooms").With(r.database.Session.Copy())

	result := []Room{}
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
