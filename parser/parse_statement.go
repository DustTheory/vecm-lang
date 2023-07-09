package parser

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IDENT:
		if p.peekToken.Type == token.IDENT {
			return p.parseTypedDeclarationStatement()
		}
		fallthrough
	default:
		return p.parseExpressionStatement()
	}
}
