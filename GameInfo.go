// GameInfo
package trueskillgo

import (
	"math"
)

const defaultInitialMean = 25.0
const defaultBeta = defaultInitialMean / 6.0
const defaultDrawProbability = 0.10
const defaultDynamicsFactor = defaultInitialMean / 300.0
const defaultInitialStandardDeviation = defaultInitialMean / 3.0

const defaultXPGainedOnWin = 100
const defaultXPLostOnLoss = 50

var Default = DefaultGameInfo()

// Comparision between two players
const (
	Win  = 1
	Draw = 0
	Lose = -1
)

type GameInfo struct {
	InitialMean              float64
	InitialStandardDeviation float64
	Beta                     float64
	DynamicsFactor           float64
	DrawProbability          float64
	XPGainedOnWin            float64
	XPLostOnLoss             float64
}

func NewGameInfo(initialMean, initialStandardDeviation, beta, dynamicsFactor, drawProbability, xpGainedOnWin, xpLostOnLoss float64) *GameInfo {
	return &GameInfo{
		initialMean,
		initialStandardDeviation,
		beta, dynamicsFactor,
		drawProbability,
		xpGainedOnWin,
		xpLostOnLoss,
	}
}

func (g *GameInfo) DefaultRating() *Rating {
	return NewRating(g.InitialMean, g.InitialStandardDeviation)
}

func DefaultGameInfo() *GameInfo {

	return &GameInfo{defaultInitialMean, defaultInitialStandardDeviation, defaultBeta, defaultDynamicsFactor, defaultDrawProbability, defaultXPGainedOnWin, defaultXPLostOnLoss}

}

func GetDrawMarginFromDrawProbability(drawProbability, beta float64) float64 {

	var margin = InverseCumulativeTo(0.5*(drawProbability+1), 0, 1) * math.Sqrt(1+1) * beta

	return margin
}
