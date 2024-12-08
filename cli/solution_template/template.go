package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Define a flag for using example files
	example := flag.Bool("e", false, "Use example input file")
	flag.Parse()

	// Determine the input file based on the flag
	inputFile := "input.txt"
	if *example {
		inputFile = "example.txt"
	}

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.TrimSpace(string(data))

	partA, partB, err := solve(input)
	if err != nil {
		fmt.Println("error in part 1", err)
	} else {
		fmt.Println("Part 1:", partA)
		fmt.Println("Part 2:", partB)
	}
}

func solve(input string) (int, int, error) {
	// Part 1 Solution

	// Part 2 Solution

	return 0, 0, nil
}
