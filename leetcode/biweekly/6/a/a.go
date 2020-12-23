package main

// github.com/EndlessCheng/codeforces-go
func isMajorityElement(a []int, target int) (ans bool) {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	mx, val := 0, 0
	for v, c := range cnt {
		if c > mx {
			mx, val = c, v
		}
	}
	return mx > len(a)/2 && val == target
}
