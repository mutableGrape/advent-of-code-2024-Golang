package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.TrimSpace(string(data))
	partA, err := part1(input)
	if err != nil {
		fmt.Println("error in part 1", err)
	} else {
		fmt.Println("Part 1:", partA)
	}

	partB, err := part2(input)
	if err != nil {
		fmt.Println("error in part 2", err)
	} else {
		fmt.Println("Part 2:", partB)
	}
}

func part1(input string) (int, error) {
	// Part 1 Solution
	return 0, nil
}

func part2(input string) (int, error) {
	// Part 2 Solution
	return 0, nil
}
