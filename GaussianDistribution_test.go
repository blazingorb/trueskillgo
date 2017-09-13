// GaussianDistributionTests
package trueskillgo

import (
	"math"
	"testing"
)

func TestCumulativeTo(t *testing.T) {
	ErrorTolerance := 0.000001

	t.Log("Testing cumulative to")
	num := CumulativeToX(0.5)
	expected := 0.691462
	result := EqualTo(expected, num, ErrorTolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", num)
	}
}

func TestAtX(t *testing.T) {

	t.Log("Testing AtX")
	num := AtX(0.5)
	expected := 0.352065
	result := EqualTo(expected, num, ErrorTolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", num)
	}
}

func AssertEqualTo(expected, actual, tolerance float64, t *testing.T) {
	result := EqualTo(expected, actual, tolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", actual)
	}
}

func EqualTo(num1, num2, tolerance float64) bool {

	diff := num1 - num2
	return math.Abs(diff) <= tolerance
}
