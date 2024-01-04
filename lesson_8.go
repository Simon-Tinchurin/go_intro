package main

import (
	"fmt"
	"go_intro_app/types"
	"go_intro_app/util"
)

// --program--
// 1. something that runs / executes code (go run somefile.go)
// => main() -> entry function of the program
// => package main

// --package / library / module--
// 2. something that someone can import into their program or library
// -> package <whatever>

// To build program
// 1. go mod init <program_or_module_name>
// 2. go build -o <binary_file_name>
// 3. to run binary - ./binary_file_name
// if something changed - you need to rebuild
// the binary file with the same command

// User -> public access
// user -> private access BUT public in its own package

func les_8_main() {
	// number := getNum ber()
	user := types.User{
		Username: util.GetUserName(),
		Age:      util.GetAge(),
	}

	fmt.Printf("The user is %+v \n", user)
}
