package main

import (
  "github.com/emicklei/go-restful"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "net/http"
)

type Actor struct {
  ID   bson.ObjectId `bson:"_id,omitempty"`
  Name string
}

type ActorResource struct { /* empty */
}

func (ActorResource) coll() *mgo.Collection {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  return session.DB("sleepy-movies").C("actors")
}

func (ActorResource) movies() *mgo.Collection {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  return session.DB("sleepy-movies").C("movies")
}

func (u ActorResource) Register(container *restful.Container) {
  ws := new(restful.WebService)
  ws.Path("/actors").
    Doc("Manage Actors").
    Consumes(restful.MIME_XML, restful.MIME_JSON).
    Produces(restful.MIME_JSON, restful.MIME_XML)

  ws.Route(ws.GET("/").To(u.findMany).
    Doc("get all actors").
    Operation("findMany").
    Writes([]Actor{}))

  ws.Route(ws.GET("/{actor-id}").To(u.findOne).
    Doc("get an actor").
    Operation("findOne").
    Param(ws.PathParameter("actor-id", "identifier of the actor").DataType("string")).
    Writes(Actor{}))

  ws.Route(ws.GET("/{actor-id}/movies").To(u.findMovies).
    Doc("get movies this actor has appeared in").
    Operation("findMovies").
    Param(ws.PathParameter("actor-id", "identifier of the actor").DataType("string")).
    Writes([]Movie{}))

  ws.Route(ws.PUT("/{actor-id}").To(u.update).
    Doc("update a actor").
    Operation("update").
    Param(ws.PathParameter("actor-id", "identifier of the actor").DataType("string")).
    Returns(409, "duplicate actor-id", nil).
    Reads(Actor{}))

  ws.Route(ws.POST("/").To(u.create).
    Doc("create a actor").
    Operation("create").
    Reads(Actor{}))

  ws.Route(ws.DELETE("/{actor-id}").To(u.delete).
    Doc("delete a actor").
    Operation("delete").
    Param(ws.PathParameter("actor-id", "identifier of the actor").DataType("string")))

  container.Add(ws)
}

func (u ActorResource) findMany(request *restful.Request, response *restful.Response) {
  var actors []Actor
  u.coll().Find(nil).All(&actors)
  response.WriteEntity(actors)
}

// GET http://localhost:8080/actors/1
//
func (u ActorResource) findOne(request *restful.Request, response *restful.Response) {
  var actor Actor
  actorID := bson.ObjectIdHex(request.PathParameter("actor-id"))

  if u.coll().FindId(actorID).One(&actor) != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusNotFound, "404: Actor could not be found.")
    return
  }
  response.WriteEntity(actor)
}

// GET http://localhost:8080/actors/1/movies
//
func (u ActorResource) findMovies(request *restful.Request, response *restful.Response) {
  var actor Actor
  actorID := bson.ObjectIdHex(request.PathParameter("actor-id"))

  if u.coll().FindId(actorID).One(&actor) != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusNotFound, "404: Actor could not be found.")
    return
  }
  var movies []Movie
  if err := u.movies().Find(bson.M{"castids": bson.M{"$in": []string{actor.ID.Hex()}}}).All(&movies); err != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusInternalServerError, err.Error())
    return
  }
  response.WriteEntity(movies)
}

// POST http://localhost:8080/actors
// <Actor><Name>Melissa</Name></Actor>
//
func (u *ActorResource) create(request *restful.Request, response *restful.Response) {
  actor := new(Actor)
  if err := request.ReadEntity(actor); err != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusInternalServerError, err.Error())
    return
  }
  if !actor.ID.Valid() {
    actor.ID = bson.NewObjectId()
  }
  if err := u.coll().Insert(actor); err != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusInternalServerError, err.Error())
    return
  }

  response.WriteHeader(http.StatusCreated)
  response.WriteEntity(actor)
}

// PUT http://localhost:8080/actors/1
// <Actor><Id>1</Id><Name>Melissa Raspberry</Name></Actor>
//
func (u *ActorResource) update(request *restful.Request, response *restful.Response) {
  actor := new(Actor)
  if err := request.ReadEntity(actor); err != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusInternalServerError, err.Error())
    return
  }
  if !actor.ID.Valid() {
    actor.ID = bson.ObjectIdHex(request.PathParameter("actor-id"))
  }
  if _, err := u.coll().UpsertId(actor.ID, actor); err != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusInternalServerError, err.Error())
    return
  }

  response.WriteHeader(http.StatusCreated)
  response.WriteEntity(actor)
}

// DELETE http://localhost:8080/actors/1
//
func (u *ActorResource) delete(request *restful.Request, response *restful.Response) {
  if err := u.coll().RemoveId(bson.ObjectIdHex(request.PathParameter("actor-id"))); err != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusInternalServerError, err.Error())
    return
  }
}
