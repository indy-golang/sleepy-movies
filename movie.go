package main

import (
	"net/http"
	"net/url"
	"gopkg.in/mgo.v2"
)

type Movie struct {
	Title       string
	Genre       string
	Released    string
	Description string
	Cast        []string
}

func (Movie) coll() *mgo.Collection {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  return session.DB("sleepy-movies").C("movies")
}

func (movie Movie) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	var movies []Movie
  movie.coll().Find(nil).All(&movies)
	data := map[string][]Movie{"movies": movies}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}

func (movie Movie) Post(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	movies := []string{"Gone with the Wind", "The Manchurian Candidate"}
	data := map[string][]string{"movies": movies}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}
