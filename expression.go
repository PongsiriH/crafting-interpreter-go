package main

// expressions (Binary, Unary, Literal, etc.)
type ExprResult interface {
  string | float32 | bool
}

type Expr interface {
	Accept(visitor BaseVisitor) any
  // String() string
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
	Value any
}

type UnaryExpr struct {
	Operator Token
	Right    Expr
}

func (expr *BinaryExpr) Accept(visitor BaseVisitor) any {
	return visitor.VisitBinaryExpr(expr)
}

func (expr *GroupingExpr) Accept(visitor BaseVisitor) any {
	return visitor.VisitGroupingExpr(expr)
}

func (expr *LiteralExpr) Accept(visitor BaseVisitor) any {
	return visitor.VisitLiteralExpr(expr)
}

func (expr *UnaryExpr) Accept(visitor BaseVisitor) any {
	return visitor.VisitUnaryExpr(expr)
}
