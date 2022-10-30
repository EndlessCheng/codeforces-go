package main

// https://space.bilibili.com/206214
func destroyTargets(nums []int, space int) (ans int) {
	g := map[int][]int{}
	for _, x := range nums {
		g[x%space] = append(g[x%space], x)
	}
	mx := 0
	for _, a := range g {
		m := len(a)
		low := a[0]
		for _, x := range a {
			if x < low {
				low = x
			}
		}
		if m > mx || m == mx && low < ans {
			ans = low
			mx = m
		}
	}
	return
}
