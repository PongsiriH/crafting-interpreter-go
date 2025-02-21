package internal

type Stmt interface {
	Apply(VisitorStmt) any
}

type VisitorStmt interface {
	VisitExpression(Expression) any
	VisitPrint(Print) any
	VisitVarDeclare(VarDeclare) any
}

type Expression struct {
	Expr Expr
}

type Print struct {
	Expr Expr
}

type VarDeclare struct {
	Name        string
	InitialExpr Expr
}

func (stmt Expression) Apply(v VisitorStmt) any {
	return v.VisitExpression(stmt)
}

func (stmt Print) Apply(v VisitorStmt) any {
	return v.VisitPrint(stmt)
}

func (stmt VarDeclare) Apply(v VisitorStmt) any {
  return v.VisitVarDeclare(stmt)
}


