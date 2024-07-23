package internal

func CalculatePoints(results []string) int {
	maxScore := 0
	for _, result := range results {
		switch result {
		case "w":
			maxScore += 3
		case "l":
			maxScore += 0
		case "d":
			maxScore += 1
		}
	}
	return maxScore
}
