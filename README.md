GET /movies
- Returns all movies. Movies have Title, Genre, Release Year, Description, Cast.
GET /movies/:id
- Returns specified movie.
GET /movies/:id/cast
- Returns a short list of the primary cast members.
PUT /movies
- Creates/updates a new movie.
DELETE /movies/:id
- Removes a movie by id
GET /genres
- Returns all genres. Genres have Name.
GET /genres/:id
- Returns specified genre.
GET /genres/:id/movies
- Returns all movies in this Genre.
PUT /genres
- Creates/updates a genre
DELETE /genre/:id
- Removes a genre by id
GET /actors
- Returns a list of actors
GET /actors/:id
- Returns a specific actor.
GET /actors/:id/movies
- Returns a list of movies this actor has been in.
PUT /actors
- Creates/updates an actor.
DELETE /actors/:id
- Removes an actor by id
