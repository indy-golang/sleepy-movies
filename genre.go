package main

import (
	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Genre struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string
}

type GenreResource struct { /* empty */
}

func (GenreResource) coll() *mgo.Collection {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session.DB("sleepy-movies").C("genres")
}

func (u GenreResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/genres").
		Doc("Manage Genres").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/").To(u.findMany).
		Doc("get all genres").
		Operation("findMany").
		Writes([]Genre{}))

	ws.Route(ws.GET("/{genre-id}").To(u.findOne).
		Doc("get a genre").
		Operation("findOne").
		Param(ws.PathParameter("genre-id", "identifier of the genre").DataType("string")).
		Writes(Genre{}))

	ws.Route(ws.PUT("/{genre-id}").To(u.update).
		Doc("update a genre").
		Operation("update").
		Param(ws.PathParameter("genre-id", "identifier of the genre").DataType("string")).
		Returns(409, "duplicate genre-id", nil).
		Reads(Genre{}))

	ws.Route(ws.POST("/").To(u.create).
		Doc("create a genre").
		Operation("create").
		Reads(Genre{}))

	ws.Route(ws.DELETE("/{genre-id}").To(u.delete).
		Doc("delete a genre").
		Operation("delete").
		Param(ws.PathParameter("genre-id", "identifier of the genre").DataType("string")))

	container.Add(ws)
}

func (u GenreResource) findMany(request *restful.Request, response *restful.Response) {
	var genres []Genre
	u.coll().Find(nil).All(&genres)
	response.WriteEntity(genres)
}

// GET http://localhost:8080/genres/1
//
func (u GenreResource) findOne(request *restful.Request, response *restful.Response) {
	var genre Genre
	genreID := bson.ObjectIdHex(request.PathParameter("genre-id"))

	if u.coll().FindId(genreID).One(&genre) != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: Genre could not be found.")
		return
	}
	response.WriteEntity(genre)
}

// POST http://localhost:8080/genres
// <User><Name>Melissa</Name></User>
//
func (u *GenreResource) create(request *restful.Request, response *restful.Response) {
	genre := new(Genre)
	if err := request.ReadEntity(genre); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	if !genre.ID.Valid() {
		genre.ID = bson.NewObjectId()
	}
	if err := u.coll().Insert(genre); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(genre)
}

// PUT http://localhost:8080/genres/1
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
//
func (u *GenreResource) update(request *restful.Request, response *restful.Response) {
	genre := new(Genre)
	if err := request.ReadEntity(genre); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	if !genre.ID.Valid() {
		genre.ID = bson.ObjectIdHex(request.PathParameter("genre-id"))
	}
	if _, err := u.coll().UpsertId(genre.ID, genre); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(genre)
}

// DELETE http://localhost:8080/genres/1
//
func (u *GenreResource) delete(request *restful.Request, response *restful.Response) {
	if err := u.coll().RemoveId(bson.ObjectIdHex(request.PathParameter("genre-id"))); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}
