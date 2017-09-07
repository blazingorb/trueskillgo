// TruncatedGaussianCorrectionFunctions
package gotrueskill

import (
	"math"
)

func VExceedsMarginWithC(teamPerformanceDifference, drawMargin, c float64) float64 {

	return VExceedsMargin(teamPerformanceDifference/c, drawMargin/c)
}

func VExceedsMargin(teamPerformanceDifference, drawMargin float64) float64 {

	denominator := CumulativeToX(teamPerformanceDifference - drawMargin)

	if denominator < 2.222758749e-162 {
		return -teamPerformanceDifference + drawMargin
	}

	return AtX(teamPerformanceDifference-drawMargin) / denominator
}

func WExceedsMarginWithC(teamPerformanceDifference, drawMargin, c float64) float64 {

	return WExceedsMargin(teamPerformanceDifference/c, drawMargin/c)
}

func WExceedsMargin(teamPerformanceDifference, drawMargin float64) float64 {

	denominator := CumulativeToX(teamPerformanceDifference - drawMargin)

	if denominator < 2.222758749e-162 {
		if teamPerformanceDifference < 0.0 {
			return 1.0
		}
		return 0
	}

	vWin := VExceedsMargin(teamPerformanceDifference, drawMargin)
	return vWin * (vWin + teamPerformanceDifference - drawMargin)
}

//
func VWithinMarginWithC(teamPerformanceDifference, drawMargin, c float64) float64 {

	return VWithinMargin(teamPerformanceDifference/c, drawMargin/c)
}

func VWithinMargin(teamPerformanceDifference, drawMargin float64) float64 {

	teamPerformanceDifferenceAbsoluteValue := math.Abs(teamPerformanceDifference)
	denominator := CumulativeToX(drawMargin-teamPerformanceDifferenceAbsoluteValue) - CumulativeToX(-drawMargin-teamPerformanceDifferenceAbsoluteValue)

	if denominator < 2.222758749e-162 {
		if teamPerformanceDifference < 0.0 {
			return -teamPerformanceDifference - drawMargin
		}
		return -teamPerformanceDifference + drawMargin
	}

	numerator := AtX(-drawMargin-teamPerformanceDifferenceAbsoluteValue) - AtX(drawMargin-teamPerformanceDifferenceAbsoluteValue)

	if teamPerformanceDifference < 0.0 {
		return -numerator / denominator
	}

	return numerator / denominator
}

func WWithinMarginWithC(teamPerformanceDifference, drawMargin, c float64) float64 {

	return WWithinMargin(teamPerformanceDifference/c, drawMargin/c)
}

func WWithinMargin(teamPerformanceDifference, drawMargin float64) float64 {

	teamPerformanceDifferenceAbsoluteValue := math.Abs(teamPerformanceDifference)
	denominator := CumulativeToX(drawMargin-teamPerformanceDifferenceAbsoluteValue) - CumulativeToX(-drawMargin-teamPerformanceDifferenceAbsoluteValue)

	if denominator < 2.222758749e-162 {
		return 1.0

	}

	vt := VWithinMargin(teamPerformanceDifferenceAbsoluteValue, drawMargin)
	temp := ((drawMargin-teamPerformanceDifferenceAbsoluteValue)*AtX(drawMargin-teamPerformanceDifferenceAbsoluteValue) - (-drawMargin-teamPerformanceDifferenceAbsoluteValue)*AtX(-drawMargin-teamPerformanceDifferenceAbsoluteValue)) / denominator

	return vt*vt + temp
}
