package database

import (
	"crypto/tls"
	"net"

	"../config"
	mgo "gopkg.in/mgo.v2"
)

// GetDatabase gets a new database session
func GetDatabase(conf config.ConfigurationInfo) *mgo.Database {
	var session *mgo.Session
	var err error
	if conf.MongoUseTLS {
		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true
		dialInfo, err := mgo.ParseURL(conf.MongoDBURL)
		if err != nil {
			panic(err)
		}
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
		session, err = mgo.DialWithInfo(dialInfo)
	} else {
		session, err = mgo.Dial(conf.MongoDBURL)
	}
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session.DB("edunav")
}
