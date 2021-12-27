package handler

import "github.com/larien/potato/business"

type V1Potato struct {
	Name string `json:"name"`
}

func (p V1Potato) toPotato() business.Potato {
	return business.Potato{
		Name: p.Name,
	}
}
