package context

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type Context interface {
	GetParentContext() *Context
	GetFunction(name string, params ...*ir.Param) *ir.Func
	DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func
}

type SharedContextProperties struct {	
	parentContext *Context
	// functionStore FunctionStore
}