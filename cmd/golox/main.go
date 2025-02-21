package main

import (
	"fmt"
	"os"
  gx "golox/internal"
)

func runFile(file_path string) {
  source_code, err := os.ReadFile(file_path)
  if err != nil {
    panic(fmt.Sprintf("Error opening file: %v\n", err))
  }
  run(source_code)
}

func run(source_code []byte) {
  scanner := gx.NewScanner(source_code)
  tokens := scanner.ScanTokens()
  fmt.Printf("Tokens from scanner: %+v\n", tokens)

  parser := gx.NewParser(tokens)
  parsedExpression := parser.Parse()
  fmt.Printf("parsedExpression: %s\n", parsedExpression) 

  interpreter := gx.NewInterpreter()
  interpreter.Interpret(parsedExpression)
}

func main() {
  if len(os.Args) != 2 {
    panic("Usage: golox <script_path.gx>")
  } else {
    runFile(os.Args[1])
  }
}
