package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) (object.Object, error) {
	functionType, err := r.evalType(node.Type, env)
	if err != nil {
		return nil, err
	}

	function := &object.Function{
		Parameters:         node.Parameters,
		Env:                env,
		Body:               node.Body,
		FunctionObjectType: *functionType.(*object.FunctionObjectType),
	}

	return function, nil
}
