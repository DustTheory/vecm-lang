package evaluator

import (
	"fmt"

	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/object"
)

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	if isTruthy(right) {
		return FALSE
	}
	return TRUE
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if !object.IsInteger(right) {
		return newError("unknown operator: -%s", right.Type().Signature())
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if object.IsError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	if object.IsNull(obj) {
		return false
	}

	if object.IsBoolean(obj) {
		if object.UnwrapReferenceObject(obj).(*object.Boolean).Value {
			return true
		} else {
			return false
		}
	}

	return true
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func evalIdentifier(
	node *ast.Identifier,
	env *object.Environment,
) object.Object {
	obj := env.Get(node.Value)
	if obj == nil {
		return newError("identifier not found: " + node.Value)
	}
	return obj
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if object.IsError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}
	return result

}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case object.IsArray(left) && object.IsInteger(index):
		return evalArrayIndexExpression(left, index)
	case object.IsHash(left):
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type().Signature())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)

	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalHashLiteral(
	node *ast.HashLiteral,
	env *object.Environment,
) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := object.UnwrapReferenceObject(Eval(keyNode, env))
		if object.IsError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type().Signature())
		}

		value := Eval(valueNode, env)
		if object.IsError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs, HashObjectType: object.HASH_OBJ()}
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type().Signature())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}

func evalTypedDeclarationStatement(node *ast.TypedDeclarationStatement, env *object.Environment) object.Object {
	objectType := evalType(node.Type, env)
	if objectType == nil {
		return newError("Unknown type: %s", node.Type.String())
	}

	error := declareVariable(&(*node).DeclarationStatement, objectType, env)
	if error != nil {
		return error
	}

	return nil
}

func evalType(typeNode ast.Type, env *object.Environment) object.ObjectType {
	switch casted := typeNode.(type) {
	case ast.HashType:
		return &object.HashObjectType{KeyType: evalType(casted.KeyType, env), ValueType: evalType(casted.ValueType, env)}
	case ast.ArrayType:
		return &object.ArrayObjectType{ElementType: evalType(casted.ElementType, env)}
	case ast.NamedType:
		namedType, _ := env.GetObjectType(casted.Token.Literal)
		return namedType
	}

	return nil
}

func declareVariable(declNode *ast.DeclarationStatement, expectedType object.ObjectType, env *object.Environment) object.Object {
	val := object.UnwrapReferenceObject(Eval(declNode.Value, env))

	if object.IsError(val) {
		return val
	}

	val = typeCast(val, expectedType, IMPLICIT_CAST)

	if object.IsError(val) {
		return val
	}

	if expectedType != nil && !object.TypesMatch(expectedType, val.Type()) {
		return newError("Expression of type %s cannot be assigned to %s", val.Type().Signature(), expectedType.Signature())
	}

	newObject := env.Declare(declNode.Name.Value, declNode.IsConstant, val)

	if newObject == nil {
		return newError("Identifier with name %s already exists.", declNode.Name.Value)
	}

	return nil
}
