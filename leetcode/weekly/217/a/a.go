package main

// github.com/EndlessCheng/codeforces-go
func maximumWealth(accounts [][]int) (ans int) {
	for _, a := range accounts {
		s := 0
		for _, v := range a {
			s += v
		}
		if s > ans {
			ans = s
		}
	}
	return
}
