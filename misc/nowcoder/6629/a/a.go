package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func solve(_ int, a []int) (ans int) {
	evens := []int{}
	for _, v := range a {
		if v&1 == 0 {
			evens = append(evens, v)
		}
	}
	sort.Ints(evens)
	vis := map[int]bool{}
	for i := len(evens) - 1; i >= 0; i-- {
		v := evens[i]
		if vis[v] {
			continue
		}
		for ; v&1 == 0; v >>= 1 {
			vis[v] = true
			ans++
		}
	}
	return
}
