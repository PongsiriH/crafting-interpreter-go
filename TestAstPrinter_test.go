package main

import (
	"testing"
)

func TestAstPrinter(t *testing.T) {
	printer := &AstPrinter[string]{}

	tests := []struct {
		name     string
		expr     Expr[string]
		expected string
	}{
		{
			name: "test literal expression",
			expr: &LiteralExpr[string]{
				Value: "123",
			},
			expected: "123",
		},
		{
			name: "test grouping expression",
			expr: &GroupingExpr[string]{
				Expression: &LiteralExpr[string]{
					Value: "123",
				},
			},
			expected: "(group 123)",
		},
		{
			name: "test unary expresion",
			expr: &UnaryExpr[string]{
				Operator: Token{MINUS, "-", nil, 1},
				Right: &LiteralExpr[string]{
					Value: "2",
				},
			},
			expected: "(- 2)",
		},
		{
			name: "test binary expression",
			expr: &BinaryExpr[string]{
				Operator: Token{
					Lexeme: "+",
				},
				Left: &LiteralExpr[string]{
					Value: "1",
				},
				Right: &LiteralExpr[string]{
					Value: "2",
				},
			},
			expected: "(+ 1 2)",
		},
		{
			name: "test complex expression",
			expr: &BinaryExpr[string]{
				Operator: Token{
					Lexeme: "*",
				},
				Left: &UnaryExpr[string]{
					Operator: Token{
						Lexeme: "-",
					},
					Right: &LiteralExpr[string]{
						Value: "123",
					},
				},
				Right: &GroupingExpr[string]{
					Expression: &LiteralExpr[string]{
						Value: "45.67",
					},
				},
			},
			expected: "(* (- 123) (group 45.67))",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := printer.Print(tt.expr)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
