package ast

import (
	"bytes"
	"fmt"
	"github.com/fansehep/monkey-lang/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {

}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("if%v %v", ie.Condition.String(), ie.Consequence.String()))
    if ie.Alternative != nil {
        out.WriteString(fmt.Sprintf("else %v", ie.Alternative.String()))
    }
    return out.String()
}


