package cupboard

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MinMax struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Game struct {
	Name   string `json:"name"`
	Player MinMax `json:"player"`
	Time   MinMax `json:"time"`
	Owner  string `json:"owner"`
	Genre  string `json:"genre"`
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
	result := []Game{}

	c := db.session.DB("test").C("games")
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
