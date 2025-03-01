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

	if p.Match(FUN) {
		p.Current++
		return p.Function("function")
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

func (p *Parser) Function(kind string) Stmt {
	p.Consume(IDENTIFIER, "Expected "+kind+" name")
	name := p.Tokens[p.Current-1]
	p.Consume(LEFT_PAREN, "Expected left_paren name")
	params := []Token{}
	if p.Tokens[p.Current].TokenType != RIGHT_PAREN {
		for {
			p.Consume(IDENTIFIER, "Expected parameter name")
			params = append(params, p.Tokens[p.Current-1])
			if p.Tokens[p.Current].TokenType != COMMA {
				break
			}
		}
	}
	p.Consume(RIGHT_PAREN, "Expected right_paren name")

  p.Consume(LEFT_BRACE, "Expected left_brace ")
	body := p.BlockStmt()
	return &FunctionStmt{name, params, body}
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

	if p.Match(VAR) {
		p.Current++
		return p.VarDeclaration()
	}

	if p.Match(IF) {
		p.Current++
		return p.IfStmt()
	}

	if p.Match(WHILE) {
		p.Current++
		return p.WhileStmt()
	}

	if p.Match(FOR) {
		p.Current++
		return p.ForStmt()
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

func (p *Parser) IfStmt() Stmt {
	p.Consume(LEFT_PAREN, "Expected opening parenthesis '(' after if.")
	condition := p.Expression()
	p.Consume(RIGHT_PAREN, "Expected closing parenthesis ')' after expression.")

	thenBranch := p.Statement()
	var elseBranch Stmt
	if p.Match(ELSE) {
		p.Current++
		elseBranch = p.Statement()
	}
	return IfStmt{condition, thenBranch, elseBranch}
}

func (p *Parser) Expression() Expr {
	return p.Assignment()
}

func (p *Parser) WhileStmt() Stmt {
	p.Consume(LEFT_PAREN, "Expected opening parenthesis '(' after while statement")
	cond := p.Expression()
	p.Consume(RIGHT_PAREN, "Expected closing parenthesis ')' after expression")
	body := p.Statement()
	return &WhileStmt{cond, body}
}

func (p *Parser) ForStmt() Stmt {
	p.Consume(LEFT_PAREN, "Expected opening parenthesis '(' after for statement")
	var Initializer Stmt
	if p.Match(SEMICOLON) {
		p.Current++
		Initializer = nil
	} else if p.Match(VAR) {
		p.Current++
		Initializer = p.VarDeclaration()
	} else {
		Initializer = p.ExpressionStmt()
	}

	var Cond Expr
	if !p.Match(SEMICOLON) {
		Cond = p.Expression()
	}
	p.Consume(SEMICOLON, "Expected semicolon `;' after for loop condition.")

	var Increment Expr
	if !p.Match(RIGHT_PAREN) {
		Increment = p.Expression()
	}
	p.Consume(RIGHT_PAREN, "Expected opening parenthesis ')' after clauses")

	Body := p.Statement()

	if Increment != nil {
		Body = Block{[]Stmt{
			Body,
			Expression{Increment},
		}}
	}

	if Cond == nil {
		Cond = &Literal{true}
	}
	Body = &WhileStmt{Cond, Body}

	if Initializer != nil {
		Body = Block{[]Stmt{
			Initializer,
			Body,
		}}
	}
	return Body
}

func (p *Parser) Assignment() Expr {
	expr := p.Or()
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

func (p *Parser) Or() Expr {
	expr := p.And()
	for p.Match(OR) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.And()
		expr = &Logic{expr, op, right}
	}
	return expr
}

func (p *Parser) And() Expr {
	expr := p.Equality()
	for p.Match(AND) {
		op := p.Tokens[p.Current]
		p.Current++
		right := p.Equality()
		expr = &Logic{expr, op, right}
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
	return p.Call()
}

func (p *Parser) Call() Expr {
	expr := p.Primary()
	for {
		if p.Match(LEFT_PAREN) {
			p.Current++
			expr = p.FinishCall(expr)
		} else {
			break
		}
	}
	return expr
}

func (p *Parser) FinishCall(callee Expr) Expr {
	arguments := []Expr{}
	if !p.Match(RIGHT_PAREN) {
		for {
			arguments = append(arguments, p.Expression())
			if !p.Match(COMMA) {
				break
			}
			p.Current++
		}
	}
	p.Consume(RIGHT_PAREN, "Expected closing parenthesis ')' after arguments.")
	paren := p.Tokens[p.Current-1]
	return &Call{callee, paren, arguments}
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
