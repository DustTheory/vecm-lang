package parser_test

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestAssignmentDeclarationStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedName  string
		expectedValue string
	}{
		{"a := 10;", "a", "10"},
		{"b := true;", "b", "true"},
		{"c := fn(b: int) -> int { return b * 2 };", "c", "fn(b:int)->int{return (b * 2);}"},
		{"d := new map{int -> int}{1: 2, 2: 3};", "d", "new map{ int -> int }{1: 2, 2: 3}"},
		{"e := new []int{1, 2, 3, 4, 5};", "e", "new []int{1, 2, 3, 4, 5}"},
		{`f := "string value";`, "f", `"string value"`},
		{"fun := fn()->void {};", "fun", "fn()->void{}"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		decl, ok := program.Statements[0].(*ast.AssignmentDeclarationStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.AssignmentDeclarationStatement. got=%T", program.Statements[0])
		}

		if decl.Name.Value != tt.expectedName {
			t.Fatalf("decl.Name is not %q. got=%q", tt.expectedName, decl.Name.Value)
		}

		if decl.Value.String() != tt.expectedValue {
			t.Fatalf("decl.Value is not %s. got=%s", tt.expectedValue, decl.Value.String())
		}
	}
}
