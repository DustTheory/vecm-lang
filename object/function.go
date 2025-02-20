package object

import (
	"bytes"
	"strings"

	"github.com/DustTheory/interpreter/ast"
)

type FunctionObjectType struct {
	ParameterTypes  []Type
	ReturnValueType Type
}

func (f FunctionObjectType) Signature() string {
	var out bytes.Buffer
	paramTypes := []string{}
	for _, p := range f.ParameterTypes {
		paramTypes = append(paramTypes, p.Signature())
	}

	out.WriteString("function(")
	out.WriteString(strings.Join(paramTypes, ", "))
	out.WriteString(") -> ")
	out.WriteString(f.ReturnValueType.Signature())

	return out.String()
}

func (f FunctionObjectType) Kind() Kind                    { return FunctionKind }
func (f FunctionObjectType) Builtins() *FunctionRepository { return FunctionKind.Builtins() }
func (f FunctionObjectType) IsConstant() bool              { return true }

type Function struct {
	FunctionObjectType
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type { return &f.FunctionObjectType }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if f.Body != nil {
		out.WriteString(" {\n")
		out.WriteString(f.Body.String())
		out.WriteString("\n}")
	} else {
		out.WriteString(";")
	}

	return out.String()
}
