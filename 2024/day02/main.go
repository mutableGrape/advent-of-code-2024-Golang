package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
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

func validDiff(report []int) bool {
	first := report[0]
	for _, n := range report {
		n_abs := math.Abs(float64(n))
		if (n_abs > 3) || (n_abs < 1) || (first*n < 0) {
			return false
		}
	}
	return true
}

func validReport(report []int) bool {
	diff := []int{}

	last := report[0]
	for _, n := range report[1:] {
		diff = append(diff, n-last)
		last = n
	}

	return validDiff(diff)
}

func stringsToInts(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = n
	}
	return ints, nil
}

func solve(input string) (int, int, error) {
	// Part 1 Solution

	partA := 0
	partB := 0
	for _, l := range strings.Split(input, "\n") {
		lineStrs := strings.Split(strings.TrimSpace(l), " ")
		report, err := stringsToInts(lineStrs)

		if err != nil {
			return 0, 0, err
		}

		if validReport(report) {
			partA++
		} else {
			// Part 2 Solution
			for i := 0; i < len(report); i++ {
				newReport := make([]int, i)
				copy(newReport, report[:i])
				newReport = append(newReport, report[i+1:]...)
				if validReport(newReport) {
					partB++
					break
				}
			}
		}
	}

	return partA, partA + partB, nil
}
