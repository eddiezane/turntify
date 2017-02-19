package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/eddiezane/turntify"
	"github.com/eddiezane/turntify/store"
)

var (
	t *turntify.Turntify
)

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "no room id provided", http.StatusBadRequest)
	}
	t.CreateRoom(id)
	w.WriteHeader(http.StatusCreated)
}

func getRoomHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	room, err := t.GetRoom(vars["id"])
	if err != nil {
		switch err.(type) {
		case store.ErrNotFound:
			http.Error(w, "not found", http.StatusNotFound)
		default:
			http.Error(w, "unknown error", http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(room)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error writing response", http.StatusInternalServerError)
	}
}

func main() {
	t = turntify.NewTurntify(store.NewLocalStore())
	r := mux.NewRouter()
	r.HandleFunc("/rooms", createRoomHandler).Methods("POST")
	r.HandleFunc("/rooms/{id}", getRoomHandler).Methods("GET")
	http.Handle("/", r)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
