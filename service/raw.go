package service

import "time"

type (
	raws map[string]raw
	raw  struct {
		Name           string    `json:"name"`
		AddedAt        time.Time `json:"added_at"`
		LastModifiedAt time.Time `json:"last_modified_at"`
		Active         bool      `json:"active"`
	}
)

func newRaw(potato Potato) raw {
	return raw{
		Name:           potato.Name,
		AddedAt:        potato.AddedAt,
		LastModifiedAt: potato.LastModifiedAt,
		Active:         true,
	}
}
