package main

import "github.com/dougblack/sleepy"

func main() {
  movie := new(Movie)

  api := sleepy.NewAPI()
  api.AddResource(movie, "/movies")
  api.Start(3000)
}
