package main

type ExprResult interface {
  isResult()
}

type StringResult string
func (StringResult) isResult() {}

type NumberResult float64
func (NumberResult) isResult() {}

type BoolResult float64
func (BoolResult) isResult() {}

type Literal interface {
  IsLiteral()
}

type StringLiteral string
func (StringLiteral) IsLiteral() {}

type NumberLiteral float64
func (NumberLiteral) IsLiteral() {}

type BoolLiteral float64
func (BoolLiteral) IsLiteral() {}

// expressions (Binary, Unary, Literal, etc.)
type Expr interface {
	Accept(visitor BaseVisitor) ExprResult
}

type BinaryExpr struct {
	Left     Expr
	Operator Token
	Right    Expr
}

type GroupingExpr struct {
	Expression Expr
}

type LiteralExpr struct {
	Value Literal
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (expr *BinaryExpr) Accept(visitor BaseVisitor) ExprResult {
	return visitor.VisitBinaryExpr(expr)
}

func (expr *GroupingExpr) Accept(visitor BaseVisitor) ExprResult {
	return visitor.VisitGroupingExpr(expr)
}

func (expr *LiteralExpr) Accept(visitor BaseVisitor) ExprResult {
	return visitor.VisitLiteralExpr(expr)
}

func (expr *UnaryExpr) Accept(visitor BaseVisitor) ExprResult {
	return visitor.VisitUnaryExpr(expr)
}
