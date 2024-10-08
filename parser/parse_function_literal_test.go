package parser_test

import (
	"testing"

	"github.com/DustTheory/interpreter/ast"
	"github.com/DustTheory/interpreter/lexer"
	"github.com/DustTheory/interpreter/parser"
)

func TestFunctionLiteral(t *testing.T) {
	input := `fn(x: int, y: int) -> int { x + y; }`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T", stmt.Expression)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], TestIdentifier{"x"})
	testLiteralExpression(t, function.Parameters[1], TestIdentifier{"y"})

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n",
			len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T",
			function.Body.Statements[0])
	}

	testInfixExpression(t, bodyStmt.Expression, TestIdentifier{"x"}, "+", TestIdentifier{"y"})
}

func TestFunctionLiteralParameters(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []TestIdentifier
	}{
		{input: "fn()->void {};", expectedParams: []TestIdentifier{}},
		{input: "fn(x: int)->void {};", expectedParams: []TestIdentifier{{"x"}}},
		{
			input:          "fn(x: int, y: string, z: map{int -> string})->void {};",
			expectedParams: []TestIdentifier{{"x"}, {"y"}, {"z"}},
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt, isExpressionStatement := program.Statements[0].(*ast.ExpressionStatement)
		if !isExpressionStatement {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		function, isFunctionLiteral := stmt.Expression.(*ast.FunctionLiteral)
		if !isFunctionLiteral {
			t.Fatalf("stmt.Expression is not ast.FunctionLiteral. got=%T", stmt.Expression)
		}

		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want %d, got=%d\n",
				len(tt.expectedParams), len(function.Parameters))
		}

		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}
