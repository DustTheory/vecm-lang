package evaluator_tests

import (
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"1 << 1", 2},
		{"1 << 62", 1 << 62},
		{"1 << 63 >> 1", -(1 << 63 >> 1)},
		{"1 << 64 >> 2", 0},
		{"1 >> 1", 0},
		{"256 >> 2", 64},
		{"3 >> 1", 1},
		{"3 << 1", 6},
		{"1 | 3", 3},
		{"4097 | 272", 4369},
		{"0 | 272", 272},
		{"0 | 0", 0},
		{"0 & 0", 0},
		{"0 & 1", 0},
		{"4097 & 272", 0},
		{"7 & 3", 3},
		{"~0", ^0},
		{"((1 << 10) - 1) ^ (1 << 8)", ((1 << 10) - 1) ^ (1 << 8)},
		{"~123", ^123},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}
