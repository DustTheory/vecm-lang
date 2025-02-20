package parser_test

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestFunctionDeclarationStatement(t *testing.T) {
	input := `fn functionName(param1: int, param2: string) -> int { return param1 + param2 }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d", 1, len(program.Statements))
	}

	functionDeclaration, ok := program.Statements[0].(*ast.FunctionDeclarationStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.FunctionDeclarationStatement. got=%T", program.Statements[0])
	}

	if functionDeclaration.Name.Value != "functionName" {
		t.Fatalf("function name wrong. want=%s, got=%s", "functionName", functionDeclaration.Name.Value)
	}

	if len(functionDeclaration.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(functionDeclaration.Parameters))
	}

	testLiteralExpression(t, functionDeclaration.Parameters[0], TestIdentifier{"param1"})
	testLiteralExpression(t, functionDeclaration.Parameters[1], TestIdentifier{"param2"})

	if len(functionDeclaration.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n",
			len(functionDeclaration.Body.Statements))
	}

	bodyStmt, ok := functionDeclaration.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ReturnStatement. got=%T",
			functionDeclaration.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.ReturnValue, TestIdentifier{"param1"}, "+", TestIdentifier{"param2"})
}

func TestExternalFunctionDeclaration(t *testing.T) {
	input := `fn functionName(param1: int, param2: string) -> int;`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	functionDeclaration, ok := program.Statements[0].(*ast.FunctionDeclarationStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.FunctionDeclarationStatement. got=%T", program.Statements[0])
	}

	if functionDeclaration.Name.Value != "functionName" {
		t.Fatalf("function name wrong. want=%s, got=%s", "functionName", functionDeclaration.Name.Value)
	}

	if len(functionDeclaration.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(functionDeclaration.Parameters))
	}

	testLiteralExpression(t, functionDeclaration.Parameters[0], TestIdentifier{"param1"})
	testLiteralExpression(t, functionDeclaration.Parameters[1], TestIdentifier{"param2"})

	if functionDeclaration.Body != nil {
		t.Fatalf("function.Body should not exist for external functions. got=%+v\n", functionDeclaration.Body)
	}
}

func TestVariadicFunctionDeclaration(t *testing.T) {
	input := `fn functionName(param1: int, ...) -> int {
		return param1 +param2 + param3
	}`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d", 1, len(program.Statements))
	}

	functionDeclaration, ok := program.Statements[0].(*ast.FunctionDeclarationStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.FunctionDeclarationStatement. got=%T", program.Statements[0])
	}

	if functionDeclaration.Name.Value != "functionName" {
		t.Fatalf("function name wrong. want=%s, got=%s", "functionName", functionDeclaration.Name.Value)
	}

	if len(functionDeclaration.Parameters) != 1 {
		t.Fatalf("function parameters wrong. want 1, got=%d\n", len(functionDeclaration.Parameters))
	}

	if !functionDeclaration.IsVariadic {
		t.Fatalf("function is not variadic. got=%t\n", functionDeclaration.IsVariadic)
	}

	testLiteralExpression(t, functionDeclaration.Parameters[0], TestIdentifier{"param1"})
}
