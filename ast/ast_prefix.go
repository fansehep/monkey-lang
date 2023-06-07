package ast

import (
	"bytes"
	"fmt"
	"github.com/fansehep/monkey-lang/token"
)

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {

}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("(%v%v)", pe.Operator, pe.Right.String()))
	return out.String()
}
