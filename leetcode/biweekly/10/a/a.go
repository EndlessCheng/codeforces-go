package main

// github.com/EndlessCheng/codeforces-go
func arraysIntersection(a, b, c []int) (ans []int) {
	c1 := map[int]bool{}
	for _, v := range a {
		c1[v] = true
	}
	c2 := map[int]bool{}
	for _, v := range b {
		c2[v] = true
	}
	for _, v := range c {
		if c1[v] && c2[v] {
			ans = append(ans, v)
		}
	}
	return
}
