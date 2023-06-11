package main

import (
	"fmt"
	"os"

	"lem-in/interface/CLI"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("ERROR: provide input file. Example: ./lem-in anthill.txt")
		return
	}

	cli := CLI.NewCLI()
	err := cli.Run(args[0])
	if err != nil {
		fmt.Println(err)
	}
}
