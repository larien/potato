package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/larien/potato/service"
	"github.com/larien/potato/utils/request"
	"github.com/larien/potato/utils/request/params"
)

var (
	serviceNew = service.New
)

func GetPotatoes(w http.ResponseWriter, r *http.Request) {
	potatoes, err := serviceNew().List(params.New(r))
	if err != nil {
		request.Error(w, http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}

	request.Success(w, http.StatusOK, potatoes)
}

func GetPotatoByID(w http.ResponseWriter, r *http.Request) {
	id := request.URLParam(r, "id")

	potato, err := serviceNew().Get(id)
	if err != nil {
		request.Error(w, http.StatusInternalServerError, nil)
		log.Println(err)
		return
	}
	if potato.Name == "" {
		request.Error(w, http.StatusNotFound, nil)
		return
	}

	request.Success(w, http.StatusOK, potato)
}

func CreatePotato(w http.ResponseWriter, r *http.Request) {
	var potato V1Potato
	if err := request.Decode(r, &potato); err != nil {
		request.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := serviceNew().Create(potato.toPotato()); err != nil {
		if errors.Is(err, service.ErrAlreadyExists) {
			request.Error(w, http.StatusConflict, err)
			return
		}
		request.Error(w, http.StatusInternalServerError, nil)
		return
	}

	request.Success(w, http.StatusCreated, potato)
}

func UpdatePotato(w http.ResponseWriter, r *http.Request) {
	id := request.URLParam(r, "id")

	var potato V1Potato
	if err := request.Decode(r, &potato); err != nil {
		request.Error(w, http.StatusBadRequest, err)
		return
	}
	potato.Name = id

	if err := serviceNew().Update(potato.toPotato()); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			request.Error(w, http.StatusNotFound, nil)
			return
		}
		request.Error(w, http.StatusInternalServerError, nil)
		return
	}

	request.Success(w, http.StatusOK, potato)
}

func DeletePotato(w http.ResponseWriter, r *http.Request) {
	id := request.URLParam(r, "id")

	if err := serviceNew().Delete(id); err != nil {
		if errors.Is(err, service.ErrNotFound) {
			request.Error(w, http.StatusNotFound, nil)
			return
		}
		request.Error(w, http.StatusInternalServerError, nil)
		return
	}

	request.Success(w, http.StatusNoContent, nil)
}
