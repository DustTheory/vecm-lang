package compiler

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/context"
)

func (c *Compiler) compileIfStatement(expr *ast.IfStatement, block *context.BlockContext) *context.BlockContext {
	condition := c.compileExpression(expr.Condition, block)

	var consequenceBlock *context.BlockContext
	var alternativeBlock *context.BlockContext
	var continueBlock *context.BlockContext

	parentContext, err := block.GetParentContext()
	if err != nil {
		c.newCompilerError(expr, "%e", err)
		return nil
	}

	consequenceBlock = context.NewBlockContext(parentContext, block.Parent.NewBlock(""))
	c.compileBlockStatement(expr.Consequence, consequenceBlock)

	if expr.Alternative != nil {
		alternativeBlock = context.NewBlockContext(parentContext, block.Parent.NewBlock(""))
		c.compileBlockStatement(expr.Alternative, alternativeBlock)
	}

	consequenceBlockHasTerm := consequenceBlock != nil && consequenceBlock.Block.Term != nil
	alternativeBlockHasTerm := alternativeBlock != nil && alternativeBlock.Block.Term != nil

	if consequenceBlockHasTerm && alternativeBlockHasTerm {
		block.NewCondBr(condition, consequenceBlock.Block, alternativeBlock.Block)
	} else {
		continueBlock = context.NewBlockContext(parentContext, block.Parent.NewBlock(""))
		block.NewCondBr(condition, consequenceBlock.Block, continueBlock.Block)

		if !consequenceBlockHasTerm && consequenceBlock != nil {
			consequenceBlock.NewBr(continueBlock.Block)
		}

		if !alternativeBlockHasTerm && alternativeBlock != nil {
			alternativeBlock.NewBr(continueBlock.Block)
		}
	}

	return continueBlock
}
