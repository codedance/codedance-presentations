
func FindBest(solutions []*Solution) *Solution {
	var best *Solution
	for _, solution := range solutions {
		if solution.Complete {
			if best == nil || solution.Cost < best.Cost {
				best := solution // HL
			}
		}
	}
	return best
}
