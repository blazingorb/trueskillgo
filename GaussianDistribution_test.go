// GaussianDistributionTests
package gotrueskill

import (
	"blazingorb/trueskill/TestUtil"
	"testing"
)

func TestCumulativeTo(t *testing.T) {
	ErrorTolerance := 0.000001

	t.Log("Testing cumulative to")
	num := CumulativeToX(0.5)
	expected := 0.691462
	result := TestUtil.EqualTo(expected, num, ErrorTolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", num)
	}
}

func TestAtX(t *testing.T) {

	t.Log("Testing AtX")
	num := AtX(0.5)
	expected := 0.352065
	result := TestUtil.EqualTo(expected, num, ErrorTolerance)

	if !result {
		t.Error("Expected Result: ", expected, "\t Actual Result: ", num)
	}
}
