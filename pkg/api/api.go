package api

import (
	"net/http"

	"../basestations"
	"../config"
	"../database"
	"../maps"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

var (
	db   *mgo.Database
	conf config.ConfigurationInfo
)

// Run starts the HTTP sercer
func Run() {
	conf = config.GetConfiguration()
	db = database.GetDatabase(conf)

	e := echo.New()
	e.GET("/", getRoot)
	e.GET("/maps", getMaps)
	e.POST("/basestations/:key/add", addBaseStation)
	e.Logger.Fatal(e.Start(":8080"))
}

// e.GET("/", getRoot)
func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "EduNav backend")
}

// e.GET("/maps", getUser)
func getMaps(c echo.Context) error {
	m := maps.New(db)
	allMaps, err := m.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, allMaps)
}

// e.POST("/basestations", addBaseStation)
func addBaseStation(c echo.Context) error {
	if c.Param("key") != conf.APIToken {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	bs := basestations.New(db)
	b := new(basestations.Basestation)
	if err := c.Bind(b); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err := bs.Add(b); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, b)
}
