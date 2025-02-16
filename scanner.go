package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int
}

func NewScanner(source string) Scanner {
	return Scanner{
		Source:  source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (scanner *Scanner) ScanTokens() []Token {
	for !scanner.IsAtEnd() {
		scanner.Start = scanner.Current
		scanner.ScanToken()
	}
	scanner.Tokens = append(scanner.Tokens, Token{EOF, "", nil, scanner.Line})
	return scanner.Tokens
}

func (scanner *Scanner) IsAtEnd() bool {
	return scanner.Current >= len(scanner.Source)
}

func (scanner *Scanner) Peak() string {
	if scanner.IsAtEnd() {
		return "\\0"
	}
	return string(scanner.Source[scanner.Current])
}
func (scanner *Scanner) PeakNext() string {
	if scanner.Current+1 >= len(scanner.Source) {
		return "\\0"
	}
	return string(scanner.Source[scanner.Current+1])
}

func (scanner *Scanner) ScanToken() {
	c := scanner.Advance()
	switch c {
	case '(':
		scanner.AddToken(LEFT_BRACE, nil)
	case ')':
		scanner.AddToken(RIGHT_PAREN, nil)
	case '{':
		scanner.AddToken(LEFT_BRACE, nil)
	case '}':
		scanner.AddToken(RIGHT_BRACE, nil)
	case ',':
		scanner.AddToken(COMMA, nil)
	case '.':
		scanner.AddToken(DOT, nil)
	case '-':
		scanner.AddToken(MINUS, nil)
	case '+':
		scanner.AddToken(PLUS, nil)
	case ';':
		scanner.AddToken(SEMICOLON, nil)
	case '*':
		scanner.AddToken(STAR, nil)
	case '!':
		var tokenType TokenType
		if scanner.Match(byte('=')) {
			tokenType = BANG_EQUAL
		} else {
			tokenType = BANG
		}
		scanner.AddToken(tokenType, nil)
	case '=':
		var tokenType TokenType
		if scanner.Match(byte('=')) {
			tokenType = EQUAL_EQUAL
		} else {
			tokenType = EQUAL
		}
		scanner.AddToken(tokenType, nil)
	case '<':
		var tokenType TokenType
		if scanner.Match(byte('=')) {
			tokenType = LESS_EQUAL
		} else {
			tokenType = LESS
		}
		scanner.AddToken(tokenType, nil)
	case '>':
		var tokenType TokenType
		if scanner.Match(byte('=')) {
			tokenType = GREATER_EQUAL
		} else {
			tokenType = GREATER
		}
		scanner.AddToken(tokenType, nil)
	case '/':
		if scanner.Match('/') {
			for scanner.Peak() != "\\n" && !scanner.IsAtEnd() {
				// Advance until end of line or end of file.
				scanner.Advance()
			}
		} else {
			scanner.AddToken(SLASH, nil)
		}
	case '\\':
		if scanner.Match('r') || scanner.Match('t') {
			// Do nothgin
		} else if scanner.Match('n') {
			scanner.Line += 1
		} else {
			error(scanner.Line, "Unexpected character")
		}
	case ' ': // Do nothing
	case '"':
		scanner.String()
	default:
		if isDigit(c) {
			scanner.Number()
		} else if isAlpha(c) {
			scanner.Identifier()
		} else {
			error(scanner.Line, "Unexpected character")
		}
	}
}

func (scanner *Scanner) Advance() byte {
	c := scanner.Source[scanner.Current]
	scanner.Current = scanner.Current + 1
	return c
}

func (scanner *Scanner) AddToken(tokenType TokenType, literal any) {
	lexeme := scanner.Source[scanner.Start:scanner.Current]
	scanner.Tokens = append(scanner.Tokens, Token{tokenType, lexeme, literal, scanner.Line})
}

func (scanner *Scanner) Match(expected byte) bool {
	if scanner.IsAtEnd() || scanner.Source[scanner.Current] != expected {
		return false
	}
	scanner.Current += 1
	return true
}

func (scanner *Scanner) String() {
	for scanner.Peak() != "\"" && !scanner.IsAtEnd() {
		if scanner.Peak() == "\n" {
			scanner.Line += 1
		}
		scanner.Advance()
		// fmt.Println(string(scanner.Source[scanner.Start:scanner.Current]))
	}

	if scanner.IsAtEnd() {
		error(scanner.Line, "Unterminated String")
		return
	}
	// fmt.Println("Now we are at: ", string( scanner.Source[scanner.Current] ))
	scanner.Advance() // to skip the closing "
	value := scanner.Source[scanner.Start+1 : scanner.Current-1]
	// fmt.Printf("value::%s\n", value)
	scanner.AddToken(STRING, value)
}

func (scanner *Scanner) Number() {
	for isDigit(byte(scanner.Peak()[0])) {
		scanner.Advance()
	}
	if scanner.Peak() == "." && isDigit(byte(scanner.PeakNext()[0])) {
		scanner.Advance() // skip the '.'
		for isDigit(byte(scanner.Peak()[0])) {
			scanner.Advance()
		}
	}
	num_str := scanner.Source[scanner.Start:scanner.Current]
	// fmt.Printf("num_str::%s\n", num_str)
	num, err := strconv.ParseFloat(num_str, 32)
	if err != nil {
		error(scanner.Line, fmt.Sprintf("Internal error: %s", err))
	}
	scanner.AddToken(NUMBER, num)
}

func (scanner *Scanner) Identifier() {
	for isAlphaNumeric(byte(scanner.Peak()[0])) {
		scanner.Advance()
	}
	text := scanner.Source[scanner.Start:scanner.Current]
	tokenType, exists := keywords[text]
	// fmt.Printf("text::%s, tokenType::%s\n", text, tokenType.String())
	if !exists {
		tokenType = IDENTIFIER
	}
	scanner.AddToken(tokenType, nil)
}
