package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileStatement(stmt ast.Statement, b *context.BlockContext) *context.BlockContext {
	switch stmt := stmt.(type) {
	case *ast.ReturnStatement:
		return c.compileReturnStatement(stmt, b)
	case *ast.IfStatement:
		return c.compileIfStatement(stmt, b)
	case *ast.BlockStatement:
		parentContext, err := b.GetParentContext()
		if err != nil {
			c.newCompilerError(stmt, "%e", err)
			return nil
		}
		parentFunctionContext, err := b.GetParentFunctionContext()
		if err != nil {
			c.newCompilerError(stmt, "%e", err)
			return nil
		}

		newBlock := context.NewBlockContext(parentContext, parentFunctionContext.NewBlock(""))
		return c.compileBlockStatement(stmt, newBlock)
	case *ast.ExpressionStatement:
		c.compileExpression(stmt.Expression, b)
		return b
	case *ast.LetStatement:
		c.compileLetStatement(stmt, b)
		return b
	// case *ast.FunctionStatement: // OH BOY I AM NOT WRITING THIS TODAY
	// 	c.compileFunctionStatement(stmt, b)
	// 	return b
	case *ast.ForStatement:
		return c.compileForStatement(stmt, b)
	case *ast.AssignmentDeclarationStatement:
		return c.compileDeclarationStatement(&stmt.DeclarationStatement, b)
	case *ast.DeclarationStatement:
		return c.compileDeclarationStatement(stmt, b)
	case *ast.TypedDeclarationStatement:
		return c.compileDeclarationStatement(&stmt.DeclarationStatement, b)
	default:
		c.newCompilerError(stmt, "Unknown statement type: %T", stmt)
		return b
	}
}
