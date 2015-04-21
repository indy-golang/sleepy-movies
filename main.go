package main

import (
	"github.com/dougblack/sleepy"
	"gopkg.in/mgo.v2"
  "fmt"
)

func main() {
	populateMovies()
	movie := new(Movie)

	api := sleepy.NewAPI()
	api.AddResource(movie, "/movies")
	fmt.Println(api.Start(3000))
}

func populateMovies() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("sleepy-movies").C("movies")
  c.DropCollection()
	c.Insert(Movie{
		Title: "The Double",
	})
	c.Insert(Movie{
		Title: "Ferngully",
	})
	c.Insert(Movie{
		Title: "Clue",
	})
}
