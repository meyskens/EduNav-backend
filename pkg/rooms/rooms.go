package rooms

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Room contains all info for a room
type Room struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string        `json:"name,omitempty" bson:"name,omitempty"`
	MapID   bson.ObjectId `json:"mapID,omitempty" bson:"mapID,omitempty"`
	X       float64       `json:"x" bson:"x"`
	Y       float64       `json:"y" bson:"y"`
	Comment string        `json:"comment,omitempty" bson:"comment,omitempty"`
}

// Rooms returns the methods for the maps database
type Rooms struct {
	database *mgo.Database
}

// New returns a new Maps interface for the database connection
func New(database *mgo.Database) Rooms {
	return Rooms{database: database}
}

// Get gets the Map for the given id
func (r *Rooms) Get(id string) (Room, error) {
	c := r.database.C("rooms").With(r.database.Session.Copy())

	result := Room{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetForTerm gives all rooms with a given term in the name
func (r *Rooms) GetForTerm(term string) ([]Room, error) {
	c := r.database.C("rooms").With(r.database.Session.Copy())
	result := []Room{}
	err := c.Find(bson.M{"name": bson.RegEx{Pattern: term, Options: "i"}}).Sort("name").All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
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
	err := c.Find(bson.M{"mapID": bson.ObjectIdHex(mapID)}).Sort("name").All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetAll gets all rooms in the database
func (r *Rooms) GetAll() ([]Room, error) {
	c := r.database.C("rooms").With(r.database.Session.Copy())

	result := []Room{}
	err := c.Find(bson.M{}).Sort("name").All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Add adds a room to the database
func (r *Rooms) Add(room *Room) error {
	c := r.database.C("rooms").With(r.database.Session.Copy())
	return c.Insert(room)
}
