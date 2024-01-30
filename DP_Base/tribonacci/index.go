package tribonacci

func tribonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	prev0, prev1, prev2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		current := prev0 + prev1 + prev2
		prev0, prev1, prev2 = prev1, prev2, current
	}

	return prev2
}
