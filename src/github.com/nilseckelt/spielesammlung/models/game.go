package models

import "github.com/byuoitav/hateoas"

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

func (g *Game) AddHosts(host string) {
	links := []hateoas.Link{}
	for _, link := range g.Links {
		links = append(links, hateoas.Link{
			Rel:  link.Rel,
			HREF: host + link.HREF,
		})
	}
	g.Links = links
}
