package parser

import "github.com/fansehep/monkey-lang/ast"


type (
    prefixParseFn func() ast.Expression
    infixParseFn func(ast.Expression) ast.Expression
)
