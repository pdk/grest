package grest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GrestGetter[Response any] func(map[string]string) (Response, error)
type GrestHandler[Request, Response any] func(map[string]string, Request) (Response, error)

func Getter[Response any](h GrestGetter[Response]) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		pathVars := mux.Vars(r)

		resp, err := h(pathVars)
		if err != nil {
			WriteError(w, fmt.Sprintf("failed to handle request: %v", err))
			return
		}

		WriteJSON(w, resp)
	}
}

func Handler[Request, Response any](h GrestHandler[Request, Response]) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			WriteError(w, fmt.Sprintf("failed to parse request (type %T): %v", req, err))
			return
		}

		resp, err := h(mux.Vars(r), req)
		if err != nil {
			WriteError(w, fmt.Sprintf("failed to handle request (type %T): %v", req, err))
			return
		}

		WriteJSON(w, resp)
	}
}

func WriteError(w http.ResponseWriter, message string) {

	w.WriteHeader(http.StatusInternalServerError)

	WriteJSON(w, map[string]string{
		"status":  "ERROR",
		"message": message,
	})
}

func WriteJSON(w http.ResponseWriter, content any) {

	w.Header().Set("Content-Type", "application/json")

	marshalled, err := json.Marshal(content)
	if err != nil {
		log.Fatalf("failed to marshal content (type %T) to JSON: %v", content, err)
	}

	_, err = w.Write(marshalled)
	if err != nil {
		log.Printf("failed to write content (%s) to client: %v", marshalled, err)
	}
}
