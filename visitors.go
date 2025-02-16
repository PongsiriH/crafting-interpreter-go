package main

import "strings"

type AstPrinter[T string] struct {}

func (printer *AstPrinter[T]) Print(expr Expr[T]) T {
  return expr.Accept(printer)
}

func (printer *AstPrinter[T]) VisitBinaryExpr(expr *BinaryExpr[T]) T {
  exprs := []Expr[T]{expr.Left, expr.Right}
  return printer.Paranthesize(expr.Operator.Lexeme, exprs)
}
func (printer *AstPrinter[T]) VisitGroupingExpr(expr *GroupingExpr[T]) T {
  exprs := []Expr[T]{expr.Expression}
  return printer.Paranthesize("group", exprs)
}
func (printer *AstPrinter[T]) VisitLiteralExpr(expr *LiteralExpr[T]) T {
  return expr.Value
}
func (printer *AstPrinter[T]) VisitUnaryExpr(expr *UnaryExpr[T]) T {
  exprs := []Expr[T]{expr.Right}
  return printer.Paranthesize(expr.Operator.Lexeme, exprs)
}

func (printer *AstPrinter[T]) Paranthesize(name string, exprs []Expr[T]) T {
  var builder strings.Builder
  builder.WriteString("(")
  builder.WriteString(name)
  for _, expr := range exprs {
    builder.WriteString(" ")
    builder.WriteString( string( expr.Accept(printer) ) )
  }
  builder.WriteString(")")
  return T(builder.String())
}
