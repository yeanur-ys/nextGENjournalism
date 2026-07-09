package ranking

import "math"

func CalculateReputation(votes float64, cVD float64, cSC float64, cF float64, w1 float64, w2 float64, w3 float64) float64 {
	return math.Log10(1+votes) + (w1 * cVD) + (w2 * cSC) - (w3 * cF)
}
