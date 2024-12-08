package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
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

func linesToInts(input string) ([][2]int, error) {
	var res [][2]int
	for _, l := range strings.Split(input, "\n") {
		var r [2]int
		for x, i := range strings.Fields(l) {
			n, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			r[x] = n
		}
		res = append(res, r)
	}
	return res, nil
}

func solve(input string) (int, int, error) {
	// Part 1 Solution
	input_pairs, err := linesToInts(input)
	n := len(input_pairs)
	if err != nil {
		return 0, 0, err
	}

	// Unpack pairs into seperate lists
	a := make([]int, n)
	for i, j := range input_pairs {
		a[i] = j[0]
	}

	b := make([]int, n)
	for i, j := range input_pairs {
		b[i] = j[1]
	}

	// Sort the lists
	slices.Sort(a)
	slices.Sort(b)

	// Sum the absolute differences
	sumA := 0
	for i := 0; i < n; i++ {
		sumA += (int)(math.Abs((float64)(b[i] - a[i])))
	}

	// Part 2 Solution

	left, right_index, counter, sumB := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		if a[i] == left {
			continue
		}

		sumB += left * counter
		left = a[i]
		counter = 0

		for (right_index < n) && (b[right_index] < left) {
			right_index++
		}

		for (right_index < n) && (b[right_index] == left) {
			counter++
			right_index++
		}

	}
	sumB += left * counter

	return sumA, sumB, nil
}
