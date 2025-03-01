package internal

import "fmt"

type Interpreter struct {
	globalEnv *Environment
	env       *Environment
}

func NewInterpreter() *Interpreter {
  i := Interpreter{
		globalEnv: &env,
		env:       &env,
	}
  i.globalEnv.Define("clock", GlobalClock{})
  return &i
}

func (i *Interpreter) Interpret(statements []Stmt) {
	for _, stmt := range statements {
		stmt.Apply(i)
	}
}

func (i *Interpreter) VisitLiteralExpr(expr Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitGroupingExpr(expr Grouping) any {
	return expr.Inside.Apply(i)
}

func (i *Interpreter) VisitBinaryExpr(expr Binary) any {
	left := expr.Left.Apply(i)
	right := expr.Right.Apply(i)
	op := expr.Operator.TokenType

	switch op {
	case MINUS:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal - rightVal
		}

	case STAR:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal * rightVal
		}

	case SLASH:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal / rightVal
		}

	case PLUS:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal + rightVal
		}

		leftString, leftStrOk := left.(string)
		rightString, rightStrOk := right.(string)
		if leftStrOk && rightStrOk {
			return leftString + rightString
		}

	case GREATER:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal > rightVal
		}

	case GREATER_EQUAL:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal >= rightVal
		}

	case LESS:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal < rightVal
		}

	case LESS_EQUAL:
		leftVal, leftOk := left.(float64)
		rightVal, rightOk := right.(float64)
		if leftOk && rightOk {
			return leftVal <= rightVal
		}
	case EQUAL_EQUAL:
		return isEqual(left, right)
	case BANG_EQUAL:
		return !isEqual(left, right)
	}
	return nil
}

func (i *Interpreter) VisitUnaryExpr(expr Unary) any {
	op := expr.Operator.TokenType
	right := expr.Right.Apply(i)
	switch op {
	case MINUS:
		rightVal, ok := right.(float64)
		if ok {
			return -rightVal
		}
	case BANG:
		return !isTruthy(right)
	}
	return nil
}

func (i *Interpreter) VisitVariableExpr(expr Variable) any {
	return i.env.Get(expr.Name.Lexeme)
}

func (i *Interpreter) VisitExpression(stmt Expression) any {
	return stmt.Expr.Apply(i)
}

func (i *Interpreter) VisitPrint(stmt Print) any {
	val := stmt.Expr.Apply(i)
	fmt.Println(">>", val)
	return val
}

func (i *Interpreter) VisitVarDeclare(stmt VarDeclare) any {
	var value any
	if stmt.InitialExpr != nil {
		value = stmt.InitialExpr.Apply(i)
	}
	i.env.Define(stmt.Name, value)
	return nil
}

func (i *Interpreter) VisitAssignmentExpr(expr Assignment) any {
	value := expr.Value.Apply(i)
	i.env.Assign(expr.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitBlock(stmt Block) any {
	upperEnv := i.env
	defer func() {
		i.env = upperEnv
	}()
	i.env = NewEnvironment(*upperEnv)

	var output any
	for _, statement := range stmt.Statements {
		output = statement.Apply(i)
	}
	return output
}

func (i *Interpreter) VisitIfStmt(stmt IfStmt) any {
	var output any
	if isTruthy(stmt.Condition.Apply(i)) {
		output = stmt.ThenBranch.Apply(i)
	} else {
		output = stmt.ElseBranch.Apply(i)
	}
	return output
}

func (i *Interpreter) VisitLogicalExpr(expr Logic) any {
	left := expr.Left.Apply(i)
	op := expr.Operator.TokenType
	switch op {
	case OR:
		if isTruthy(left) {
			return true
		}
	case AND:
		if !isTruthy(left) {
			return false
		}
	}
	right := expr.Right.Apply(i)
	return isTruthy(right)
}

func (i *Interpreter) VisitWhileStmt(expr WhileStmt) any {
	cond := expr.Condition.Apply(i)
	for isTruthy(cond) {
		expr.Body.Apply(i)
		cond = expr.Condition.Apply(i)
	}
	return nil
}

func (i *Interpreter) VisitCallExpr(expr Call) any {
	callee := expr.Callee.Apply(i)

	args := []any{}
	for _, arg := range expr.Arguments {
	  args = append(args, arg)
	}
	callable, ok := callee.(LoxCallable)
  if !ok {
    fmt.Println("Not a function", callable)
    panic("Trying to call not a function")
  } 
  if callable.Arity() != len(args) {
    panic(fmt.Sprintf("Runtime error: Expected %d arguments but got %d arguments", callable.Arity(), len(args)))
  }

  callable.Call(i, &args)
	return nil
}

func (i *Interpreter) VisitFunctionStmt(stmt FunctionStmt) any {
  function := &LoxFunction{
    Declaration: stmt,
    Closure: i.globalEnv,
  }
  i.env.Define(stmt.Name.Lexeme, function)
  return nil
}
