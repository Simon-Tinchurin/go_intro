package main

import "fmt"

// variables and constants naming styles:
// name of the var should be in only one word - userName (standart style)
// without '-', not starting with numbers
// if first letter of the var name is capital:
// this var can be accessed outside of the package
// if first letter of the var name is lowercase:
// this var can be accessed only in package where it was declared

func types_2_main() {

	// it's possible to declare few variables in one line
	j, k, l := "shark", 2.05, 15
	fmt.Println(j, k, l)

	var text = "Hello"
	text2 := "World"
	i := 1
	var h float32
	h = 10.3
	res := float32(i) + h

	res2 := fmt.Sprintf("%s - %d - %s", text, i, text2)
	types_var := fmt.Sprintf("%T - %T - %T", text, i, text2)
	fmt.Println(res, res2, "\n", types_var)
	// fmt.Println(text, text2)

	// get user input
	var input float64
	fmt.Scanf("%f", &input)

	output := input + 1
	fmt.Println(output, "\n")
}
