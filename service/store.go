package service

import (
	"errors"

	"github.com/larien/potato/utils/drivers/kvs"
	"github.com/larien/potato/utils/request/params"
)

var (
	errNotFound      = errors.New("not found")
	errAlreadyExists = errors.New("already exists")
)

type store interface {
	list(params params.Queries) (raws, error)
	getByIDs(ids []string) (raws, error)
	create(raw raw) error
	update(raw raw) error
	delete(id string) error
}

type potatoStore struct {
	kvs *kvs.KeyValueStore
}

func newStore() store {
	return &potatoStore{
		kvs: kvs.New(),
	}
}

func (s *potatoStore) list(params params.Queries) (raws, error) {
	results := s.kvs.GetAll()

	raws := make(raws, len(results))
	for _, result := range results {
		r := result.(raw)
		if !r.Active {
			continue
		}
		raws[r.Name] = r
	}

	return raws, nil
}

func (s *potatoStore) getByIDs(ids []string) (raws, error) {
	raws := make(raws, len(ids))

	for _, id := range ids {
		r, ok := s.kvs.Get(id).(raw)
		if !ok && r.Name == "" && !r.Active {
			return nil, errNotFound
		}
		raws[r.Name] = r
	}

	return raws, nil
}

func (s *potatoStore) create(r raw) error {
	result, ok := s.kvs.Get(r.Name).(raw)
	if ok && result.Name != "" && result.Active {
		return errAlreadyExists
	}

	r.Active = true

	s.kvs.Set(r.Name, r)

	return nil
}

func (s *potatoStore) update(r raw) error {
	result := s.kvs.Get(r.Name).(raw)
	if result.Name == "" || !result.Active {
		return errNotFound
	}

	s.kvs.Set(r.Name, r)

	return nil
}

func (s *potatoStore) delete(id string) error {
	result := s.kvs.Get(id).(raw)
	if result.Name == "" || !result.Active {
		return errNotFound
	}

	result.Active = false
	s.kvs.Set(id, result)

	return nil
}
