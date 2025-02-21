package internal

import "fmt"

type Parser struct {
	Tokens  []Token
	Current int
}

func NewParser(tokens []Token) Parser {
	return Parser{
		Tokens: tokens,
	}
}

func (p *Parser) Parse() []Stmt {
	statements := []Stmt{}
	for p.Tokens[p.Current].TokenType != EOF {
		statements = append(statements, p.Declaration())
	}
	return statements
}

func (p *Parser) Declaration() Stmt {
	if p.Match(VAR) {
		p.Current++
		return p.VarDeclaration()
	}
	return p.Statement()
}

func (p *Parser) VarDeclaration() Stmt {
  p.Consume(IDENTIFIER, "Expected variable name.")
  name := p.Tokens[p.Current-1]
  var val Expr
  if p.Match(EQUAL) {
    p.Current++
    val = p.Expression()
  }
  p.Consume(SEMICOLON, "Expected semicolon `;' after variable declaration.")
  return VarDeclare{name.Lexeme, val}
}

func (p *Parser) Statement() Stmt {
	if p.Match(PRINT) {
		p.Current++
		return p.PrintStmt()
	}

  if p.Match(LEFT_BRACE) {
    p.Current++
    return p.BlockStmt()
  }

	return p.ExpressionStmt()
}

func (p *Parser) PrintStmt() Stmt {
	expr := p.Expression()
	p.Consume(SEMICOLON, "Expected semicolon `;' after expression.")
	return Print{expr}
}

func (p *Parser) ExpressionStmt() Stmt {
	expr := p.Expression()
	p.Consume(SEMICOLON, "Expected semicolon `;' after expression.")
	return Expression{expr}
}

func (p *Parser) BlockStmt() Stmt {
  statements := []Stmt{}
  for !p.Match(RIGHT_BRACE) && p.Tokens[p.Current].TokenType != EOF {
    statements = append(statements, p.Declaration())
  }
  p.Consume(RIGHT_BRACE, "Expected closing brace `}' after block.")
  return Block{statements}
}

func (p *Parser) Expression() Expr {
  return p.Assignment()
}

func (p *Parser) Assignment() Expr {
  expr := p.Equality()
  if p.Match(EQUAL) {
    p.Current++
    value := p.Assignment()
    expr, ok := expr.(*Variable)
    if ok {
      return &Assignment{expr.Name, value}
    }

    panic("Invalid assignment target")
  }
 	return expr
}

func (p *Parser) Equality() Expr {
	expr := p.Comparison()
	for p.Match(EQUAL_EQUAL, BANG_EQUAL) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.Comparison()
		expr = &Binary{expr, op, right}
	}
	return expr
}

func (p *Parser) Comparison() Expr {
	expr := p.Term()
	for p.Match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.Term()
		expr = &Binary{expr, op, right}
	}
	return expr
}

func (p *Parser) Term() Expr {
	expr := p.Factor()
	for p.Match(MINUS, PLUS) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.Factor()
		expr = &Binary{expr, op, right}
	}
	return expr
}

func (p *Parser) Factor() Expr {
	expr := p.Unary()
	for p.Match(SLASH, STAR) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.Unary()
		expr = &Binary{expr, op, right}
	}
	return expr
}

func (p *Parser) Unary() Expr {
	if p.Match(MINUS, BANG) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.Unary()
		return &Unary{op, right}

	}
	return p.Primary()
}

func (p *Parser) Primary() Expr {
	switch p.Tokens[p.Current].TokenType {
	case FALSE:
		p.Current++
		return &Literal{false}
	case TRUE:
		p.Current++
		return &Literal{true}
	case NIL:
		p.Current++
		return &Literal{nil}
	case NUMBER, STRING:
		literal := p.Tokens[p.Current].Literal
		p.Current++
		return &Literal{literal}
	case IDENTIFIER:
		token := p.Tokens[p.Current]
		p.Current++
		return &Variable{token}
	case LEFT_PAREN:
		p.Current++
		expr := p.Expression()
		p.Consume(RIGHT_PAREN, "Expected closing parenthesis ')' after expression.")
		return &Grouping{expr}
	default:
		return &Literal{"Token not supported"}
	}
}

// helper
func (p *Parser) Match(tokenTypes ...TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.Tokens[p.Current].TokenType == tokenType {
			return true
		}
	}
	return false
}

func (p *Parser) Consume(expected TokenType, message string) {
	if p.Tokens[p.Current].TokenType == expected {
		p.Current++
		return
	}
	panic(fmt.Sprintf("%s.. Got %s ", message, p.Tokens[p.Current]))
}

func (p *Parser) Synchronize() {
	for p.Current < len(p.Tokens) {
		if p.Tokens[p.Current-1].TokenType == SEMICOLON {
			return
		}

		switch p.Tokens[p.Current].TokenType {
		case CLASS:
		case FUN:
		case VAR:
		case FOR:
		case IF:
		case WHILE:
		case PRINT:
		case RETURN:
			return
		}
	}
	p.Current++
}
