package internal

type Stmt interface {
	Apply(VisitorStmt) any
}

type VisitorStmt interface {
	VisitExpression(Expression) any
	VisitPrint(Print) any
	VisitVarDeclare(VarDeclare) any
	VisitBlock(Block) any
	VisitIfStmt(IfStmt) any
	VisitWhileStmt(WhileStmt) any
	VisitFunctionStmt(FunctionStmt) any
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

type Block struct {
	Statements []Stmt
}

type IfStmt struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

type FunctionStmt struct {
	Name   Token
	Params []Token
	Body   Stmt
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

func (stmt Block) Apply(v VisitorStmt) any {
	return v.VisitBlock(stmt)
}

func (stmt IfStmt) Apply(v VisitorStmt) any {
	return v.VisitIfStmt(stmt)
}

func (stmt WhileStmt) Apply(v VisitorStmt) any {
	return v.VisitWhileStmt(stmt)
}

func (stmt FunctionStmt) Apply(v VisitorStmt) any {
	return v.VisitFunctionStmt(stmt)
}
