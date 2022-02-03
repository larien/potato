package handler

import (
	"github.com/larien/potato/internal/request/params"
	"github.com/larien/potato/service"
)

type mockPotatoes struct {
	fnGet    func(id string) (service.Potato, error)
	fnList   func(params params.Queries) ([]service.Potato, error)
	fnCreate func(potato service.Potato) error
	fnUpdate func(potato service.Potato) error
	fnDelete func(id string) error
}

func (m mockPotatoes) Get(id string) (service.Potato, error) {
	if m.fnGet == nil {
		return service.Potato{}, nil
	}
	return m.fnGet(id)
}

func (m mockPotatoes) List(params params.Queries) ([]service.Potato, error) {
	if m.fnList == nil {
		return []service.Potato{}, nil
	}
	return m.fnList(params)
}

func (m mockPotatoes) Create(potato service.Potato) error {
	if m.fnCreate == nil {
		return nil
	}
	return m.fnCreate(potato)
}

func (m mockPotatoes) Update(potato service.Potato) error {
	if m.fnUpdate == nil {
		return nil
	}
	return m.fnUpdate(potato)
}

func (m mockPotatoes) Delete(id string) error {
	if m.fnDelete == nil {
		return nil
	}
	return m.fnDelete(id)
}
