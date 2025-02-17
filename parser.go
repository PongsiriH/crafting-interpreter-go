package main

import "fmt"

type ParseResult struct {
	Expression Expr
	HadError   bool
}

func ParseSuccess(expression Expr) ParseResult {
	return ParseResult{
		Expression: expression,
		HadError:   false,
	}
}

func ParseFailed() ParseResult {
	return ParseResult{
		HadError: true,
	}
}

type Parser struct {
	Tokens  []Token
	Current int
}

func NewParser(tokens []Token) Parser {
	return Parser{
		Tokens:  tokens,
		Current: 0,
	}
}

func (parser *Parser) Match(tokenTypes []TokenType) bool {
	// If match the tokenTypes, advance and return true.
	for _, tokenType := range tokenTypes {
		if parser.Check(tokenType) {
			parser.Advance()
			return true
		}
	}
	return false
}

func (parser *Parser) Check(tokenType TokenType) bool {
	fmt.Println("Check", parser.IsAtEnd(), parser.Peek())
	if parser.IsAtEnd() {
		return false
	}
	return parser.Peek().Type == tokenType
}

func (parser *Parser) Advance() Token {
	if !parser.IsAtEnd() {
		parser.Current += 1
	}
	return parser.Previous()
}

func (parser *Parser) IsAtEnd() bool {
	return parser.Peek().Type == EOF
}

func (parser *Parser) Peek() Token {
	return parser.Tokens[parser.Current]
}

func (parser *Parser) Previous() Token {
	return parser.Tokens[parser.Current-1]
}

// // Start
func (parser *Parser) Expression() Expr {
	return parser.Equality()
}

func (parser *Parser) Equality() Expr {
	expr := parser.Comparison()
	for parser.Match([]TokenType{BANG_EQUAL, EQUAL_EQUAL}) {
		operator := parser.Previous()
		right := parser.Comparison()
		expr = &BinaryExpr{expr, operator, right}
	}
	return expr
}

func (parser *Parser) Comparison() Expr {
	expr := parser.Term()
	for parser.Match([]TokenType{GREATER, GREATER_EQUAL, LESS, LESS_EQUAL}) {
		operator := parser.Previous()
		right := parser.Term()
		expr = &BinaryExpr{expr, operator, right}
	}
	return expr
}

func (parser *Parser) Term() Expr {
	expr := parser.Factor()
	for parser.Match([]TokenType{MINUS, PLUS}) {
		operator := parser.Previous()
		right := parser.Factor()
		expr = &BinaryExpr{expr, operator, right}
	}
	return expr
}

func (parser *Parser) Factor() Expr {
	expr := parser.Unary()
	for parser.Match([]TokenType{SLASH, STAR}) {
		operator := parser.Previous()
		right := parser.Unary()
		expr = &BinaryExpr{expr, operator, right}
	}
	return expr
}

func (parser *Parser) Unary() Expr {
	for parser.Match([]TokenType{}) {
		operator := parser.Previous()
		right := parser.Unary()
		return &UnaryExpr{operator, right}
	}
	parseResult := parser.Primary()
	if parseResult.HadError {
		fmt.Println("Failed to parse: ")
	}
	return parseResult.Expression
}

func (parser *Parser) Primary() ParseResult {
	fmt.Println("bye? ", parser.Tokens, parser.Current, parser.Peek())
	if parser.Match([]TokenType{FALSE}) {
		return ParseSuccess(&LiteralExpr{BoolLiteral(FALSE)})
	}
	if parser.Match([]TokenType{TRUE}) {
		return ParseSuccess(&LiteralExpr{Value: BoolLiteral(TRUE)})
	}
	if parser.Match([]TokenType{NIL}) {
		return ParseSuccess(&LiteralExpr{Value: nil})
	}
	if parser.Match([]TokenType{NUMBER}) {
		return ParseSuccess(&LiteralExpr{NumberLiteral(parser.Previous().Literal.(float64))})
	}
	if parser.Match([]TokenType{STRING}) {
		return ParseSuccess(&LiteralExpr{StringLiteral(parser.Previous().Literal.([]byte))})
	}
	if parser.Match([]TokenType{IDENTIFIER}) {
		// return ParseSuccess(&VariableExpr{parser.Previous()}) // Assuming you have a VariableExpr type
	}
	if parser.Match([]TokenType{LEFT_BRACE}) {
		expr := parser.Expression()
		parser.Consume(RIGHT_BRACE, "Expect ')' after expression.")
		return ParseSuccess(&GroupingExpr{expr})
	}
	fmt.Println("hello? ", parser.Current, parser.Peek())
	filade := ParseFailed()
	filade.Expression = &LiteralExpr{StringLiteral("please work")}
	return filade
}

func (parser *Parser) Consume(tokenType TokenType, message string) Token {
	if parser.Check(tokenType) {
		return parser.Advance()
	}
	parser.Error(parser.Peek(), message)
	return parser.Peek() // ERRRO
}

func (parser *Parser) Error(token Token, message string) {
	error(token, message)
}

func (parser *Parser) Parse() Expr {
	return parser.Expression()
}
