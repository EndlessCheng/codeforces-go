package main

func mincostTickets(days []int, costs []int) (ans int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	exp := [3]int{1, 7, 30}
	dp := make([]int, len(days))
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(p int) (res int) {
		dv := &dp[p]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for _i, e := range exp {
			costSum := costs[_i]
			i := p + 1
			for ; i < len(days) && days[i] < days[p]+e; i++ {
			}
			if i < len(days) {
				costSum += f(i)
			}
			res = min(res, costSum)
		}
		return
	}
	return f(0)
}
