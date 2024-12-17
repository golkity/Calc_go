package calc

import (
	"calc_go/internal/Errors"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name         string
		expression   string
		expected     float64
		expecttedERR error
	}{
		//ЧЕТКИЕ ТЕСТЫ
		{"Simple addition", "2+2", 4, nil},
		{"Simple subtraction", "5-3", 2, nil},
		{"Simple multiplication", "4*3", 12, nil},
		{"Simple division", "10/2", 5, nil},
		{"Complex expression", "2+3*4", 14, nil},
		{"Expression with brackets", "(2+3)*4", 20, nil},
		{"Nested brackets", "((2+3)*4)+10", 30, nil},

		//НЕЧЕТКИЕ ТЕСТЫ
		{"Invalid character", "2+a", 0, Errors.ErrInvalidInput},
		{"Division by zero", "4/0", 0, Errors.ErrDivisionByZero},
		{"Mismatched brackets", "2+(3*4", 0, Errors.ErrMismatchedBrackets},
		{"Empty expression", "", 0, Errors.ErrEmptyExpression},
		{"Invalid syntax", "+2*3", 0, Errors.ErrInvalidExpression},
		{"Invalid syntax at end", "2*3+", 0, Errors.ErrInvalidExpression},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calc(tt.expression)
			if result != tt.expected {
				t.Errorf("Calc() expected %v, actual %v", tt.expected, result)
			}
			if err != tt.expecttedERR {
				if err == nil || err.Error() != tt.expecttedERR.Error() {
					t.Errorf("Calc() expected %v, actual %v", tt.expecttedERR, err)
				}
			}
		})
	}
}

func TestIsSign(t *testing.T) {
	tests := []struct {
		name     string
		expr     rune
		expected bool
	}{
		{"Addition sign", '+', true},
		{"Subtraction sign", '-', true},
		{"Multiplication sign", '*', true},
		{"Division sign", '/', true},
		{"Invalid sign", 'a', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSign(tt.expr); got != tt.expected {
				t.Errorf("IsSign() expected %v, actual %v", tt.expected, got)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		expected float64
	}{
		{"Positive integer", "123", 123},
		{"Negative integer", "-123", -123},
		{"Zero", "0", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToFloat64(tt.expr)
			if result != tt.expected {
				t.Errorf("StringToFloat64() expected %v, actual %v", tt.expected, result)
			}
		})
	}
}
