package ast

import (
	"github.com/fansehep/monkey-lang/token"
)


type StringLiteral struct {
    Token token.Token
    Value string
}

func (s *StringLiteral) expressionNode() {

}

func (s *StringLiteral) TokenLiteral() string {
    return s.Token.Literal
}

func (s *StringLiteral) String() string {
    return s.Token.Literal
}