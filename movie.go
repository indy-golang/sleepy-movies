package main

import (
	"net/http"
	"gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/emicklei/go-restful"
)

type Movie struct {
  ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string
	Genre       string
	Released    string
	Description string
	Cast        []string
}

type MovieResource struct { /* empty */}

func (MovieResource) coll() *mgo.Collection {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  return session.DB("sleepy-movies").C("movies")
}

func (u MovieResource) Register(container *restful.Container) {
  ws := new(restful.WebService)
  ws.Path("/movies").
    Doc("Manage Movies").
    Consumes(restful.MIME_XML, restful.MIME_JSON).
    Produces(restful.MIME_JSON, restful.MIME_XML)

  ws.Route(ws.GET("/").To(u.findMany).
    Doc("get all movies").
    Operation("findMany").
    Writes([]Movie{}))

  ws.Route(ws.GET("/{movie-id}").To(u.findOne).
    Doc("get a movie").
    Operation("findOne").
    Param(ws.PathParameter("movie-id", "identifier of the movie").DataType("string")).
    Writes(Movie{}))

  // ws.Route(ws.PUT("/{movie-id}").To(u.update).
  //   Doc("update a movie").
  //   Operation("update").
  //   Param(ws.PathParameter("movie-id", "identifier of the movie").DataType("string")).
  //   Returns(409, "duplicate movie-id", nil).
  //   Reads(Movie{}))

  // ws.Route(ws.POST("").To(u.create).
  //   Doc("create a movie").
  //   Operation("create").
  //   Reads(Movie{}))

  // ws.Route(ws.DELETE("/{movie-id}").To(u.remove).
  //   Doc("delete a movie").
  //   Operation("remove").
  //   Param(ws.PathParameter("movie-id", "identifier of the movie").DataType("string")))

  container.Add(ws)
}

func (u MovieResource) findMany(request *restful.Request, response *restful.Response) {
  var movies []Movie
  u.coll().Find(nil).All(&movies)
  response.WriteEntity(movies)
}

// GET http://localhost:8080/movies/1
//
func (u MovieResource) findOne(request *restful.Request, response *restful.Response) {
  var movie Movie
  movieID := bson.ObjectIdHex(request.PathParameter("movie-id"))

  if u.coll().FindId(movieID).One(&movie) != nil {
    response.AddHeader("Content-Type", "text/plain")
    response.WriteErrorString(http.StatusNotFound, "404: Movie could not be found.")
    return
  }
  response.WriteEntity(movie)
}

// // POST http://localhost:8080/movies
// // <User><Name>Melissa</Name></User>
// //
// func (u *MovieResource) create(request *restful.Request, response *restful.Response) {
//   usr := new(User)
//   err := request.ReadEntity(usr)
//   if err != nil {
//     response.AddHeader("Content-Type", "text/plain")
//     response.WriteErrorString(http.StatusInternalServerError, err.Error())
//     return
//   }
//   usr.Id = strconv.Itoa(len(u.users) + 1) // simple id generation
//   u.users[usr.Id] = *usr
//   response.WriteHeader(http.StatusCreated)
//   response.WriteEntity(usr)
// }

// // PUT http://localhost:8080/movies/1
// // <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
// //
// func (u *MovieResource) update(request *restful.Request, response *restful.Response) {
//   usr := new(User)
//   err := request.ReadEntity(&usr)
//   if err != nil {
//     response.AddHeader("Content-Type", "text/plain")
//     response.WriteErrorString(http.StatusInternalServerError, err.Error())
//     return
//   }
//   u.users[usr.Id] = *usr
//   response.WriteEntity(usr)
// }

// // DELETE http://localhost:8080/movies/1
// //
// func (u *MovieResource) remove(request *restful.Request, response *restful.Response) {
//   id := request.PathParameter("movie-id")
//   delete(u.users, id)
// }