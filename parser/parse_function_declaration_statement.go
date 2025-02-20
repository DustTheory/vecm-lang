package parser

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/token"
)

func (p *Parser) parseFunctionDeclarationStatement() *ast.FunctionDeclarationStatement {
	funcToken := p.curToken
	p.nextToken() // Swallow "fn" token

	functionName := p.parseIdentifier()

	// Swallow "(" token
	if !p.expectPeek(token.LeftParen) {
		return nil
	}

	functionParams, functionParamTypes, isVariadic := p.parseFunctionParameters()

	// Swallow ")" token
	if !p.expectPeek(token.DashArrow) {
		return nil
	}

	p.nextToken() // Swallow "->" token

	// Parse return type
	functionReturnType := p.parseType()
	if functionReturnType == nil {
		return nil
	}

	var functionBody *ast.BlockStatement
	if p.peekTokenIs(token.Semicolon) {
		// Swallow ";" token
		p.nextToken()
	} else {
		// Swallow "{" token
		if !p.expectPeek(token.LeftBrace) {
			return nil
		}

		functionBody = p.parseBlockStatement()
	}

	stmt := &ast.FunctionDeclarationStatement{
		Token:      funcToken,
		Name:       functionName,
		Body:       functionBody,
		Parameters: functionParams,
		IsVariadic: isVariadic,
		Type: ast.FunctionType{
			Token:          funcToken,
			ParameterTypes: functionParamTypes,
			ReturnType:     functionReturnType,
		},
	}

	return stmt
}
