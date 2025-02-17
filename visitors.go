package main

import (
	"fmt"
	"strings"
)

// visitors (Printer, Interpreter, TypeChecker)
type BaseVisitor interface {
	VisitBinaryExpr(expr *BinaryExpr) ExprResult
	VisitUnaryExpr(expr *UnaryExpr) ExprResult
	VisitLiteralExpr(expr *LiteralExpr) ExprResult
	VisitGroupingExpr(expr *GroupingExpr) ExprResult
}

type Visitor[R ExprResult] interface {
	BaseVisitor
}

// ---- AstPrinter ----
type AstPrinter struct{}

func (printer *AstPrinter) Print(expr Expr) ExprResult {
	return expr.Accept(printer)
}

func (printer *AstPrinter) VisitBinaryExpr(expr *BinaryExpr) ExprResult {
	exprs := []Expr{expr.Left, expr.Right}
	return printer.Parenthesize(expr.Operator.Lexeme, exprs)
}
func (printer *AstPrinter) VisitGroupingExpr(expr *GroupingExpr) ExprResult {
	exprs := []Expr{expr.Expression}
	return printer.Parenthesize("group", exprs)
}

func (printer *AstPrinter) VisitLiteralExpr(expr *LiteralExpr) ExprResult {
	exprLiteral := expr.Value
	var output ExprResult
	switch exprLiteral.(type) {
	case StringLiteral:
		output = StringResult(exprLiteral.(StringLiteral))
	case NumberLiteral:
		output = NumberResult(exprLiteral.(NumberLiteral))
	case BoolLiteral:
		output = BoolResult(exprLiteral.(BoolLiteral))
	}
	return output
}
func (printer *AstPrinter) VisitUnaryExpr(expr *UnaryExpr) ExprResult {
	exprs := []Expr{expr.Right}
	return printer.Parenthesize(expr.Operator.Lexeme, exprs)
}

func (printer *AstPrinter) Parenthesize(name string, exprs []Expr) ExprResult {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)

	for _, expr := range exprs {
		builder.WriteString(" ")

		result := expr.Accept(printer)
		switch result.(type) {
		case StringResult:
			builder.WriteString(fmt.Sprintf("%s", result))
		case NumberResult:
			builder.WriteString(fmt.Sprintf("%g", result))
		case BoolResult:
			builder.WriteString(fmt.Sprintf("%g", result))
		default:
			builder.WriteString(fmt.Sprintf("%v", result))
		}
	}

	builder.WriteString(")")
	return StringResult(builder.String())
}

// Interpreter
type Interpreter struct{}
