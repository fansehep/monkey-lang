package executor

import (
	"testing"
)

func TestReturnStatements2(t *testing.T) {
    tests := []struct {
        input string
        expected int64
    } {
        {"return 10;", 10},
        {"return 10; 9", 10},
        {"return 2 * 5; 9;", 10},
        {"9; return 2 * 10; 9;", 20},
        {`if (10 > 1) { 
            if (10 > 1) {
                return 10;
            }
            return 1;
        }`, 10},
    }
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}

