package cupboard

import (
	"log"

	"github.com/byuoitav/hateoas"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MinMax struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Game struct {
	Links  []hateoas.Link `json:"links,omitempty"`
	Name   string         `json:"name"`
	Player MinMax         `json:"player"`
	Time   MinMax         `json:"time"`
	Owner  string         `json:"owner"`
	Genre  string         `json:"genre"`
}

type Db struct {
	session *mgo.Session
}

func New(host string) *Db {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	return &Db{
		session,
	}
}

func (db *Db) Close() {
	db.session.Close()
}

func (db *Db) AddGame(game Game) {

	db.session.SetMode(mgo.Monotonic, true)

	c := db.session.DB("test").C("games")
	err := c.Insert(game)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Db) RetrieveGames() []Game {
	gameCollection := []Game{}

	c := db.session.DB("test").C("games")
	err := c.Find(bson.M{}).All(&gameCollection)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(gameCollection); i++ {
		log.Println("Adding links for %s...", gameCollection[i].Name)
		links, err := hateoas.AddLinks("/gc", []string{gameCollection[i].Name})
		if err != nil {
			log.Fatal(err)
		}
		gameCollection[i].Links = links
	}

	return gameCollection
}
func (db *Db) GetGame(name string) Game {
	game := Game{}
	// links, err := hateoas.AddLinks(c, []string{game})
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, helpers.ReturnError(err))
	// }
	// game.Links = links
	return game
}
func (db *Db) CreateGame(name string) {}
func (db *Db) DeleteGame(name string) {}
