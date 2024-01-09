package main

// testing
// if we want to test everything in the directory:
// 1. "go mod init <project_name>"
// 2. "go test ./..."

// if you want to see prints/logs in test cases - "go test -v ./..."
// if you need to run one test -> "go test ./ -v -run TestFuncName"
// use "-count=1" to disable caching test results

func calculateValues(x int, y int) int {
	return x + y
}

type Player2 struct {
	name string
	hp   int
}

func les_13_main() {

}
