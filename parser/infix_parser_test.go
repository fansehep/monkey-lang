package parser

import (
	"testing"

	"github.com/fansehep/monkey-lang/ast"
	"github.com/fansehep/monkey-lang/lexer"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        if len(program.Statements) != 1 {
            t.Fatalf("program.Statements does not contain %d statements, get=%d", 1, len(program.Statements))
        }
        stmt, ok := program.Statements[0].(*ast.ExprStatement)
        if !ok {
            t.Fatalf("expr is not ast.InfixExpression, get=%T", stmt.Expr)
        }

        exp, ok := stmt.Expr.(*ast.InfixExpression)
        if !ok {
            t.Fatalf("exp is not ast.InfixExpression, get=%T", stmt.Expr)
        }

        if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
            return
        }
        if exp.Operator != tt.operator {
            t.Fatalf("exp.Operator is not '%s', get=%s", tt.operator, exp.Operator)
        }
        if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
            return
        }
    }
}
