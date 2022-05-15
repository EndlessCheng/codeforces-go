package main

// github.com/EndlessCheng/codeforces-go
func largestCombination(candidates []int) (ans int) {
	for i := 0; i < 24; i++ {
		s := 0
		for _, v := range candidates {
			s += v >> i & 1
		}
		ans = max(ans, s)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
