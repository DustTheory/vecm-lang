package runtime

import (
	"fmt"
	"math"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/object"
)

func (r *Runtime) evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) (object.Object, error) {
	right, err := r.Eval(node.Right, env)
	if err != nil {
		return nil, err
	}

	switch node.Operator {
	case "!":
		return evalBangPrefixOperatorExpression(right)
	case "-":
		return r.evalMinusPrefixOperatorExpression(right)
	case "~":
		return r.evalTildePrefixOperatorExpression(right)
	default:
		return nil, fmt.Errorf("unknown operator: %s%s", node.Operator, right.Type())
	}
}

func evalBangPrefixOperatorExpression(right object.Object) (object.Object, error) {
	if isTruthy(right) {
		return False, nil
	}
	return True, nil
}

func (r *Runtime) evalMinusPrefixOperatorExpression(right object.Object) (object.Object, error) {
	number, isNumber := right.(*object.Number)
	if !isNumber {
		return nil, fmt.Errorf("operator - not defined on type %s", right.Type().Signature())
	}

	if object.IsInteger(number) && number.IsUnsigned() {
		return nil, fmt.Errorf("operator - not defined on unsigned integer type %s", number.Kind.Signature())
	}
	if object.IsInteger(number) && number.IsSigned() {
		return &object.Number{Value: object.Int64Bits(-number.GetInt64()), Kind: number.Kind}, nil
	}
	if number.Kind == object.Float32Kind {
		return &object.Number{Value: uint64(math.Float32bits(-number.GetFloat32())), Kind: number.Kind}, nil
	}
	if number.Kind == object.Float64Kind {
		return &object.Number{Value: math.Float64bits(-number.GetFloat64()), Kind: number.Kind}, nil
	}

	return nil, fmt.Errorf("operator - not defined on number type %s", right.Type().Signature())
}

func (r *Runtime) evalTildePrefixOperatorExpression(right object.Object) (object.Object, error) {
	number, isNumber := right.(*object.Number)

	if !isNumber {
		return nil, fmt.Errorf("operator ~ not defined on type %s", right.Type().Signature())
	}

	return &object.Number{Value: ^number.Value, Kind: number.Kind}, nil
}
