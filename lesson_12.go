package main

import "fmt"

// pointers
// points to the memory adress / slot in memory

type Player1 struct {
	HP int
}

// if player is not a pointer, we are adjusting
// the copy of the player not the actual player
func TakeDamage(palyer *Player1, amount int) {
	palyer.HP -= amount
	fmt.Println("player is taking damage. New HP ->", palyer.HP)

}

// function receiver
func (p *Player1) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("player is taking damage. New HP ->", p.HP)
}

type Database struct {
	user string
}

type Server1 struct {
	db *Database // uintptr -> 8 bytes long
}

func (s *Server1) GetUserFromDB() string {
	// golang going to "dereference" the db pointer
	// its going to lookup the memory address of the pointer
	if s.db == nil {
		panic("database is not initialized")
	}
	return s.db.user
}

func main() {
	player := Player1{
		HP: 100,
	}
	TakeDamage(&player, 10)
	player.TakeDamage(10)
	fmt.Printf("%+v\n", player)

	s := &Server1{db: &Database{}}
	s.GetUserFromDB()
}
