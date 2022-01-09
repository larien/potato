package service

import (
	"errors"
	"log"
	"time"

	"github.com/larien/potato/utils/request/params"
)

var (
	ErrNotFound      = errors.New("potato not found")
	ErrAlreadyExists = errors.New("potato already exists")
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

	if !raws[id].Active {
		return Potato{}, ErrNotFound
	}

	return newPotato(raws[id]), nil
}

func (p potatoes) Create(potato Potato) error {
	potato.AddedAt = time.Now()
	potato.LastModifiedAt = potato.AddedAt

	err := p.store.create(newRaw(potato))
	if err == nil {
		log.Println("Created potato:", potato.Name)
		return nil
	}

	if errors.Is(err, errAlreadyExists) {
		return ErrAlreadyExists
	}

	return err
}

func (p potatoes) Update(potato Potato) error {
	raws, err := p.store.getByIDs([]string{potato.Name})
	if err != nil {
		return err
	}

	oldPotato := newPotato(raws[potato.Name])
	potato.AddedAt = oldPotato.AddedAt
	potato.LastModifiedAt = time.Now()

	err = p.store.update(newRaw(potato))
	if err == nil {
		log.Println("Updated potato:", potato.Name)
		return nil
	}

	if errors.Is(err, errNotFound) {
		return ErrNotFound
	}

	return err
}

func (p potatoes) Delete(id string) error {
	err := p.store.delete(id)
	if err == nil {
		log.Println("Deleted potato:", id)
		return nil
	}

	if errors.Is(err, errNotFound) {
		return ErrNotFound
	}

	return err
}
