package main

// github.com/EndlessCheng/codeforces-go
func removeInterval(intervals [][]int, toBeRemoved []int) (ans [][]int) {
	l, r := toBeRemoved[0], toBeRemoved[1]
	for _, p := range intervals {
		if p[1] <= l || p[0] >= r {
			ans = append(ans, p)
			continue
		}
		if p[0] < l {
			ans = append(ans, []int{p[0], l})
		}
		if p[1] > r {
			ans = append(ans, []int{r, p[1]})
		}
	}
	return
}
