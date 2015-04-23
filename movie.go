package main

import (
	"github.com/emicklei/go-restful"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Movie struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Title       string
	GenreID     string
	Genre       Genre `bson:"-"`
	Released    string
	Description string
	Cast        []Actor `bson:"-"`
	CastIDs     []string `xml:"CastIDs>CastID"`
}

type MovieResource struct { /* empty */
}

func (MovieResource) genre() *mgo.Collection {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session.DB("sleepy-movies").C("genre")
}

func (MovieResource) actors() *mgo.Collection {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session.DB("sleepy-movies").C("actors")
}

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

	ws.Route(ws.PUT("/{movie-id}").To(u.update).
		Doc("update a movie").
		Operation("update").
		Param(ws.PathParameter("movie-id", "identifier of the movie").DataType("string")).
		Returns(409, "duplicate movie-id", nil).
		Reads(Movie{}))

	ws.Route(ws.POST("/").To(u.create).
		Doc("create a movie").
		Operation("create").
		Reads(Movie{}))

	ws.Route(ws.DELETE("/{movie-id}").To(u.delete).
		Doc("delete a movie").
		Operation("delete").
		Param(ws.PathParameter("movie-id", "identifier of the movie").DataType("string")))

	container.Add(ws)
}

func (u MovieResource) getGenre(genreID string) (string, Genre) {
  id := ""
  var genre Genre
  if len(genreID) == 24 {
    if bsonID := bson.ObjectIdHex(genreID); bsonID.Valid() {
      u.genre().FindId(bsonID).One(&genre)
      id = genreID
    }
  }
  return id, genre
}

func (u MovieResource) getCast(castIDs []string) ([]string, []Actor) {
  var ids []string
  var actors []Actor
  for _, actorID := range castIDs {
  	if len(actorID) == 24 {
	    if bsonID := bson.ObjectIdHex(actorID); bsonID.Valid() {
	    	var actor Actor
	      u.actors().FindId(bsonID).One(&actor)
	      actors = append(actors, actor)
	      ids = append(ids, actorID)
	    }
  	}
  }
  return ids, actors
}

func (u MovieResource) findMany(request *restful.Request, response *restful.Response) {
	var movies []Movie
	u.coll().Find(nil).All(&movies)
  for i, movie := range movies {
    movies[i].GenreID, movies[i].Genre = u.getGenre(movie.GenreID)
  	movies[i].CastIDs, movies[i].Cast = u.getCast(movie.CastIDs)
  }
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
  movie.GenreID, movie.Genre = u.getGenre(movie.GenreID)
  movie.CastIDs, movie.Cast = u.getCast(movie.CastIDs)
	response.WriteEntity(movie)
}

// POST http://localhost:8080/movies
// <Movie><Title>Melissa</Title></Movie>
//
func (u *MovieResource) create(request *restful.Request, response *restful.Response) {
	movie := new(Movie)
	if err := request.ReadEntity(movie); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	if !movie.ID.Valid() {
		movie.ID = bson.NewObjectId()
	}
  movie.GenreID, movie.Genre = u.getGenre(movie.GenreID)
  movie.CastIDs, movie.Cast = u.getCast(movie.CastIDs)
	if err := u.coll().Insert(movie); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error() + "\n")
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(movie)
}

// PUT http://localhost:8080/movies/1
// <Movie><Id>1</Id><Title>Melissa Raspberry</Title></Movie>
func (u *MovieResource) update(request *restful.Request, response *restful.Response) {
	movie := new(Movie)
	if err := request.ReadEntity(movie); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	if !movie.ID.Valid() {
		movie.ID = bson.ObjectIdHex(request.PathParameter("movie-id"))
	}
  movie.GenreID, movie.Genre = u.getGenre(movie.GenreID)
  movie.CastIDs, movie.Cast = u.getCast(movie.CastIDs)
	if _, err := u.coll().UpsertId(movie.ID, movie); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.WriteEntity(movie)
}

// DELETE http://localhost:8080/movies/1
//
func (u *MovieResource) delete(request *restful.Request, response *restful.Response) {
	if err := u.coll().RemoveId(bson.ObjectIdHex(request.PathParameter("movie-id"))); err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}
