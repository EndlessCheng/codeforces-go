package main

// github.com/EndlessCheng/codeforces-go
func minMoves(v, maxDoubles int) (ans int) {
	for v > 1 {
		if maxDoubles == 0 {
			return ans + v - 1
		}
		if v%2 > 0 {
			v--
			ans++
		}
		maxDoubles--
		v /= 2
		ans++
	}
	return
}
