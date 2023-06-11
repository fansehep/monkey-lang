package ast

import (
	"bytes"
	"fmt"
	"github.com/fansehep/monkey-lang/token"
	"strings"
)

type FuncLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FuncLiteral) expressionNode() {

}

func (fl *FuncLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FuncLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fmt.Sprintf("%v (%v) %v", fl.TokenLiteral(),
		strings.Join(params, ", "), fl.Body.String()))
	return out.String()
}
