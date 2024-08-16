package compiler

import (
	"testing"
)

func TestEmptyFunctionDeclaration(t *testing.T) {

	t.Fatalf("No statements in function declaration crashes the compiler")
	// code := `fn main() -> void {}`

	// module := compileAndVerifyCode(code, t)

	// if len(module.Funcs) != 1 {
	// 	t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	// }

	// fn := module.Funcs[0]

	// if len(fn.Blocks) != 1 {
	// 	t.Fatalf("Expected 1 block, got %d", len(fn.Blocks))
	// }

	// block := fn.Blocks[0]

	// if len(block.Insts) != 0 {
	// 	t.Fatalf("Expected 0 instructions, got %d", len(block.Insts))
	// }

}

func TestFunctionDeclarationWithIntegerReturn(t *testing.T) {
	code := `fn main() -> int { return 5; }`

	module := compileAndVerifyCode(code, t)

	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	if len(fn.Blocks) != 1 {
		t.Fatalf("Expected 1 block, got %d", len(fn.Blocks))
	}

	block := fn.Blocks[0]

	if len(block.Insts) != 1 {
		t.Fatalf("Expected 1 instruction, got %d", len(block.Insts))
	}

	inst := block.Insts[0]

	if inst.LLString() != "ret i32 5" {
		t.Fatalf("Expected ret i32 5, got %s", inst.LLString())
	}

}
func TestFunctionDeclarationWithNoParameters(t *testing.T) {
	t.Fatalf("No statements in function declaration crashes the compiler")

	// code := `fn main() -> void {}`

	// module := compileAndVerifyCode(code, t)

	// if len(module.Funcs) != 1 {
	// 	t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	// }

	// fn := module.Funcs[0]

	// if len(fn.Params) != 0 {
	// 	t.Fatalf("Expected 0 parameters, got %d", len(fn.Params))
	// }
}

func TestFunctionDeclarationWithOneParameter(t *testing.T) {
	code := `fn add(x: int) -> int { return x + 5; }`

	module := compileAndVerifyCode(code, t)

	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	if len(fn.Params) != 1 {
		t.Fatalf("Expected 1 parameter, got %d", len(fn.Params))
	}

	param := fn.Params[0]

	if param.Type().LLString() != "i32" {
		t.Fatalf("Expected parameter type i32, got %s", param.Typ.LLString())
	}
}

func TestFunctionDeclarationWithMultipleParameters(t *testing.T) {
	code := `fn add(x: int, y: int) -> int { return x + y; }`

	module := compileAndVerifyCode(code, t)

	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	if len(fn.Params) != 2 {
		t.Fatalf("Expected 2 parameters, got %d", len(fn.Params))
	}

	param1 := fn.Params[0]
	param2 := fn.Params[1]

	if param1.Typ.LLString() != "i32" {
		t.Fatalf("Expected parameter 1 type i32, got %s", param1.Typ.LLString())
	}

	if param2.Typ.LLString() != "i32" {
		t.Fatalf("Expected parameter 2 type i32, got %s", param2.Typ.LLString())
	}
}

func TestFunctionDeclarationWithDifferentReturnTypes(t *testing.T) {
	code := `fn getBool() -> bool { return true; }`

	module := compileAndVerifyCode(code, t)

	if len(module.Funcs) != 1 {
		t.Fatalf("Expected 1 function, got %d", len(module.Funcs))
	}

	fn := module.Funcs[0]

	if fn.Typ.LLString() != "i1" {
		t.Fatalf("Expected return type i1, got %s", fn.Typ.LLString())
	}
}