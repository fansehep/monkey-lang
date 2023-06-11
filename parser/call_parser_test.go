package parser

import (
	"testing"

	"github.com/fansehep/monkey-lang/ast"
	"github.com/fansehep/monkey-lang/lexer"
)


func TestCallExpressionParsing(t *testing.T) {
    input := "add(1, 2 * 3, 4 + 5);"
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParserErrors(t, p)
    if len(program.Statements) != 1 {
        t.Fatalf("program.Statements does not contain %d statements get=%d", 1, len(program.Statements))
    }
    stmt, ok := program.Statements[0].(*ast.ExprStatement)
    if !ok {
        t.Fatalf("stmt is not ast.ExpressionStatement, get=%T", program.Statements[0])
    }
    if !ok {
        t.Fatalf("stmt.Expression is not ast.CallExpresion. get=%T", stmt.Expr)
    }
    //
    exp, ok := stmt.Expr.(*ast.CallExpression)
    if !ok {
        t.Fatalf("stmt.Expression is not ast.CallExpression, get=%T", stmt.Expr)
    }
    if !testIdentifier(t, exp.Function, "add") {
        return
    }
    if len(exp.Arguments) != 3 {
        t.Fatalf("wrong length of arguments, get=%d", len(exp.Arguments))
    }
    testLiteralExpression(t, exp.Arguments[0], 1)
    testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
    testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}