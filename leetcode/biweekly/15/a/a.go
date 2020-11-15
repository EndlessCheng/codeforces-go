package main

// github.com/EndlessCheng/codeforces-go
func findSpecialInteger(a []int) (ans int) {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	for v, c := range cnt {
		if c > cnt[ans] {
			ans = v
		}
	}
	return
}
