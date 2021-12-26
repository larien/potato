package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Potatos map[string]Potato

type Potato struct {
	Name           string    `json:"name"`
	AddedAt        time.Time `json:"added_at"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}

var (
	potatos Potatos
	lock    sync.RWMutex
)

func GetPotatos(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(potatos)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "could not marshal response"}`))
		return
	}

	_, _ = w.Write([]byte(response))
}

func GetPotatoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	potato, ok := potatos[id]
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "potato not found"}`))
		return
	}

	response, err := json.Marshal(potato)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "could not marshal response"}`))
		return
	}

	_, _ = w.Write([]byte(response))
}

func CreatePotato(w http.ResponseWriter, r *http.Request) {
	var potato Potato
	if err := json.NewDecoder(r.Body).Decode(&potato); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "could not decode request body"}`))
		return
	}
	potato.AddedAt = time.Now()
	potato.LastModifiedAt = potato.AddedAt

	lock.Lock()
	defer lock.Unlock()
	if potatos == nil {
		potatos = make(map[string]Potato)
	}
	if _, ok := potatos[potato.Name]; ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		_, _ = w.Write([]byte(`{"error": "potato already exists"}`))
		return
	}
	potatos[potato.Name] = potato

	response, err := json.Marshal(potato)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "could not marshal response"}`))
		return
	}

	_, _ = w.Write([]byte(response))
}

func DeletePotato(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	defer lock.Unlock()
	id := mux.Vars(r)["id"]

	if _, ok := potatos[id]; !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "potato not found"}`))
		return
	}
	delete(potatos, id)

	w.WriteHeader(http.StatusNoContent)
}
