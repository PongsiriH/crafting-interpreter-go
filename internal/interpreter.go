package internal

import "fmt"

type Interpreter struct {
	globalEnv *Environment
	env       *Environment
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		globalEnv: &env,
		env:       &env,
	}
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
	i.env.Define(expr.Name.Lexeme, value)
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
