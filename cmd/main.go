package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: provide input file. Example: ./lem-in anthill.txt")
		return
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("ERROR: file not found")
		return
	}
	fmt.Println(content)
	// 1) read content build ant farm
	// 2) push anthill into solving algorithm
	// 3) printout result
}
