package executor

import (
	"testing"

	"github.com/fansehep/monkey-lang/object"
)

func TestErrorHandling(t *testing.T) {
    tests := [] struct {
        input string
        expectedMessage string
    } {
        {
            "5 + true",
            "type mismatch: Integer + Boolean",
        },
        {
            "5 + true; 5;",
            "type mismatch: Integer + Boolean",
        },
        {
            "-true",
            "unknown operator: -Boolean",
        },
        {
            "true + false",
            "unknown operator: Boolean + Boolean",
        },
        {
            "5; true + false; 5",
            "unknown operator: Boolean + Boolean",
        },
        {
            "if (10 > 1) { true + false; }",
            "unknown operator: Boolean + Boolean",
        },
        {
            `
            if (10 > 1) {
                if (10 > 1) {
                    return true + false;
                }
            }
            return 1;
            `,
            "unknown operator: Boolean + Boolean",
        },
        {
            "foobadsada",
            "identifier not found: foobadsada",
        },
    }

    for _, tt := range tests {
        evaluated := testEval(tt.input)
        errObj, ok := evaluated.(*object.Error)
        if !ok {
            t.Errorf("no error object returned, get: %T (%+v)", evaluated, evaluated)
            continue
        }
        if errObj.Message != tt.expectedMessage {
            t.Errorf("wrong error message, expected=%q, get=%q", tt.expectedMessage, errObj.Message)
        }
    }
}