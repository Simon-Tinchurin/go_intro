package main

import "fmt"

// generics???
type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func foo[T any](val T) {
	fmt.Println(val)
}

func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func les_11_main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("foo", 1)
	m1.Insert("bar", 2)

	fmt.Println(m1)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 9.99)
	m2.Insert(2, 10.123)

	fmt.Println(m2)

	foo[int](1)

	intSlice := []int{1, 2, 3, 4, 5}
	stringSlice := []string{"apple", "banana", "cherry"}

	PrintSlice(intSlice)
	PrintSlice(stringSlice)
}
