package main

import (
	"fmt"
	"strings"
)

// 1. advanced interface mechanics
// 2. typed functions

type Putter interface {
	Put(id int, val any) error
}

type Storage interface {
	Get(id int) (any, error)
	Putter
}

type simpePutter struct{}

func (s *simpePutter) Put(id int, val any) error {
	return nil
}

type FooStorage struct{}

func (s *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (s *FooStorage) Put(id int, val any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, val any, p Putter) error {
	// store := &Storage{}
	return p.Put(id, val)
}

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func les_10_main() {
	s := &Server{
		store: &FooStorage{},
	}
	s.store.Get(1)
	s.store.Put(1, "foo")
	updateValue(1, "foo", s.store)
	sputter := &simpePutter{}
	updateValue(2, "boo", sputter)
	fmt.Println(transformString("hello!", Uppercase))
	fmt.Println(transformString("hello!", Prefixer("FOOO_")))
	fmt.Println(transformString("hello!", Prefixer("BAR-")))
}
