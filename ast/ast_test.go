package ast

import (
	"github.com/fansehep/monkey-lang/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVal",
					},
					Value: "myVal",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVal",
				},
			},
		},
	}
	if program.String() != "let myVal = anotherVal;" {
		t.Errorf("program.String() wrong, get=%q", program.String())
	}
}
