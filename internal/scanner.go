package internal

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	Source  []byte
	Tokens  []Token
	Start   int
	Current int
	Line    uint16
}

func NewScanner(source_code []byte) Scanner {
	return Scanner{
		Source: source_code,
		Line:   1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for s.Current < len(s.Source) {
		s.scanToken()
		s.Start = s.Current
	}
	s.Tokens = append(s.Tokens, NewToken(EOF, "EOF", nil, s.Line))
	return s.Tokens
}

func (s *Scanner) scanToken() {
	c := s.Source[s.Current]
	s.Current++
	switch c {
	case '(':
		s.AddToken(LEFT_PAREN, nil)
	case ')':
		s.AddToken(RIGHT_PAREN, nil)
	case '{':
		s.AddToken(LEFT_BRACE, nil)
	case '}':
		s.AddToken(RIGHT_BRACE, nil)
	case '*':
		s.AddToken(STAR, nil)
	case '+':
		s.AddToken(PLUS, nil)
	case '-':
		s.AddToken(MINUS, nil)
	case ';':
		s.AddToken(SEMICOLON, nil)
	case '.':
		s.AddToken(DOT, nil)
	case ' ':
		{
		}
	case '\n':
		{
		}
	case '<':
		if s.Source[s.Current] == '=' {
			s.Current++
			s.AddToken(LESS_EQUAL, nil)
		} else {
			s.AddToken(LESS, nil)
		}
	case '>':
		if s.Source[s.Current] == '=' {
			s.Current++
			s.AddToken(GREATER_EQUAL, nil)
		} else {
			s.AddToken(GREATER, nil)
		}
	case '=':
		if s.Source[s.Current] == '=' {
			s.Current++
			s.AddToken(EQUAL_EQUAL, nil)
		} else {
			s.AddToken(EQUAL, nil)
		}
	case '!':
		if s.Source[s.Current] == '=' {
			s.Current++
			s.AddToken(BANG_EQUAL, nil)
		} else {
			s.AddToken(BANG, nil)
		}
	case '/':
		if s.Source[s.Current] == '/' {
			for s.Current < len(s.Source) && s.Source[s.Current] != '\n' {
				s.Current++
			}
		} else {
			s.AddToken(SLASH, nil)
		}
	case '"':
		s.ProcessString()
	default:
		if isAlpha(c) {
			s.ProcessIdentifier()
		} else if isDigit(c) {
			s.ProcessNumber()
		} else {
			panic(fmt.Sprintf("Found unexpected character: %v", string(c)))
		}
	}
}

func (s *Scanner) ProcessString() {
	for s.Current < len(s.Source) && s.Source[s.Current] != '\n' {
		if s.Source[s.Current] == '"' {
			s.Current++
			text := string(s.Source[s.Start+1 : s.Current-1])
			s.Tokens = append(s.Tokens, NewToken(STRING, text, text, s.Line))
      return
		}
		s.Current++
	}
  panic("Expected closing '\"' for a string")
}

func (s *Scanner) ProcessNumber() {
	for s.Current < len(s.Source) && isDigit(s.Source[s.Current]) {
		s.Current++
	}

	if s.Source[s.Current] == '.' &&
		s.Current+1 < len(s.Source) && isDigit(s.Source[s.Current+1]) {
		s.Current++
		for s.Current < len(s.Source) && isDigit(s.Source[s.Current]) {
			s.Current++
		}
	}

	numStr := string(s.Source[s.Start:s.Current])
	literal, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Error parsing number:", numStr)
		return
	}
	s.Tokens = append(s.Tokens, NewToken(NUMBER, numStr, literal, s.Line))
}

func (s *Scanner) ProcessIdentifier() {
	for isAlphanumeric(s.Source[s.Current]) {
		s.Current++
	}

	str := string(s.Source[s.Start:s.Current])
	keyword, ok := keywords[str]
	if ok {
		s.Tokens = append(s.Tokens, NewToken(keyword, str, nil, s.Line))
	} else {
		s.Tokens = append(s.Tokens, NewToken(IDENTIFIER, str, str, s.Line))
	}
}

// helper
func (s *Scanner) AddToken(tokenType TokenType, literal any) {
	var text string
	if literal == nil {
		text = ""
	} else {
		text = string(s.Source[s.Start:s.Current])
	}
	s.Tokens = append(s.Tokens, NewToken(tokenType, text, literal, s.Line))
}
