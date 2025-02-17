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
	run(file)
	if hadError {
		os.Exit(1)
	}
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if scanner.Scan() {
			line := scanner.Bytes() // This only read line at a time?
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

func run(source []byte) {
	scanner := NewScanner(source)
	tokens := scanner.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token.toString())
	}

  parser := NewParser(tokens)
  expressions := parser.Parse()
  if hadError {
    return
  }
  astPrinter := AstPrinter{}
  fmt.Println(astPrinter.Print(expressions))
  astPrinter.Print(expressions)
}

func error(token Token, message string) {
  if token.Type == EOF {
	fmt.Printf("[line %d] Error at end %s\n", token.Line, message)
  }
  fmt.Printf("[line %d] Error at '%s': %s\n", token.Line, token.Lexeme, message)
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
