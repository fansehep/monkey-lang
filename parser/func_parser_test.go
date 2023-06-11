package parser

import (
	"testing"

	"github.com/fansehep/monkey-lang/ast"
	"github.com/fansehep/monkey-lang/lexer"
)

func TestFuncLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExprStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExprStatment, get=%T", program.Statements[0])
	}

	function, ok := stmt.Expr.(*ast.FuncLiteral)

	if !ok {
		t.Fatalf("stmt.Expression is not a st.FunctionLiteral, get=%T", stmt.Expr)
	}
	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong, want 2, get=%T", len(function.Parameters))
	}
	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements, get=%d", len(function.Body.Statements))
	}
	bodyStmt, ok := function.Body.Statements[0].(*ast.ExprStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement, get=%T", function.Body.Statements[0])
	}
	testInfixExpression(t, bodyStmt.Expr, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: "fn() {}", expectedParams: []string{}},
		{input: "fn(x) {}", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}
    for _, tt := range tests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)
        stmt := program.Statements[0].(*ast.ExprStatement)
        function := stmt.Expr.(*ast.FuncLiteral)
        if len(function.Parameters) != len(tt.expectedParams) {
            t.Errorf("length parameters wront want: %d get=%d\n", len(tt.expectedParams), len(function.Parameters))
        }
        for i, ident := range tt.expectedParams {
            testLiteralExpression(t, function.Parameters[i], ident)
        }
    }
}
