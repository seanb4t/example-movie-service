// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Actor struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Movies []*Movie `json:"movies,omitempty"`
}

type Director struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Movies []*Movie `json:"movies,omitempty"`
}

type Genre struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Movies []*Movie `json:"movies,omitempty"`
}

type Movie struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Year      int         `json:"year"`
	Rating    float64     `json:"rating"`
	Actors    []*Actor    `json:"actors,omitempty"`
	Directors []*Director `json:"directors,omitempty"`
	Genres    []*Genre    `json:"genres,omitempty"`
}
