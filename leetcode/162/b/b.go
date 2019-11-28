package main

func reconstructMatrix(upper int, lower int, colsum []int) (ans [][]int) {
	ans = [][]int{}
	sum := 0
	for _, s := range colsum {
		sum += s
	}
	if sum != upper+lower {
		return
	}
	n := len(colsum)
	ans = [][]int{make([]int, n), make([]int, n)}
	// fill 2
	for i, sum := range colsum {
		if sum == 2 {
			ans[0][i] = 1
			ans[1][i] = 1
			upper--
			lower--
			// FIXME 反思：比赛时太进张，没有仔细想想其余 return [] 的情况
			if upper < 0 || lower < 0 {
				return [][]int{}
			}
		}
	}
	// fill 1
	for i, sum := range colsum {
		if sum == 1 {
			if upper > 0 {
				ans[0][i] = 1
				upper--
			} else if lower > 0 {
				ans[1][i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}
	return
}
