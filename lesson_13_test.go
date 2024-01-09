package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayers(t *testing.T) {
	expected := Player2{
		name: "Mark",
		hp:   100,
	}
	have := Player2{
		name: "Alice",
		hp:   59,
	}
	// if you want to campare deep nested objects,
	// use -> reflect.DeepEqual(obj1, obj2)
	if !reflect.DeepEqual(expected, have) {
		t.Errorf("expected %+v but got %+v", expected, have)
	}
}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 4
		b        = 5
	)
	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but have %d", expected, have)
	}
	// fmt.Println("hello from test case")
}
