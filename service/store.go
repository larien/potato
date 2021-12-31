package service

import (
	"encoding/json"
	"os"

	"github.com/larien/potato/utils/request/params"
)

// in a regular microservice, we'd probably implement a SQL/NoSQL driver. To make things simple here,
// let's update a file for now. Never use files as database in the real world, they are (among many disadvantages)
// not performatic, not thread safe and not reliable.

const (
	filename = "db/potatoes.json"
)

type store interface {
	list(params params.Queries) (raws, error)
	getByIDs(ids []string) (raws, error)
	create(raw raw) error
	update(raw raw) error
	delete(id string) error
}

type fileStore struct{}

func newStore() store {
	return fileStore{}
}

func (s fileStore) list(params params.Queries) (raws, error) {
	var result raws
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(content, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s fileStore) getByIDs(ids []string) (raws, error) {
	list, err := s.list(params.Queries{})
	if err != nil {
		return nil, err
	}
	result := make(raws)
	for _, raw := range list {
		for _, id := range ids {
			if raw.Name == id {
				result[id] = raw
			}
		}
	}
	return result, nil
}

func (s fileStore) create(raw raw) error {
	list, err := s.list(params.Queries{})
	if err != nil {
		return err
	}
	list[raw.Name] = raw

	content, err := json.Marshal(list)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(content), 0644)
}

func (s fileStore) update(raw raw) error {
	list, err := s.list(params.Queries{})
	if err != nil {
		return err
	}
	for i, potato := range list {
		if potato.Name == raw.Name {
			list[i] = raw
		}
	}
	content, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}

func (s fileStore) delete(id string) error {
	list, err := s.list(params.Queries{})
	if err != nil {
		return err
	}
	raw := list[id]
	raw.Active = false
	list[id] = raw
	content, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}
