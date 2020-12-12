package main

// github.com/EndlessCheng/codeforces-go
func findTheDistanceValue(a, b []int, d int) (ans int) {
o:
	for _, v := range a {
		for _, w := range b {
			if abs(v-w) <= d {
				continue o
			}
		}
		ans++
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
