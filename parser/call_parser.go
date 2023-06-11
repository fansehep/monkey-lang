package parser

import (
    "github.com/fansehep/monkey-lang/token"
	"github.com/fansehep/monkey-lang/ast"
)

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{
		Token:    p.curToken,
		Function: function,
	}
	exp.Arguments = p.parseCallArguments()
	return exp
}


func (p *Parser) parseCallArguments() []ast.Expression {
    args := []ast.Expression{}
    if p.peekTokenIs(token.RPAREN) {
        p.nextToken()
        return args
    }
    p.nextToken()
    args = append(args, p.parseExpression(LOWEST))
    for p.peekTokenIs(token.COMMA) {
        p.nextToken()
        p.nextToken()
        args = append(args, p.parseExpression(LOWEST))
    }
    if !p.expectPeek(token.RPAREN) {
        return nil
    }
    return args
}