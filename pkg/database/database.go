package database

import (
	"../config"
	mgo "gopkg.in/mgo.v2"
)

// GetDatabase gets a new database session
func GetDatabase(conf config.ConfigurationInfo) *mgo.Database {
	session, err := mgo.Dial(conf.MongoDBURL)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session.DB("edunav")
}
