package main

import (
	"bufio"
	"fmt"
	"os"
)

// global config
var hadError bool = false

func runFile(file_path string) {
	file, err := os.ReadFile(file_path)
	if err != nil {
		fmt.Println("Failed to open file provided.")
		os.Exit(1)
	}
	run(string(file))
	if hadError {
		os.Exit(1)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			line := scanner.Text() // This only read line at a time?
			run(line)
			hadError = false
		} else {
			// Check for errors or EOF
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading input:", err)
				return
			}
			// If no error, we hit EOF (Ctrl+D)
			return
		}
	}
}

func run(source string) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token.toString())
	}
}

func error(line int, message string) {
	fmt.Printf("[line %d] Error: %s\n", line, message)
}

func main() {
  init_keywords()
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
