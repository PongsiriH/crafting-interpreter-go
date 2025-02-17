package main

import (
	"testing"
)

func TestAstPrinter(t *testing.T) {
	printer := &AstPrinter{}

	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{
			name: "test literal expression",
			expr: &LiteralExpr{
				Value: "123",
			},
			expected: "123",
		},
		{
			name: "test grouping expression",
			expr: &GroupingExpr{
				Expression: &LiteralExpr{
					Value: "123",
				},
			},
			expected: "(group 123)",
		},
		{
			name: "test unary expresion",
			expr: &UnaryExpr{
				Operator: Token{MINUS, "-", nil, 1},
				Right: &LiteralExpr{
					Value: "2",
				},
			},
			expected: "(- 2)",
		},
		{
			name: "test binary expression",
			expr: &BinaryExpr{
				Operator: Token{
					Lexeme: "+",
				},
				Left: &LiteralExpr{
					Value: "1",
				},
				Right: &LiteralExpr{
					Value: "2",
				},
			},
			expected: "(+ 1 2)",
		},
		{
			name: "test complex expression",
			expr: &BinaryExpr{
				Operator: Token{
					Lexeme: "*",
				},
				Left: &UnaryExpr{
					Operator: Token{
						Lexeme: "-",
					},
					Right: &LiteralExpr{
						Value: "123",
					},
				},
				Right: &GroupingExpr{
					Expression: &LiteralExpr{
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
