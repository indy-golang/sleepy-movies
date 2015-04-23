package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	wsContainer := restful.NewContainer()
	m := MovieResource{}
	m.Register(wsContainer)
	g := GenreResource{}
	g.Register(wsContainer)
  a := ActorResource{}
  a.Register(wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
