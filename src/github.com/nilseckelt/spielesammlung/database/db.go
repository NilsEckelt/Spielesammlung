package database

import (
	"log"

	"github.com/byuoitav/hateoas"
	"github.com/nilseckelt/spielesammlung/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

func (db *Db) AddGame(game models.Game) {

	db.session.SetMode(mgo.Monotonic, true)

	c := db.session.DB("test").C("games")
	err := c.Insert(game)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Db) RetrieveGames() []models.Game {
	gameCollection := []models.Game{}

	c := db.session.DB("test").C("games")
	err := c.Find(bson.M{}).All(&gameCollection)
	if err != nil {
		log.Fatal(err)
	}

	hateoas.Load("https://raw.githubusercontent.com/NilsEckelt/Spielesammlung/master/src/swagger.yml")

	for i := 0; i < len(gameCollection); i++ {
		links, err := hateoas.AddLinks("/gc", []string{gameCollection[i].Name})
		if err != nil {
			log.Fatal(err)
		}
		gameCollection[i].Links = links
	}

	return gameCollection
}

func (db *Db) Find(name string) (*models.Game, error) {
	game := models.Game{}

	c := db.session.DB("test").C("games")
	err := c.Find(bson.M{"name": name}).One(&game)
	if err != nil {
		return nil, err
	}
	links, err := hateoas.AddLinks("/gc", []string{name})
	game.Links = links

	return &game, nil
}
