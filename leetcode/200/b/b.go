package main

// github.com/EndlessCheng/codeforces-go
func getWinner(a []int, k int) (ans int) {
	c := 0
	for i, v := range a {
		if v > ans {
			ans, c = v, 0
			if i > 0 {
				c++
			}
		} else {
			c++
		}
		if c == k {
			break
		}
	}
	return
}
