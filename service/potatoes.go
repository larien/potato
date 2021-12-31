package service

import (
	"errors"
	"log"
	"time"

	"github.com/larien/potato/utils/request/params"
)

var (
	ErrAlreadyExists = errors.New("potato already exists")
	ErrNotFound      = errors.New("potato not found")
)

type Potatoes interface {
	List(params params.Queries) ([]Potato, error)
	Get(id string) (Potato, error)
	Create(potato Potato) error
	Update(potato Potato) error
	Delete(id string) error
}

type potatoes struct {
	store store
}

func New() Potatoes {
	return potatoes{
		store: newStore(),
	}
}

func (p potatoes) List(params params.Queries) ([]Potato, error) {
	var result []Potato
	raws, err := p.store.list(params)
	if err != nil {
		return nil, err
	}
	for _, raw := range raws {
		result = append(result, newPotato(raw))
	}
	return result, nil
}

func (p potatoes) Get(id string) (Potato, error) {
	raws, err := p.store.getByIDs([]string{id})
	if err != nil {
		return Potato{}, err
	}
	return newPotato(raws[id]), nil
}

func (p potatoes) Create(potato Potato) error {
	raws, err := p.store.list(params.Queries{})
	if err != nil {
		return err
	}

	if _, ok := raws[potato.Name]; ok {
		return ErrAlreadyExists
	}

	potato.AddedAt = time.Now()
	potato.LastModifiedAt = potato.AddedAt

	p.store.create(newRaw(potato))
	log.Println("Created potato:", potato.Name)

	return nil
}

func (p potatoes) Update(potato Potato) error {
	raws, err := p.store.getByIDs([]string{potato.Name})
	if err != nil {
		return err
	}

	raw, ok := raws[potato.Name]
	if !ok {
		return ErrNotFound
	}

	oldPotato := newPotato(raw)
	potato.AddedAt = oldPotato.AddedAt
	potato.LastModifiedAt = time.Now()

	if err := p.store.update(newRaw(potato)); err != nil {
		return err
	}
	log.Println("Updated potato:", potato.Name)

	return nil
}

func (p potatoes) Delete(id string) error {
	raws, err := p.store.getByIDs([]string{id})
	if err != nil {
		return err
	}

	_, ok := raws[id]
	if !ok {
		return ErrNotFound
	}

	if err := p.store.delete(id); err != nil {
		return err
	}
	log.Println("Deleted potato:", id)

	return nil
}
