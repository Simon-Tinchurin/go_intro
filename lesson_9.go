package main

import "fmt"

type Position struct {
	x int
	y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

// struct embedding
type SpecialEntity struct {
	Entity
	specialField float64
}

type Color int

// fmt.Stringer
func (c Color) String() string {
	switch c {
	case ColorBlue:
		return "BLUE"
	case ColorBlack:
		return "BLACK"
	case ColorYellow:
		return "YELLOW"
	case ColorPink:
		return "PINK"
	default:
		panic("invalid color given")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func les_9_main() {
	e := SpecialEntity{
		specialField: 1.234,
		Entity: Entity{
			name:    "my entity",
			id:      "123",
			version: "V.1.0",
			Position: Position{
				x: 100,
				y: 200,
			},
		},
	}
	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", e.name)
	fmt.Printf("%+v\n", e.x)

	fmt.Println("the color is:", ColorBlue)
}
