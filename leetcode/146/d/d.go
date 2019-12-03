package main

func maxAbsValExpr(a1 []int, a2 []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	const inf int = 1e9
	mins := [4]int{inf, inf, inf, inf}
	maxs := [4]int{-inf, -inf, -inf, -inf}
	for i := range a1 {
		for j, v := range [4]int{
			a1[i] + a2[i] + i,
			a1[i] - a2[i] + i,
			-a1[i] + a2[i] + i,
			-a1[i] - a2[i] + i,
		} {
			mins[j] = min(mins[j], v)
			maxs[j] = max(maxs[j], v)
		}
	}
	ans := -inf
	for i := range mins {
		ans = max(ans, maxs[i]-mins[i])
	}
	return ans
}
