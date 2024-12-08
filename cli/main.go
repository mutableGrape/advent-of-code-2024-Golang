package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	day := createCmd.Int("day", 1, "Day of the challenge")

	if len(os.Args) < 2 {
		fmt.Println("expected 'create' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		createDay(*day)
	default:
		fmt.Println("expected 'create' subcommand")
		os.Exit(1)
	}
}

func createDay(day int) {
	dayStr := fmt.Sprintf("day%02d", day)
	baseDir := filepath.Join("2024", dayStr)

	if err := os.MkdirAll(baseDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	mainFile := filepath.Join(baseDir, "main.go")
	inputFile := filepath.Join(baseDir, "input.txt")
	readmeFile := filepath.Join(baseDir, "README.md")

	mainContent, err := os.ReadFile("cli/solution_template/template.go")
	if err != nil {
		fmt.Printf("Error reading template: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(mainFile, []byte(mainContent), 0644); err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(inputFile, []byte(""), 0644); err != nil {
		fmt.Printf("Error creating input.txt: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(readmeFile, []byte("# Day "+fmt.Sprintf("%d", day)), 0644); err != nil {
		fmt.Printf("Error creating README.md: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created files for day %d\n", day)
}
