package calculator_test

import (
	"testing"

	"github.com/Dmitrylolo/golang-tests/calculator"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalcualteIsAmstrong(t *testing.T) {
	testCase := TestCase{
		value:    371,
		expected: true,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

func TestNegativeCalculateIsAmstrong(t *testing.T) {
	testCase := TestCase{
		value:    350,
		expected: false,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
