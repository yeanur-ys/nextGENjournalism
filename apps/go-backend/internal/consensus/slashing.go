package consensus

func SlashReputation(current float64, penalty float64) float64 {
	next := current - penalty
	if next < 0 {
		return 0
	}
	return next
}
