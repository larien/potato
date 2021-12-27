package handler

import (
	"github.com/larien/potato/business"
	"github.com/larien/potato/utils/request/params"
)

type mockPotatoes struct {
	fnGet    func(id string) business.Potato
	fnList   func(params params.QueryParams) ([]business.Potato, error)
	fnCreate func(potato business.Potato) error
	fnUpdate func(potato business.Potato) error
	fnDelete func(id string) error
}

func (m mockPotatoes) Get(id string) business.Potato {
	if m.fnGet == nil {
		return business.Potato{}
	}
	return m.fnGet(id)
}

func (m mockPotatoes) List(params params.QueryParams) ([]business.Potato, error) {
	if m.fnList == nil {
		return []business.Potato{}, nil
	}
	return m.fnList(params)
}

func (m mockPotatoes) Create(potato business.Potato) error {
	if m.fnCreate == nil {
		return nil
	}
	return m.fnCreate(potato)
}

func (m mockPotatoes) Update(potato business.Potato) error {
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
