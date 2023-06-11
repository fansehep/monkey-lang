package executor


import (
    "github.com/fansehep/monkey-lang/object"
	"github.com/fansehep/monkey-lang/lexer"
    "github.com/fansehep/monkey-lang/parser"
    "testing"
)

func TestExecutorExpression(t *testing.T) {
    tests := []struct {
        input string
        expected int64
    } {
        {"5", 5},
        {"10", 10},
    }
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}

func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()
    return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
    result, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("object is not Integer. get=%T (%+v)", obj, obj)
        return false
    }
    if result.Value != expected {
        t.Errorf("object has wrong value. get=%v, want=%d", result.Value, expected)
        return false
    }
    return true
}

func TestIfElseExpressions(t *testing.T) {
    tests := [] struct {
        input string
        expected interface{}
    } {
        {"if (true) { 10 }", 10},
        {"if (false) { 10 }", nil},
        {"if (1) {10}", 10},
        {""}
    }
}