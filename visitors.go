package main

import (
	"fmt"
	"strings"
)

// visitors (Printer, Interpreter, TypeChecker)
type BaseVisitor interface {
	VisitBinaryExpr(expr *BinaryExpr) any
	VisitUnaryExpr(expr *UnaryExpr) any
	VisitLiteralExpr(expr *LiteralExpr) any
	VisitGroupingExpr(expr *GroupingExpr) any
}

type Visitor[R ExprResult] interface {
	BaseVisitor
}

// ---- AstPrinter ---- 
type AstPrinter struct{}

func (printer *AstPrinter) Print(expr Expr) any {
	return expr.Accept(printer)
}

func (printer *AstPrinter) VisitBinaryExpr(expr *BinaryExpr) any {
	exprs := []Expr{expr.Left, expr.Right}
	return printer.Parenthesize(expr.Operator.Lexeme, exprs)
}
func (printer *AstPrinter) VisitGroupingExpr(expr *GroupingExpr) any {
	exprs := []Expr{expr.Expression}
	return printer.Parenthesize("group", exprs)
}
func (printer *AstPrinter) VisitLiteralExpr(expr *LiteralExpr) any {
	return expr.Value
}
func (printer *AstPrinter) VisitUnaryExpr(expr *UnaryExpr) any {
	exprs := []Expr{expr.Right}
	return printer.Parenthesize(expr.Operator.Lexeme, exprs)
}

func (printer *AstPrinter) Parenthesize(name string, exprs []Expr) any {
	var builder strings.Builder
	builder.WriteString("(")
	builder.WriteString(name)

	for _, expr := range exprs {
		builder.WriteString(" ")

		result := expr.Accept(printer)
		switch v := result.(type) {
		case string:
			builder.WriteString(v)
		case float32:
			builder.WriteString(fmt.Sprintf("%g", v))
		case float64:
			builder.WriteString(fmt.Sprintf("%g", v))
		default:
			builder.WriteString(fmt.Sprintf("%v", v))
		}
	}

	builder.WriteString(")")
	return builder.String()
}

// Interpreter
type Interpreter struct {}

