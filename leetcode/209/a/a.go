package main

// github.com/EndlessCheng/codeforces-go
func specialArray(a []int) (ans int) {
	for x := 0; ; x++ {
		c := 0
		for _, v := range a {
			if v >= x {
				c++
			}
		}
		if c == x {
			return x
		}
		if c < x {
			return -1
		}
	}
}
