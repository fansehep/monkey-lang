package ast

import (
	"bytes"
    "fmt"
    "strings"
	"github.com/fansehep/monkey-lang/token"
)

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
    return ce.Token.Literal
}

func (ce *CallExpression) String() string {
    var out bytes.Buffer
    args := []string{}
    for _, a := range ce.Arguments {
        args = append(args, a.String())
    }
    out.WriteString(fmt.Sprintf("%v(%v)", ce.Function.String(), strings.Join(args, ", ")))
    return out.String()
}
