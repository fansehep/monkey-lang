package parser

import (
	"fmt"
	"github.com/fansehep/monkey-lang/ast"
	"github.com/fansehep/monkey-lang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. get = %d",
			len(program.Statements))
	}
	tests := []struct {
		expected_identifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expected_identifier) {
			return
		}
	}
}

func TestLetStatements2(t *testing.T) {
	tests := []struct {
		input               string
		expected_identifier string
		expected_value      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}
	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements, get=%d", len(program.Statements))
		}
		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expected_identifier) {
			return
		}
		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expected_value) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let' get = %q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("LetStmt.Name.Value not '%s'. got='%s'", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
    return 5;
    return 10;
    return 933 322;
    `
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, get=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', get=%q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements get=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExprStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.Expression, get=%T", program.Statements[0])
	}
	ident, ok := stmt.Expr.(*ast.Identifier)
	if !ok {
		t.Fatalf("expr not *ast.Identifier get=%T", stmt.Expr)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident value not %s get = %s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident TokenLiteral not %s get=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements, get=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExprStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement get=%T", program.Statements[0])
	}

	literal, ok := stmt.Expr.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral, get=%T", stmt.Expr)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d, get=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s get=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue interface{}
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements doest not contain %d statements, get=%d\n", 1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExprStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement get=%T", program.Statements[0])
		}
		exp, ok := stmt.Expr.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expr.Operator is not '%s' get=%s", tt.operator, exp.Operator)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("expr.Operator is not '%s', get=%s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue.(int64)) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral, get=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. get=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d, got=%s", value, integ.TokenLiteral())
		return false
	}
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier get=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s get=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s get=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s get=%s", value, ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled, get=%T", exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression, get=%T", exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s', get=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean get=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t get=%t", value, bo.Value)
		return false
	}
	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral not %t get=%s", value, bo.TokenLiteral())
		return false
	}
	return true
}
