package main

import "fmt"

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	name       string  = "Foo"
	intVar32   int32   = 10
	intVar64   int64   = 513434
	intVar     int     = -2
	// unsigned integer, always >=0
	uintVar   uint   = 1
	uintVar32 uint32 = 444
	byteVar   byte   = 0x1
	// for saving symbols in unicode, always in single quotes
	runVar rune = 'a'
)

// creating custom data type
// struct type, used to group together variables
// of different data types under a single name.
type Player struct {
	name   string
	health int
	power  float64
}

// Method with a value receiver, also there are a pointer receivers.
func (player Player) getHealth() int {
	return player.health
}

// passing the parameter to the func like:
// name_of_the_variable type_of_the_variable

// func getHealth(player Player) int {
// 	return player.health
// }

func main() {
	// empty struct
	// player := Player{}
	// assign values for variables in the struct
	player := Player{
		name:   "Jack",
		health: 50,
		power:  100.10,
	}
	// print formated
	// use +v to print names of the variables
	// %d - for integers

	// fmt.Printf("Health: %d\n", getHealth(player))

	// get variable of the struct with function receiver
	// fmt.Printf("Health: %d\n", player.getHealth())
	// or directly from struct
	fmt.Printf("Health: %d\n", player.health)

	// maps
	// empty map - map[type]type
	// most used method
	users := map[string]int{
		"user1": 20,
	}

	// second way to initialize the map
	// another_map := make(map[string]int)

	// add to map
	users["user2"] = 30

	// delete from the map - delete(map_name, "var_name")
	delete(users, "user1")

	// range over the map
	for k, v := range users {
		fmt.Printf("Key - %s and value - %d\n", k, v)
	}

	// access a map variable
	age := users["user2"]
	// ok - exists - type bool
	age2, ok := users["Jack"]
	if !ok {
		fmt.Println("Jack is not in the map")
	} else {
		fmt.Println("Jack is in the map: ", age2)
	}

	fmt.Printf("age: %d\n", age)

	fmt.Printf("Map: %+v\n", users)
}
