package parser

import (
	"testing"
    "github.com/fansehep/monkey-lang/ast"
	"github.com/fansehep/monkey-lang/lexer"
)

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x }`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements, got=%d\n", 1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExprStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement, get=%T", program.Statements[0])
	}
	exp, ok := stmt.Expr.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not a st.IfExpression, get=%T", stmt.Expr)
	}
	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements, get=%d", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExprStatement)
	if !ok {
		t.Fatalf("Statements[0] is not as t.ExpressionStatement, get=%T", exp.Consequence.Statements[0])
	}
	if !testIdentifier(t, consequence.Expr, "x") {
		return
	}
	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil, get=%+v", exp.Alternative)
	}
}
