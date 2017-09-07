// Rating
package gotrueskill

import (
	"fmt"
)

const RankUpThreshold = 1
const RankDownThreshold = -4
const conservativeStandardDeviationMultiplierConst = 3

type Rating struct {
	ConservativeStandardDeviationMultiplier float64
	Mean                                    float64
	StandardDeviation                       float64
}

func NewRating(mean, standardDeviation float64) *Rating {

	return &Rating{conservativeStandardDeviationMultiplierConst, mean, standardDeviation}

}

func NewRatingWithMultiplier(mean, standardDeviation, conservativeStandardDeviationMultiplier float64) *Rating {

	return &Rating{conservativeStandardDeviationMultiplier, mean, standardDeviation}

}

func (r *Rating) ConservativeRating() float64 {
	return r.Mean - (r.ConservativeStandardDeviationMultiplier * r.StandardDeviation)
}

func GetPartialUpdate(prior, fullPosterior *Rating, updatePercentage float64) *Rating {

	priorGaussian := NewGaussianDistribution(prior.Mean, prior.StandardDeviation)
	posteriorGaussian := NewGaussianDistribution(fullPosterior.Mean, fullPosterior.StandardDeviation)

	precisionDifference := posteriorGaussian.Precision() - priorGaussian.Precision()
	partialPrecisionDifference := updatePercentage * precisionDifference

	precisionMeanDifference := posteriorGaussian.PrecisionMean() - priorGaussian.PrecisionMean()
	partialPrecisionMeanDifference := updatePercentage * precisionMeanDifference

	partialPosteriorGaussion := FromPrecisionMean(priorGaussian.PrecisionMean()+partialPrecisionMeanDifference, priorGaussian.Precision()+partialPrecisionDifference)

	return NewRatingWithMultiplier(partialPosteriorGaussion.Mean(), partialPosteriorGaussion.StandardDeviation(), prior.ConservativeStandardDeviationMultiplier)
}

func (r *Rating) PowerLevel() float64 {
	result := r.Mean - 2*r.StandardDeviation

	if result < 1 { // Make 1 the minimum, so it won't become negative
		result = 1
	}

	return result
}

// to string
func (r *Rating) String() string {
	return fmt.Sprintf("μ=%.4f, σ=%.4f", r.Mean, r.StandardDeviation)
}
