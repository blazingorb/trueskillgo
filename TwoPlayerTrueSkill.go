// TwoPlayerTrueSkill
package trueskillgo

import (
	"math"
)

// public CalculateNewRatings
func CalculateNewRatings2(gameInfo *GameInfo, selfRating, opponentRating *Rating, comparison int) (*Rating, *Rating) {
	selfRatingNew := CalculateNewRatings(gameInfo, selfRating, opponentRating, comparison)
	oppComparison := comparison

	if oppComparison == Win {
		oppComparison = Lose
	} else if oppComparison == Lose {
		oppComparison = Win
	}

	oppRatingNew := CalculateNewRatings(gameInfo, opponentRating, selfRating, oppComparison)
	return selfRatingNew, oppRatingNew
}

func CalculateXPGained2(gameInfo *GameInfo, selfRating, opponentRating *Rating, comparison int) (int32, int32) {
	selfXP := CalculateXPGained(gameInfo, selfRating, opponentRating, comparison)

	oppComparison := comparison

	if oppComparison == Win {
		oppComparison = Lose
	} else if oppComparison == Lose {
		oppComparison = Win
	}

	oppXP := CalculateXPGained(gameInfo, opponentRating, selfRating, oppComparison)
	return selfXP, oppXP
}

func CalculateXPGained(gameInfo *GameInfo, selfRating, opponentRating *Rating, comparison int) int32 {
	winningPower := selfRating.PowerLevel()
	losingPower := opponentRating.PowerLevel()
	xp := gameInfo.XPGainedOnWin

	switch comparison {
	case Win, Draw:
		break
	case Lose:
		winningPower = opponentRating.PowerLevel()
		losingPower = selfRating.PowerLevel()
		xp = gameInfo.XPLostOnLoss
	}

	return int32(float64(comparison) * (xp * (losingPower / winningPower)))
}

func CalculateNewRatings(gameInfo *GameInfo, selfRating, opponentRating *Rating, comparison int) *Rating {
	var drawMargin = GetDrawMarginFromDrawProbability(gameInfo.DrawProbability, gameInfo.Beta)

	c := math.Sqrt(square(selfRating.StandardDeviation) + square(opponentRating.StandardDeviation) + 2*square(gameInfo.Beta))
	winningMean := selfRating.Mean
	losingMean := opponentRating.Mean

	switch comparison {

	case Win, Draw:
		break
	case Lose:
		winningMean = opponentRating.Mean
		losingMean = selfRating.Mean
		break
	default:
	}

	meanDelta := winningMean - losingMean
	var v, w, rankMultiplier = 0.0, 0.0, 0

	if comparison != Draw {
		v = VExceedsMarginWithC(meanDelta, drawMargin, c)
		w = WExceedsMarginWithC(meanDelta, drawMargin, c)
		rankMultiplier = comparison
	} else {
		v = VWithinMarginWithC(meanDelta, drawMargin, c)
		w = WWithinMarginWithC(meanDelta, drawMargin, c)
		rankMultiplier = 1
	}

	meanMultiplier := (square(selfRating.StandardDeviation) + square(gameInfo.DynamicsFactor)) / c

	varianceWithDynamics := square(selfRating.StandardDeviation) + square(gameInfo.DynamicsFactor)
	stdDevMultiplier := varianceWithDynamics / square(c)

	newMean := selfRating.Mean + (float64(rankMultiplier) * meanMultiplier * v)
	newStdDev := math.Sqrt(varianceWithDynamics * (1 - w*stdDevMultiplier))

	return NewRating(newMean, newStdDev)
}

func CalculateMatchQuality(g *GameInfo, player1, player2 *Rating) float64 {
	// Use equation 4.1 from pg 8 of 2006 paper
	betaSquared := square(g.Beta)
	player1SigmaSquared := square(player1.StandardDeviation)
	player2SigmaSquared := square(player2.StandardDeviation)

	sqrtPart := math.Sqrt((2 * betaSquared) / (2*betaSquared + player1SigmaSquared + player2SigmaSquared))

	expPart := math.Exp((-1 * square(player1.Mean-player2.Mean)) / (2 * (2*betaSquared + player1SigmaSquared + player2SigmaSquared)))

	return sqrtPart * expPart
}

func square(num float64) float64 {
	return num * num
}
