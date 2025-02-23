package internal

import "fmt"

type Expr interface {
	Apply(VisitorExpr) any
}

type VisitorExpr interface {
	VisitBinaryExpr(expr Binary) any
	VisitUnaryExpr(expr Unary) any
	VisitLiteralExpr(expr Literal) any
	VisitGroupingExpr(expr Grouping) any
	VisitVariableExpr(expr Variable) any
	VisitAssignmentExpr(expr Assignment) any
	VisitCallExpr(expr Call) any
	VisitLogicalExpr(expr Logic) any
}

type Binary struct {
	Left     Expr
	Operator Token
	Right    Expr
}

type Unary struct {
	Operator Token
	Right    Expr
}

type Literal struct {
	Value any
}

type Grouping struct {
	Inside Expr
}

type Variable struct {
	Name Token
}

type Assignment struct {
	Name  Token
	Value Expr
}

type Call struct {
	Callee    Expr
	Arguments []Expr
}

type Logic struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (expr *Binary) Apply(v VisitorExpr) any {
	return v.VisitBinaryExpr(*expr)
}

func (expr *Unary) Apply(v VisitorExpr) any {
	return v.VisitUnaryExpr(*expr)
}

func (expr *Literal) Apply(v VisitorExpr) any {
	return v.VisitLiteralExpr(*expr)
}

func (expr *Grouping) Apply(v VisitorExpr) any {
	return v.VisitGroupingExpr(*expr)
}

func (expr *Variable) Apply(v VisitorExpr) any {
	return v.VisitVariableExpr(*expr)
}

func (expr *Assignment) Apply(v VisitorExpr) any {
	return v.VisitAssignmentExpr(*expr)
}

func (expr *Call) Apply(v VisitorExpr) any {
	return v.VisitCallExpr(*expr)
}

func (expr *Logic) Apply(v VisitorExpr) any {
	return v.VisitLogicalExpr(*expr)
}

func (expr *Binary) String() string {
	return fmt.Sprintf("Binary(%v, %v, %v)", expr.Left, expr.Operator, expr.Right)
}

func (expr *Unary) String() string {
	return fmt.Sprintf("Unary(%v, %v)", expr.Operator, expr.Right)
}

func (expr *Literal) String() string {
	return fmt.Sprintf("Literal(%v)", expr.Value)
}

func (expr *Grouping) String() string {
	return fmt.Sprintf("Grouping(%v)", expr.Inside)
}

func (expr *Variable) String() string {
	return fmt.Sprintf("Variable(%v)", expr.Name)
}

func (expr *Assignment) String() string {
	return fmt.Sprintf("Assignment(%v, %v)", expr.Name, expr.Value)
}

func (expr *Logic) String() string {
	return fmt.Sprintf("Logic(%v, %v, %v)", expr.Left, expr.Operator, expr.Right)
}
