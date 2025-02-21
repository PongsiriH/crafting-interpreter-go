package internal

import "fmt"

type Interpreter struct{}

func (i *Interpreter) Interpret(statements []Stmt) {
	for _, stmt := range statements {
		result := stmt.Apply(i)
		fmt.Println(">", result)
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
	return env.Get(expr.Name.Lexeme)
}

func (i *Interpreter) VisitExpression(stmt Expression) any {
	return stmt.Expr.Apply(i)
}

func (i *Interpreter) VisitPrint(stmt Print) any {
	return stmt.Expr.Apply(i)
}

func (i *Interpreter) VisitVarDeclare(stmt VarDeclare) any {
	var value any
	if stmt.InitialExpr != nil {
		value = stmt.InitialExpr.Apply(i)
	}
	env.Define(stmt.Name, value)
	return nil
}

func (i *Interpreter) VisitAssignmentExpr(stmt Assignment) any {
	value := stmt.Value.Apply(i)
	env.Define(stmt.Name.Lexeme, value)
	return nil
}
