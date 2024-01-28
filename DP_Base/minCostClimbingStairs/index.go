package minCostClimbingStairs

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}

	prev0, prev1 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		current := min(prev0, prev1) + cost[i]
		prev0, prev1 = prev1, current
	}

	return min(prev0, prev1)
}
