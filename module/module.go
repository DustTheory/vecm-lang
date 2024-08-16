package module

import (
	"github.com/0xM-D/interpreter/ast"
	"github.com/0xM-D/interpreter/lexer"
	"github.com/0xM-D/interpreter/object"
	"github.com/0xM-D/interpreter/parser"
)

type Module struct {
	ModuleKey       string
	RootEnvironment object.Environment
	Program         *ast.Program
}

func ParseModule(moduleKey, code string) (*Module, []string) {
	l := lexer.New(string(code))
	p := parser.New(l)

	program := p.ParseProgram()

	module := &Module{
		ModuleKey:       moduleKey,
		RootEnvironment: *object.NewEnvironment(),
		Program:         program,
	}

	return module, p.Errors()
}