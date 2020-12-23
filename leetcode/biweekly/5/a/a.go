package main

// github.com/EndlessCheng/codeforces-go
func largestUniqueNumber(a []int) (ans int) {
	ans = -1
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	for v, c := range cnt {
		if c == 1 && v > ans {
			ans = v
		}
	}
	return
}
