package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pdk/grest"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/{greeting}", grest.Handler(PostExample)).Methods("POST")
	r.HandleFunc("/{greeting}", grest.Getter(GetExample)).Methods("GET")

	log.Printf("starting http server at :8080")
	err := http.ListenAndServe(":8080", r)
	log.Printf("server stopped: %v", err)
}

type ExampleRequest struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	FavoriteFood string `json:"favoriteFood"`
}

type ExampleResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func PostExample(vars map[string]string, request ExampleRequest) (ExampleResponse, error) {

	greeting := vars["greeting"]

	log.Printf("PostExample invoked with greeting=%#v and request=%#v", greeting, request)

	return ExampleResponse{
		Status: "OK",
		Message: fmt.Sprintf("%s, %s! You are %d years old. Your favorite food is %s.",
			greeting, request.Name, request.Age, request.FavoriteFood),
		Timestamp: time.Now().UTC().String(),
	}, nil
}

func GetExample(vars map[string]string) (ExampleResponse, error) {

	greeting := vars["greeting"]

	log.Printf("GetExample invoked with greeting=%#v", greeting)

	return ExampleResponse{
		Status:    "OK",
		Message:   fmt.Sprintf("%s!", greeting),
		Timestamp: time.Now().UTC().String(),
	}, nil
}
