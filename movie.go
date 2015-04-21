package main

import (
  "net/url"
  "net/http"
  "fmt"
)

type Movie struct { /* empty */ }

func (movie Movie) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
  movies := []string{"Gone with the Wind", "The Manchurian Candidate"}
  data := map[string][]string{"movies": movies}
  return 200, data, http.Header{"Content-type": {"application/json"}}
}

func (movie Movie) Post(values url.Values, headers http.Header) (int, interface{}, http.Header) {
  fmt.Println(values)
  movies := []string{"Gone with the Wind", "The Manchurian Candidate"}
  data := map[string][]string{"movies": movies}
  return 200, data, http.Header{"Content-type": {"application/json"}}
}
