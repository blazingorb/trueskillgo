// GaussianDistribution
package gotrueskill

import (
	"math"
)

var coefficients = []float64{-1.3026537197817094, 6.4196979235649026e-1,
	1.9476473204185836e-2, -9.561514786808631e-3, -9.46595344482036e-4,
	3.66839497852761e-4, 4.2523324806907e-5, -2.0278578112534e-5,
	-1.624290004647e-6, 1.303655835580e-6, 1.5626441722e-8, -8.5238095915e-8,
	6.529054439e-9, 5.059343495e-9, -9.91364156e-10, -2.27365122e-10,
	9.6467911e-11, 2.394038e-12, -6.886027e-12, 8.94487e-13, 3.13092e-13,
	-1.12708e-13, 3.81e-16, 7.106e-15, -1.523e-15, -9.4e-17, 1.21e-16, -2.8e-17}

type GaussianDistribution struct {
	mean              float64
	standardDeviation float64
	Variance          float64
	precision         float64
	precisionMean     float64
}

func NewGaussianDistribution(mean, standardDeviation float64) *GaussianDistribution {

	var variance = standardDeviation * standardDeviation
	var precision = 1.0 / variance
	var precisionMean = precision * mean

	return &GaussianDistribution{mean, standardDeviation, variance, precision, precisionMean}
}

func (g *GaussianDistribution) Mean() float64 {
	return g.mean
}

func (g *GaussianDistribution) StandardDeviation() float64 {
	return g.standardDeviation
}

func (g *GaussianDistribution) Precision() float64 {
	return g.precision
}

func (g *GaussianDistribution) PrecisionMean() float64 {
	return g.precisionMean
}

func (g *GaussianDistribution) NormalizationConstant() float64 {
	return 1.0 / (math.Sqrt(2*math.Pi) * g.standardDeviation)
}

func (g *GaussianDistribution) Clone() *GaussianDistribution {
	return &GaussianDistribution{g.mean, g.standardDeviation, g.Variance, g.precision, g.precisionMean}
}

func FromPrecisionMean(precisionMean, precision float64) *GaussianDistribution {

	var variance = 1.0 / precision
	var standardDeviation = math.Sqrt(variance)
	var mean = precisionMean / precision

	return &GaussianDistribution{mean, standardDeviation, variance, precision, precisionMean}
}

func Multiply(left, right *GaussianDistribution) *GaussianDistribution {

	return FromPrecisionMean(left.precisionMean+right.precisionMean, left.precision+right.precision)
}

// AbsoluteDifference
// Minus
// LogProductNormalization
// Divide

// LogRatioNormalization

// At
func AtX(x float64) float64 {
	return At(x, 0, 1)
}

func At(x, mean, standardDeviation float64) float64 {
	multiplier := 1.0 / (standardDeviation * math.Sqrt(2*math.Pi))
	expPart := math.Exp((-1.0 * math.Pow(x-mean, 2.0)) / (2 * (standardDeviation * standardDeviation)))

	return multiplier * expPart
}

// CumulativeTo
func CumulativeTo(x, mean, standardDeviation float64) float64 {

	invsqrt2 := -0.707106781186547524400844362104
	result := ErrorFunctionCumulativeTo(invsqrt2 * x)

	return 0.5 * result
}

func CumulativeToX(x float64) float64 {
	return CumulativeTo(x, 0, 1)
}

func ErrorFunctionCumulativeTo(x float64) float64 {

	z := math.Abs(x)

	t := 2.0 / (2.0 + z)
	ty := 4*t - 2

	ncof := len(coefficients)
	d := 0.0
	dd := 0.0

	for j := ncof - 1; j > 0; j-- {
		tmp := d
		d = ty*d - dd + coefficients[j]
		dd = tmp
	}

	ans := t * math.Exp(-z*z+0.5*(coefficients[0]+ty*d)-dd)

	if x < 0.0 {
		ans = 2.0 - ans
	}

	return ans
}

func InverseErrorFunctionCumulativeTo(p float64) float64 {
	if p >= 2.0 {
		return -100
	} else if p <= 0.0 {
		return 100
	}

	var pp = p
	if !(p < 1.0) {
		pp = 2 - p
	}
	var t = math.Sqrt(-2 * math.Log(pp/2.0)) // initial guess
	var x = -0.70711 * ((2.30753+t*0.27061)/(1.0+t*(0.99229+t*0.04481)) - t)

	for j := 0; j < 2; j++ {

		err := ErrorFunctionCumulativeTo(x) - pp
		x += err / (1.12837916709551257*math.Exp(-(x*x)) - x*err)
	}

	if p < 1.0 {
		return x
	} else {
		return -x
	}

}

func InverseCumulativeTo(x, mean, standardDeviation float64) float64 {

	return mean - math.Sqrt(2)*standardDeviation*InverseErrorFunctionCumulativeTo(2*x)

}
