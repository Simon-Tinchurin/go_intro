package main

import (
	"encoding/json"
	"fmt"
)

// defining structure
type status struct {
	Name string
}

type info struct {
	Friends []struct {
		Id   int
		Name string
	}
}

func main() {
	text := `{"friends": [{"id": 1,"name": "Annabelle Carey"},
						{"id": 2,"name": "Nellie Hansen"},
						{"id": 3,"name": "Milagros Nash"}]}`

	var st = status{Name: "status_200"}
	fmt.Println(st.Name)

	var infos info
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos)
	fmt.Println(infos.Friends)
	for i := range infos.Friends {
		fmt.Println(i, infos.Friends[i].Id, infos.Friends[i].Name)
	}
}
