package basestations

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Basestation contains all info of a specific basestation
type Basestation struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	BSSID string        `json:"BSSID,omitempty" bson:"BSSID,omitempty" form:"BSSID" query:"BSSID"`
	X     int           `json:"x" bson:"x" form:"x" query:"x"`
	Y     int           `json:"y" bson:"y" form:"y" query:"y"`
	MapID bson.ObjectId `json:"mapID,omitempty" bson:"mapID,omitempty" form:"mapID" query:"mapID"`
}

// Basestations contains the methods for the basestations database
type Basestations struct {
	database *mgo.Database
}

// New returns a new Basestations interface for the database connection
func New(database *mgo.Database) Basestations {
	return Basestations{database: database}
}

// GetAll gets all basestations in the database
func (b *Basestations) GetAll() ([]Basestation, error) {
	c := b.database.C("basestations").With(b.database.Session.Copy())

	result := []Basestation{}
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Add adds a basestation to the database
func (b *Basestations) Add(basestation *Basestation) error {
	c := b.database.C("basestations").With(b.database.Session.Copy())
	return c.Insert(basestation)
}