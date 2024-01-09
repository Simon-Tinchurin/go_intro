package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// TEST TASKS FROM CODEWARS

func SumMix(arr []any) int {
	var result int
	for _, v := range arr {
		switch value := v.(type) {
		case int:
			result += value
		case string:
			intValue, _ := strconv.Atoi(value)
			result += intValue
		}
	}
	return result
}

func Invert(arr []int) []int {
	var result []int
	for _, val := range arr {
		result = append(result, -val)
	}
	return result
}

func Digitize(n int) []int {
	var result []int
	lengthOfN := len(strconv.Itoa(n))
	for i := 0; i < lengthOfN; i++ {
		remainder := n % 10
		result = append(result, remainder)
		n = n / 10
	}
	return result
}

func Summation(n int) int {
	var result int
	for i := 0; i < n+1; i++ {
		result += i
	}
	return result
}

func RepeatStr(repetitions int, value string) string {
	var result string
	for i := 0; i < repetitions; i++ {
		result += value
	}
	return result
}

func NoSpace(word string) string {
	var result string
	for _, char := range word {
		if unicode.IsSpace(char) {
			continue
		} else {
			result += string(char)
		}
	}
	return result
}

func Feast(beast string, dish string) bool {
	var result bool
	if beast[0] == dish[0] && beast[len(beast)-1] == dish[len(dish)-1] {
		result = true
	} else {
		result = false
	}
	return result
}

func main() {
	// result := SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7})
	// result := Invert([]int{1, 2, 3, 4})
	// result := Digitize(43512)
	// result := Summation(8)
	// result := RepeatStr(3, "Hello")
	// result := NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B")
	result := Feast("great blue heron", "garlic naan")

	fmt.Println(result)
}
