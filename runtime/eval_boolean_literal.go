package runtime

import (
	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalBooleanLiteral(node *ast.BooleanLiteral) *object.Boolean {
	return nativeBoolToBooleanObject(node.Value)
}
