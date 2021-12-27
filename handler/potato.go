package handler

import "github.com/larien/potato/service"

type V1Potato struct {
	Name string `json:"name"`
}

func (p V1Potato) toPotato() service.Potato {
	return service.Potato{
		Name: p.Name,
	}
}
