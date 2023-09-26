package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func earliestFullBloom(plantTime, growTime []int) (ans int) {
	type pair struct{ p, g int }
	a := make([]pair, len(plantTime))
	for i, p := range plantTime {
		a[i] = pair{p, growTime[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].g > a[j].g })
	days := 0
	for _, p := range a {
		days += p.p
		ans = max(ans, days+p.g)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
