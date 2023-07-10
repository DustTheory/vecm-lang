package parser

import (
	"fmt"

	"github.com/0xM-D/interpreter/ast"
)

func (p *Parser) parseConstDeclarationStatement() ast.Statement {
	p.nextToken()

	stmt := p.parseStatement()

	switch declStmt := stmt.(type) {
	case *ast.TypedDeclarationStatement:
		declStmt.IsConstant = true
	case *ast.AssignmentDeclarationStatement:
		declStmt.IsConstant = true
	default:
		msg := fmt.Sprintf("const cannot be applied to statement of type %T", declStmt)
		p.errors = append(p.errors, msg)
	}

	return stmt
}