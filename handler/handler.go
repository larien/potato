package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larien/potato/business"
	"github.com/larien/potato/utils/request/params"
)

var (
	businessNew = business.New
)

func GetPotatoes(w http.ResponseWriter, r *http.Request) {
	potatoes, err := businessNew().List(params.New(r))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "could not list"}`))
		return
	}

	response, _ := json.Marshal(potatoes)

	_, _ = w.Write([]byte(response))
}

func GetPotatoByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	potato := businessNew().Get(id)
	if potato.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "potato not found"}`))
		return
	}

	response, _ := json.Marshal(potato)

	_, _ = w.Write([]byte(response))
}

func CreatePotato(w http.ResponseWriter, r *http.Request) {
	var potato V1Potato
	if err := json.NewDecoder(r.Body).Decode(&potato); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "could not decode request body"}`))
		return
	}

	if err := businessNew().Create(potato.toPotato()); err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.Is(err, business.ErrAlreadyExists) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	response, _ := json.Marshal(potato)

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(response))
}

func UpdatePotato(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var potato V1Potato
	if err := json.NewDecoder(r.Body).Decode(&potato); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "could not decode request body"}`))
		return
	}
	potato.Name = id

	if err := businessNew().Update(potato.toPotato()); err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.Is(err, business.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "potato not found"}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
}

func DeletePotato(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := businessNew().Delete(id); err != nil {
		w.Header().Set("Content-Type", "application/json")
		if errors.Is(err, business.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"error": "potato not found"}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
