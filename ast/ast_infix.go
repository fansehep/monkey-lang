package ast

import (
	"bytes"
    "fmt"
	"github.com/fansehep/monkey-lang/token"
)


type InfixExpression struct {
    Token token.Token
    Left Expression
    Operator string
    Right Expression
}

func (ie *InfixExpression) expressionNode() {}


func (ie *InfixExpression) TokenLiteral() string {
    return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
    var out bytes.Buffer
    out.WriteString(fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String()))
    return out.String()
}