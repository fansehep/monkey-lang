package ast

import (
	"github.com/fansehep/monkey-lang/token"
)

type ExprStatement struct {
	Token token.Token
	Expr  Expression
}

func (es *ExprStatement) statementNode() {}

func (es *ExprStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExprStatement) String() string {
    if es.Expr != nil {
        return es.Expr.String()
    }
    return " "
}