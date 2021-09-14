package main

// github.com/EndlessCheng/codeforces-go
func mostCompetitive(a []int, k int) (s []int) {
	for i, v := range a {
		for len(s) > 0 && len(s)+len(a)-1-i >= k && v < s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, v)
		}
	}
	return
}
