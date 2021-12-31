package service

import "time"

type Potato struct {
	Name           string    `json:"name"`
	AddedAt        time.Time `json:"added_at"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}

func newPotato(raw raw) Potato {
	return Potato{
		Name:           raw.Name,
		AddedAt:        raw.AddedAt,
		LastModifiedAt: raw.LastModifiedAt,
	}
}
