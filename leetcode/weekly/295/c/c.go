package main

// https://space.bilibili.com/206214/dynamic
func totalSteps(nums []int) (ans int) {
	type pair struct{ v, t int }
	st := []pair{}
	for _, num := range nums {
		maxT := 0
		for len(st) > 0 && st[len(st)-1].v <= num {
			maxT = max(maxT, st[len(st)-1].t)
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			maxT++
			ans = max(ans, maxT)
		} else {
			maxT = 0
		}
		st = append(st, pair{num, maxT})
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
