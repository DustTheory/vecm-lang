package runtime_test

import (
	"math/big"
	"testing"

	"github.com/DustTheory/interpreter/object"
)

func TestArrayLiterals(t *testing.T) {
	input := "new []int{1, 2 * 2, 3 + 3}"

	evaluated, err := testEval(input)
	if err != nil {
		t.Fatal(err)
	}

	result, isResultArray := evaluated.(*object.Array)

	if !isResultArray {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], big.NewInt(1))
	testIntegerObject(t, result.Elements[1], big.NewInt(4))
	testIntegerObject(t, result.Elements[2], big.NewInt(6))
}

func TestHashLiterals(t *testing.T) {
	input := `let two = "two";
	new map{string -> int }{
	"one": 10 - 9,
	two: 1 + 1,
	"thr" + "ee": 6 / 2
	}`

	evaluated, err := testEval(input)
	if err != nil {
		t.Fatal(err)
	}

	result, isResultHash := evaluated.(*object.Hash)
	if !isResultHash {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
	}

	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}

	for expectedKey, expectedValue := range expected {
		pair, keyExists := result.Pairs[expectedKey]
		if !keyExists {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, big.NewInt(expectedValue))
	}
}
