package main

import (
  "fmt"
  "net/http"
)

type Strang string

func (s Strang) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprint(w, s)
}


type Welcome struct {
  Greeting string
  Punct string
  Who string
}

func (h Welcome) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprint(w, h)
}

func main() {
  // your http.Handle calls here
  http.Handle("/string", Strang("I'm a frayed knot."))
  http.Handle("/struct", &Welcome{"Hello", ":", "Gophers!"})
  http.ListenAndServe("localhost:4000", nil)
}
