// GaussianDistributionTests
package gotrueskill

import (
	"math"
	"testing"
)

var ErrorTolerance = 0.000001

func TestGetDrawMarginFromDrawProbability(t *testing.T) {

	beta := 25.0 / 6.0

	assertDrawMargin(0.10, beta, 0.74046637542690541, t)
	assertDrawMargin(0.25, beta, 1.87760059883033, t)
	assertDrawMargin(0.33, beta, 2.5111010132487492, t)
}

func assertDrawMargin(drawProbability, beta, expected float64, t *testing.T) {

	actual := GetDrawMarginFromDrawProbability(drawProbability, beta)

	result := FloatEqualTo(expected, actual, ErrorTolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", actual)
	}
}

func AssertFloatEqualTo(expected, actual, tolerance float64, t *testing.T) {
	result := FloatEqualTo(expected, actual, tolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", actual)
	}
}

func FloatEqualTo(num1, num2, tolerance float64) bool {

	diff := num1 - num2
	return math.Abs(diff) <= tolerance
}
