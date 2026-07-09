package consensus

type Vote struct {
	AuditorID string
	ClaimID   string
	Weight    float64
	Decision  bool
}

func WeightedScore(votes []Vote) float64 {
	var score float64
	for _, vote := range votes {
		if vote.Decision {
			score += vote.Weight
		} else {
			score -= vote.Weight
		}
	}
	return score
}
