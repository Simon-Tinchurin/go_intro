package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
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

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func Xor(a, b bool) bool {
	if a != b {
		return true
	} else {
		return false
	}
}

func PowersOfTwo(n int) []uint64 {
	var result []uint64
	for i := 0; i <= n; i++ {
		val := math.Pow(2.0, float64(i))
		result = append(result, uint64(val))
	}
	return result
}

func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}

func ReverseSeq(n int) (result []int) {
	// var result []int
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return result
}

func SquareSum(numbers []int) (result int) {
	for _, val := range numbers {
		result += val * val
	}
	return result
}

func HighAndLow(in string) (result string) {
	var lowest, highest int
	splitted := strings.Split(in, " ")

	for ind, char := range splitted {
		charToStr := string(char)
		charToInt, _ := strconv.Atoi(charToStr)
		if ind == 0 {
			lowest = charToInt
			highest = charToInt
		}
		if charToInt < lowest {
			lowest = charToInt
		}
		if charToInt > highest {
			highest = charToInt
		}
	}
	result = fmt.Sprint(highest) + " " + fmt.Sprint(lowest)
	return result
}

func PrinterError(s string) (res string) {
	var errors int
	totalLength := len(s)
	for _, char := range s {
		if char > 109 {
			errors++
		}
	}
	return fmt.Sprint(errors) + "/" + fmt.Sprint(totalLength)
}

func TwoToOne(s1 string, s2 string) (res string) {
	bytes := []byte(s1 + s2)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	sortedStr := string(bytes)

	for _, char := range sortedStr {
		if strings.Contains(res, string(char)) {
			continue
		} else {
			res += string(char)
		}
	}
	return res
}

func Maps(x []int) (res []int) {
	for _, val := range x {
		res = append(res, val+val)
	}
	return res
}

func EvenOrOdd(number int) string {
	remainder := number % 2
	if remainder == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func GetCount(str string) (count int) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, char := range str {
		for _, vowel := range vowels {
			if char == vowel {
				count++
			}
		}
	}
	return count
}

func PositiveSum(numbers []int) (res int) {
	for _, val := range numbers {
		if val > 0 {
			res += val
		}
	}
	return res
}

func OtherAngle(a int, b int) int {
	return 180 - a - b
}

func Century(year int) (cent int) {
	cent = year / 100
	if year%100 > 0 {
		cent++
	}
	return cent
}

func CountBy(x, n int) (res []int) {
	for i := 1; i <= n; i++ {
		res = append(res, x*i)
	}
	return res
}

func DNAStrand(dna string) (res string) {
	DNAMap := map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}
	for _, sym := range dna {
		res += DNAMap[string(sym)]
	}
	return res
}

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	return laLigaGoals + copaDelReyGoals + championsLeagueGoals
}

func Grow(arr []int) int {
	res := 1
	for _, val := range arr {
		res *= val
	}
	return res
}

func Solution(str, ending string) bool {
	if str == "" && ending != "" {
		return false
	}
	for i := 0; i < len(ending); i++ {
		if ending[len(ending)-i-1] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	// res := Century(1704)
	// res := CountBy(2, 5)
	// res := DNAStrand("ATTGC")
	// res := CheckForFactor(10, 2)

	res := Solution("", "t")

	fmt.Println(res)

	// result := SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7})
	// result := Invert([]int{1, 2, 3, 4})
	// result := Digitize(43512)
	// result := Summation(8)
	// result := RepeatStr(3, "Hello")
	// result := NoSpace("8 j 8   mBliB8g  imjB8B8  jl  B")
	// result := Feast("great blue heron", "garlic naan")
	// result := Greet("Jack")
	// result := PowersOfTwo(5)
	// result := ReverseSeq(5)
	// result := SquareSum([]int{1, 2, 2})
	// result := HighAndLow("1 2 3")
	// res := PrinterError("aaaxbbbbyyhwawiwjjjwwm")
	// res := TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq")
	// res := EvenOrOdd(20)
	// res := GetCount("abracadabra")
}
