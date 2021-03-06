package api

import (
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"../basestations"
	"../config"
	"../database"
	gh "../github"
	"../maps"
	"../rooms"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	goCache "github.com/patrickmn/go-cache"
	mgo "gopkg.in/mgo.v2"
)

var (
	db    *mgo.Database
	conf  config.ConfigurationInfo
	cache *goCache.Cache
)

// Run starts the HTTP server
func Run() {
	conf = config.GetConfiguration()
	db = database.GetDatabase(conf)
	cache = goCache.New(5*time.Minute, 30*time.Second)

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", getRoot)
	e.GET("/contributors", getContributors)
	e.GET("/maps", getMaps)
	e.GET("/maps/:id", getMap)
	e.GET("/rooms/:id", getRoom)
	e.GET("/rooms/search", getRoomsForSearch)
	e.GET("/rooms/map/:mapID", getRoomsForMapID)
	e.GET("/basestation/bssid/:bssid", getBasestationsBSSID)
	e.GET("/basestations/map/:mapID", getBasestationsForMapID)
	e.POST("/basestations/:key/add", addBaseStation)
	e.POST("/rooms/:key/add", addRoom)
	if conf.AutoTLS {
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist(conf.Hostname)
		e.AutoTLSManager.Cache = autocert.DirCache(conf.CertCache)
		e.Logger.Fatal(e.StartAutoTLS(conf.Bind))
	} else {
		e.Logger.Fatal(e.Start(conf.Bind))
	}
}

// e.GET("/", getRoot)
func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "EduNav backend")
}

// e.GET("/contributors", getContributors)
func getContributors(c echo.Context) error {
	if cachedContributors, found := cache.Get("/contributors"); found {
		return c.JSON(http.StatusOK, cachedContributors.([]gh.Contributor))
	}
	contributors := gh.GetContributors()
	cache.Set("/contributors", contributors, goCache.DefaultExpiration)
	return c.JSON(http.StatusOK, contributors)
}

// e.GET("/maps", getMaps)
func getMaps(c echo.Context) error {
	m := maps.New(db)
	allMaps, err := m.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, allMaps)
}

// e.GET("/maps/:id", getMap)
func getMap(c echo.Context) error {
	m := maps.New(db)
	mapForID, err := m.Get(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, mapForID)
}

// e.GET("/rooms/:id", getRoom)
func getRoom(c echo.Context) error {
	r := rooms.New(db)
	room, err := r.Get(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, room)
}

// e.GET("/rooms/search", getRoomsForSearch)
func getRoomsForSearch(c echo.Context) error {
	r := rooms.New(db)
	rooms, err := r.GetForTerm(c.QueryParam("term"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}

// e.GET("/rooms/map/:mapID", getRoomsForMapID)
func getRoomsForMapID(c echo.Context) error {
	r := rooms.New(db)
	room, err := r.GetForMap(c.Param("mapID"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, room)
}

// e.GET("/basestation/bssid/:bssid", getBasestationsBSSID)
func getBasestationsBSSID(c echo.Context) error {
	b := basestations.New(db)
	bs, err := b.GetForBSSID(c.Param("bssid"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bs)
}

// e.GET("/basestations/map/:mapID", getBasestationsForMapID)
func getBasestationsForMapID(c echo.Context) error {
	b := basestations.New(db)
	bs, err := b.GetForMap(c.Param("mapID"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, bs)
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

// e.POST("/rooms/:key/add", addRoom)
func addRoom(c echo.Context) error {
	if c.Param("key") != conf.APIToken {
		return c.String(http.StatusUnauthorized, "Invalid API key")
	}
	r := rooms.New(db)
	room := new(rooms.Room)
	if err := c.Bind(room); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err := r.Add(room); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, room)
}
