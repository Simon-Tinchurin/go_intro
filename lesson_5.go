package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife

type WeaponType int

// const (
// 	Axe         WeaponType = 1
// 	Sword       WeaponType = 2
// 	WoodenStick WeaponType = 3
// 	Knife       WeaponType = 4
// )

// ???? some kind of magic like __repr__ in python happens here...
func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "AXE"
	case Sword:
		return "SWORD"
	case WoodenStick:
		return "WOODENSTICK"
	case Knife:
		return "KNIFE"
	}
	return ""
}

const (
	Axe WeaponType = iota // increment everything below
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType WeaponType) int {
	// control statement?
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		// like an exception?
		panic("weapon does not exist")
	}
}

func les_5_main() {
	fmt.Printf("Damage of the given weapon (%s) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("Damage of the given weapon (%s) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("Damage of the given weapon (%s) (%d):\n", WoodenStick, getDamage(WoodenStick))
	fmt.Printf("Damage of the given weapon (%s) (%d):\n", Knife, getDamage(Knife))
}
