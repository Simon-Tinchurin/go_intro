package main

import "fmt"

// interfaces

// Storage
// Storer
// io.Read(er)

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(number int) error
}

type PostgresNumberStore struct {
	// postgres db connection and stuff
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7, 8}, nil
}

func (p PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the Postgres storage")
	return nil
}

type MongoDBNumberStore struct {
	// some values
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the storage")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: PostgresNumberStore{},
	}

	// short way to deal with errors
	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	// long way to deal with errors
	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(numbers, err)
}
