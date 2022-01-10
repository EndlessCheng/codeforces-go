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
	day := 0
	for _, p := range a {
		day += p.p
		if day+p.g > ans {
			ans = day + p.g
		}
	}
	return
}
