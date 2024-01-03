package main

import (
	"fmt"
	"sort"
)

func sort_main() {

	// ----------- INT
	sliceFloat64 := []float64{4.2, -1.3, 0.8, -3.8, 12.6}
	// -------- check if slice is sorted
	fmt.Println(sort.Float64sAreSorted(sliceFloat64))
	// sort the slice
	sort.Float64s(sliceFloat64)
	fmt.Println(sliceFloat64, sort.Float64sAreSorted(sliceFloat64))
	// search for item in slice,
	// returns index of this item in the slice
	// or if the value is not found in the slice,
	// it returns the index where the value would be inserted.
	fmt.Println(sort.SearchFloat64s(sliceFloat64, 0.9))
	// ----------- STRING
	sliceOfStr := []string{"one", "two", "three"}
	sort.Strings(sliceOfStr)
	fmt.Println(sliceOfStr, sort.StringsAreSorted(sliceOfStr))

	peopleInfo := []struct {
		Name string
		Age  int
	}{
		{"Be", 12},
		{"Foo", 13},
		{"Bar", 14},
		{"Text", 15},
	}
	// sort by name
	sort.Slice(peopleInfo, func(i, j int) bool {
		return peopleInfo[i].Name < peopleInfo[j].Name
	})
	fmt.Println(peopleInfo)
	// sort by age
	sort.SliceStable(peopleInfo, func(i, j int) bool {
		return peopleInfo[i].Age < peopleInfo[j].Age
	})
	fmt.Println(peopleInfo)
}
