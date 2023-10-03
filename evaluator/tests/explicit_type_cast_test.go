package evaluator_tests

import (
	"math/big"
	"testing"
)

func TestExplicitTypeCast(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`123 as uint8`, big.NewInt(123)},
		{`[]int16{13, 14, 15} as []uint32`, []*big.Int{big.NewInt(13), big.NewInt(14), big.NewInt(15)}},
		{`123 as string`, "123"},
	}
	for _, tt := range tests {
		testLiteralObject(t, testEval(tt.input), tt.expected)
	}
}
