package main

// https://space.bilibili.com/206214
func numberOfAlternatingGroups(colors []int, k int) (ans int) {
	n := len(colors)
	cnt := 0
	for i := range n + k - 1 {
		if colors[i%n] == colors[(i+1)%n] {
			cnt = 0
		}
		cnt++
		if cnt >= k {
			ans++
		}
	}
	return
}

func numberOfAlternatingGroups2(colors []int, k int) (ans int) {
	n := len(colors)
	cnt := 0
	for i := range n * 2 {
		if colors[i%n] == colors[(i+1)%n] {
			cnt = 0
		}
		cnt++
		if i >= n && cnt >= k {
			ans++
		}
	}
	return
}
