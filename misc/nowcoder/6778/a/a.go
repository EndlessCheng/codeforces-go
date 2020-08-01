package main

// github.com/EndlessCheng/codeforces-go
func solve(n int, a []int) int {
	if n == 1 {
		return 0
	}
	ans := int(1e9)
	for a0 := a[0] - 1; a0 <= a[0]+1; a0++ {
	o:
		for a1 := a[1] - 1; a1 <= a[1]+1; a1++ {
			c, D := 0, a1-a0
			for i, v := range a {
				d := abs(a0 + i*D - v)
				if d > 1 {
					continue o
				}
				c += d
			}
			if c < ans {
				ans = c
			}
		}
	}
	if ans == 1e9 {
		ans = -1
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
