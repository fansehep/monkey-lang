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
    env := object.NewEnvironment()
    return Eval(program, env)
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
        {"if (1 < 2) { 10 }", 10},
        {"if (1 > 2) { 10 }", nil},
        {"if (1 > 2) { 10} else { 20 }", 20},
        {"if (1 < 2) { 10} else { 20 }", 10},
    }
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        integer, ok := tt.expected.(int)
        if ok {
            testIntegerObject(t, evaluated, int64(integer))
        } else {
           testNullObject(t, evaluated)
        }
    }
}

func testNullObject(t *testing.T, obj object.Object) bool {
    if obj != kNull {
        t.Errorf("object is not Null, get: %T (%+v)", obj, obj)
        return false
    }
    return true
}