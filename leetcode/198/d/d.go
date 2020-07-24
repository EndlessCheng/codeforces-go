package main

// github.com/EndlessCheng/codeforces-go
func closestToTarget(a []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	ans := int(1e9)
	s := []int{}
	for _, v := range a {
		for i := range s {
			s[i] &= v
		}
		s = append(s, v)
		j := 0
		for i := 1; i < len(s); i++ {
			if s[j] != s[i] {
				j++
				s[j] = s[i]
			}
		}
		s = s[:j+1]
		for _, v := range s {
			ans = min(ans, abs(v-target))
		}
	}
	return ans
}
