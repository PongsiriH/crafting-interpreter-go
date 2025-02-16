package main


type Expr[T any] interface {
    Accept(visitor Visitor[T]) T
}
            
type Visitor[T any] interface {
	VisitBinaryExpr(expr *BinaryExpr[T]) T
	VisitGroupingExpr(expr *GroupingExpr[T]) T
	VisitLiteralExpr(expr *LiteralExpr[T]) T
	VisitUnaryExpr(expr *UnaryExpr[T]) T
}
        
type BinaryExpr[T any] struct {
	Left Expr[T]
	Operator Token
	Right Expr[T]
}
                
func (expr * BinaryExpr[T]) Accept(visitor Visitor[T]) T {
    return visitor.VisitBinaryExpr(expr)}
                
type GroupingExpr[T any] struct {
	Expression Expr[T]
}
                
func (expr * GroupingExpr[T]) Accept(visitor Visitor[T]) T {
    return visitor.VisitGroupingExpr(expr)}
                
type LiteralExpr[T any] struct {
	Value T
}
                
func (expr * LiteralExpr[T]) Accept(visitor Visitor[T]) T {
    return visitor.VisitLiteralExpr(expr)}
                
type UnaryExpr[T any] struct {
	Operator Token
	Right Expr[T]
}
                
func (expr * UnaryExpr[T]) Accept(visitor Visitor[T]) T {
    return visitor.VisitUnaryExpr(expr)}
                