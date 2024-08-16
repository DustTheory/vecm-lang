package compiler

import (
	"github.com/0xM-D/interpreter/module"
)

type Compiler struct {
	Modules     map[string]*module.Module

	EntryModule *module.Module
	Errors []CompilerError
}

func InitializeCompiler() (*Compiler, error) {
	return &Compiler{Modules: map[string]*module.Module{}}, nil
}

func (c *Compiler) LoadModule(moduleKey, code string) (*module.Module, bool) {
	module, failedToLoad := c.loadModule(moduleKey, code)
	return module, failedToLoad
}

func (c *Compiler) CompileModule(moduleKey string) (string, bool){
	module := c.Modules[moduleKey]
	ctx := c.compileProgram(module.Program)

	return ctx.Module.String(), c.hasCompilerErrors();
}