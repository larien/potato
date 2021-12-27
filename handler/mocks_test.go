package handler

import (
	"github.com/larien/potato/service"
	"github.com/larien/potato/utils/request/params"
)

type mockPotatoes struct {
	fnGet    func(id string) service.Potato
	fnList   func(params params.QueryParams) ([]service.Potato, error)
	fnCreate func(potato service.Potato) error
	fnUpdate func(potato service.Potato) error
	fnDelete func(id string) error
}

func (m mockPotatoes) Get(id string) service.Potato {
	if m.fnGet == nil {
		return service.Potato{}
	}
	return m.fnGet(id)
}

func (m mockPotatoes) List(params params.QueryParams) ([]service.Potato, error) {
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
