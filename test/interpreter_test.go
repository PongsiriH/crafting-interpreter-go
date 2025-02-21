package main

import (
	gx "golox/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterpreter_Interpret_Literal(t *testing.T) {
	// Test literal value (e.g., number 5)
	interpreter := &gx.Interpreter{}
	literalExpr := &gx.Literal{Value: 5.0}
	result := interpreter.Interpret(literalExpr)

	// Assert the result is the value of the literal
	assert.Equal(t, 5.0, result)
}

func TestInterpreter_Interpret_Binary(t *testing.T) {
	// Test binary operations (e.g., addition, subtraction)

	// 5 + 3
	left := &gx.Literal{Value: 5.0}
	right := &gx.Literal{Value: 3.0}
	operator := gx.Token{TokenType: gx.PLUS}
	binaryExpr := &gx.Binary{Left: left, Right: right, Operator: operator}
	interpreter := &gx.Interpreter{}
	result := interpreter.Interpret(binaryExpr)

	// Assert the result is the sum of the two values
	assert.Equal(t, 8.0, result)

	// 5 - 3
	operator = gx.Token{TokenType: gx.MINUS}
	binaryExpr = &gx.Binary{Left: left, Right: right, Operator: operator}
	result = interpreter.Interpret(binaryExpr)
	assert.Equal(t, 2.0, result)

	// 5 * 3
	operator = gx.Token{TokenType: gx.STAR}
	binaryExpr = &gx.Binary{Left: left, Right: right, Operator: operator}
	result = interpreter.Interpret(binaryExpr)
	assert.Equal(t, 15.0, result)

	// 5 / 3
	operator = gx.Token{TokenType: gx.SLASH}
	binaryExpr = &gx.Binary{Left: left, Right: right, Operator: operator}
	result = interpreter.Interpret(binaryExpr)
	assert.Equal(t, 1.6666666666666667, result)

	// 5 == 3
	operator = gx.Token{TokenType: gx.EQUAL_EQUAL}
	binaryExpr = &gx.Binary{Left: left, Right: right, Operator: operator}
	result = interpreter.Interpret(binaryExpr)
	assert.Equal(t, false, result)
}

func TestInterpreter_Interpret_Unary(t *testing.T) {
	// Test unary operations (e.g., negation, logical NOT)

	// -5
	left := &gx.Literal{Value: 5.0}
	operator := gx.Token{TokenType: gx.MINUS}
	unaryExpr := &gx.Unary{Operator: operator, Right: left}
	interpreter := &gx.Interpreter{}
	result := interpreter.Interpret(unaryExpr)

	// Assert the result is the negation of the literal value
	assert.Equal(t, -5.0, result)

	// !true
	right := &gx.Literal{Value: true}
	operator = gx.Token{TokenType: gx.BANG}
	unaryExpr = &gx.Unary{Operator: operator, Right: right}
	result = interpreter.Interpret(unaryExpr)

	// Assert the result is the negation of true (i.e., false)
	assert.Equal(t, false, result)

	// !false
	right = &gx.Literal{Value: false}
	unaryExpr = &gx.Unary{Operator: operator, Right: right}
	result = interpreter.Interpret(unaryExpr)

	// Assert the result is the negation of false (i.e., true)
	assert.Equal(t, true, result)
}

func TestInterpreter_Interpret_Comparison(t *testing.T) {
	// Test comparison operations (e.g., greater than, less than)

	// 5 > 3
	left := &gx.Literal{Value: 5.0}
	right := &gx.Literal{Value: 3.0}
	operator := gx.Token{TokenType: gx.GREATER}
	binaryExpr := &gx.Binary{Left: left, Right: right, Operator: operator}
	interpreter := &gx.Interpreter{}
	result := interpreter.Interpret(binaryExpr)

	// Assert the result is true
	assert.Equal(t, true, result)

	// 5 <= 3
	operator = gx.Token{TokenType: gx.LESS_EQUAL}
	binaryExpr = &gx.Binary{Left: left, Right: right, Operator: operator}
	result = interpreter.Interpret(binaryExpr)

	// Assert the result is false
	assert.Equal(t, false, result)
}

func TestInterpreter_Interpret_Equality(t *testing.T) {
	// Test equality operations (e.g., ==, !=)

	// 5 == 5
	left := &gx.Literal{Value: 5.0}
	right := &gx.Literal{Value: 5.0}
	operator := gx.Token{TokenType: gx.EQUAL_EQUAL}
	binaryExpr := &gx.Binary{Left: left, Right: right, Operator: operator}
	interpreter := &gx.Interpreter{}
	result := interpreter.Interpret(binaryExpr)

	// Assert the result is true
	assert.Equal(t, true, result)

	// 5 != 3
	operator = gx.Token{TokenType: gx.BANG_EQUAL}
	binaryExpr = &gx.Binary{Left: left, Right: right, Operator: operator}
	result = interpreter.Interpret(binaryExpr)

	// Assert the result is false
	assert.Equal(t, false, result)
}

func TestInterpreter_Interpret_Grouping(t *testing.T) {
	// Test grouping expressions (e.g., nested expressions)

	// (5 + 3)
	left := &gx.Literal{Value: 5.0}
	right := &gx.Literal{Value: 3.0}
	operator := gx.Token{TokenType: gx.PLUS}
	binaryExpr := &gx.Binary{Left: left, Right: right, Operator: operator}
	groupingExpr := &gx.Grouping{Inside: binaryExpr}
	interpreter := &gx.Interpreter{}
	result := interpreter.Interpret(groupingExpr)

	// Assert the result is the sum of the two values
	assert.Equal(t, 8.0, result)
}


func TestInterpreter_Interpret_Literal_EdgeCases(t *testing.T) {
	interpreter := &gx.Interpreter{}

	// Test 0
	literalZero := &gx.Literal{Value: 0.0}
	result := interpreter.Interpret(literalZero)
	assert.Equal(t, 0.0, result)

	// Test negative numbers
	literalNegative := &gx.Literal{Value: -42.5}
	result = interpreter.Interpret(literalNegative)
	assert.Equal(t, -42.5, result)

	// Test string literals
	literalString := &gx.Literal{Value: "Hello, Lox!"}
	result = interpreter.Interpret(literalString)
	assert.Equal(t, "Hello, Lox!", result)

	// Test boolean literals
	literalTrue := &gx.Literal{Value: true}
	result = interpreter.Interpret(literalTrue)
	assert.Equal(t, true, result)

	literalFalse := &gx.Literal{Value: false}
	result = interpreter.Interpret(literalFalse)
	assert.Equal(t, false, result)
}


func TestInterpreter_Interpret_DivisionByZero(t *testing.T) {
	interpreter := &gx.Interpreter{}

	left := &gx.Literal{Value: 5.0}
	right := &gx.Literal{Value: 0.0}
	operator := gx.Token{TokenType: gx.SLASH}
	binaryExpr := &gx.Binary{Left: left, Right: right, Operator: operator}

	// Expect an error or a specific behavior (e.g., return nil or panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected division by zero to cause an error or panic")
		}
	}()
	interpreter.Interpret(binaryExpr)
}

func TestInterpreter_Interpret_StringConcatenation(t *testing.T) {
  interpreter := &gx.Interpreter{}

  left := &gx.Literal{Value: "Hello"}
  right := &gx.Literal{Value: " World"}
  operator := gx.Token{TokenType: gx.PLUS}
  binaryExpr := &gx.Binary{Left: left, Right: right, Operator: operator}

  result := interpreter.Interpret(binaryExpr)
  assert.Equal(t, "Hello World", result)
}

func TestInterpreter_Interpret_3StringConcatenation(t *testing.T) {
  interpreter := &gx.Interpreter{}

  left := &gx.Literal{Value: "Hello"}
  mid := &gx.Literal{Value: " Good"}
  right := &gx.Literal{Value: " World"}
  operator := gx.Token{TokenType: gx.PLUS}

  leftMid := &gx.Binary{Left: left, Right: mid, Operator: operator}
  binaryExpr := &gx.Binary{Left: leftMid, Right: right, Operator: operator}

  result := interpreter.Interpret(binaryExpr)
  assert.Equal(t, "Hello Good World", result)
}

