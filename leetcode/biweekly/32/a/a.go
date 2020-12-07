package main

// github.com/EndlessCheng/codeforces-go
func findKthPositive(a []int, k int) (ans int) {
	has := map[int]bool{}
	for _, v := range a {
		has[v] = true
	}
	for i := 1; ; i++ {
		if !has[i] {
			if k == 1 {
				return i
			}
			k--
		}
	}
}
