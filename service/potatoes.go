package service

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/larien/potato/utils/request/params"
)

var (
	list map[string]Potato
	lock sync.RWMutex

	ErrAlreadyExists = errors.New("potato already exists")
	ErrNotFound      = errors.New("potato not found")
)

type Potatoes interface {
	List(params params.QueryParams) ([]Potato, error)
	Get(id string) Potato
	Create(potato Potato) error
	Update(potato Potato) error
	Delete(id string) error
}

type potatoes struct{}

func New() Potatoes {
	if list == nil {
		list = make(map[string]Potato)
	}
	return potatoes{}
}

func (p potatoes) List(params params.QueryParams) ([]Potato, error) {
	var result []Potato
	for _, potato := range list {
		result = append(result, potato)
	}
	return result, nil
}

func (p potatoes) Get(id string) Potato {
	potato, ok := list[id]
	if !ok {
		return Potato{}
	}
	return potato
}

func (p potatoes) Create(potato Potato) error {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := list[potato.Name]; ok {
		return ErrAlreadyExists
	}

	potato.AddedAt = time.Now()
	potato.LastModifiedAt = potato.AddedAt
	list[potato.Name] = potato
	log.Println("Created potato:", potato.Name)

	return nil
}

func (p potatoes) Update(potato Potato) error {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := list[potato.Name]; !ok {
		return ErrNotFound
	}

	potato.LastModifiedAt = time.Now()
	list[potato.Name] = potato
	log.Println("Updated potato:", potato.Name)

	return nil
}

func (p potatoes) Delete(id string) error {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := list[id]; !ok {
		return ErrNotFound
	}

	delete(list, id)
	log.Println("Deleted potato:", id)

	return nil
}
