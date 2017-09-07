// GaussianDistributionTests
package gotrueskill

import (
	"testing"
)

var RatingErrorTolerance = 0.085

func TestMassiveUpsetDraw(t *testing.T) {
	t.Log("Testing cumulative to")

	game := DefaultGameInfo()
	p1 := game.DefaultRating()
	p2 := NewRating(50, 12.5)

	p1Temp, p2Temp := CalculateNewRatings2(game, p1, p2, Draw)
	matchQuality := CalculateMatchQuality(game, p1, p2)

	expected1 := NewRating(31.662, 7.137)
	result := RatingEqualTo(expected1, p1Temp)

	if !result {
		t.Error("Expected Result: ", expected1, "\t Actual Result: ", p1Temp)
	}

	expected2 := NewRating(35.010, 7.910)
	result = RatingEqualTo(expected2, p2Temp)

	if !result {
		t.Error("Expected Result: ", expected2, "\t Actual Result: ", p2Temp)
	}

	expectedMatchQuality := 0.110
	resultMatchQuality := FloatEqualTo(expectedMatchQuality, matchQuality, RatingErrorTolerance)
	if !resultMatchQuality {
		t.Error("Expected Match Quality: ", expectedMatchQuality, "\t Actual Result: ", matchQuality)
	}

}

func RatingEqualTo(expectedRating, rating *Rating) bool {
	return FloatEqualTo(expectedRating.Mean, rating.Mean, RatingErrorTolerance) && FloatEqualTo(expectedRating.StandardDeviation, rating.StandardDeviation, RatingErrorTolerance)
}
