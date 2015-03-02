package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	args := os.Args[1:]
	fmt.Println(args)
	return 0
}
