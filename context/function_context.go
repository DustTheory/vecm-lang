package context

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/util"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

type FunctionContext struct {
	sharedContextProperties SharedContextProperties
	functionParams VariableStore
	*ir.Func
}

func (ctx *FunctionContext) GetParentContext() Context {
	return ctx.sharedContextProperties.parentContext;
}

func NewFunctionContext(parent Context, fn *ir.Func, parameterNames []*ast.Identifier, parameterTypes []ast.Type) *FunctionContext {
	ctx := &FunctionContext{
		sharedContextProperties: SharedContextProperties{parentContext: parent},
		Func: fn,
	}

	ctx.functionParams = VariableStore{variables: map[string]Variable{}}
	for i, name := range parameterNames {
		t, error := util.GetLLVMType(parameterTypes[i])
		if error != nil {
			panic(error)
		}
		ctx.functionParams.DeclareVariable(name.Value, t, fn.Params[i])
	}

	return ctx
}

func (ctx *FunctionContext) LookUpIdentifier(name string) (Variable, bool) {
	variable, ok := ctx.functionParams.LookUpVariable(name)
	if ok {
		return variable, ok
	}

	if ctx.GetParentContext() == nil {
		return nil, false
	}

	return ctx.GetParentContext().LookUpIdentifier(name)
}

func (ctx *FunctionContext) GetFunction(signature types.FuncType) (*ir.Func, bool) {
	return ctx.GetParentContext().GetFunction(signature)
}

func (ctx *FunctionContext) DeclareFunction(name string, retType types.Type, params ...*ir.Param) *ir.Func {
	return ctx.GetParentContext().DeclareFunction(name, retType, params...)
}

func (ctx *FunctionContext) DeclareLocalVariable(name string, t types.Type) *ir.InstAlloca {
	return nil // TODO: Throw error
}