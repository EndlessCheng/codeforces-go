package main

// github.com/EndlessCheng/codeforces-go
func numberOfWeeks(milestones []int) int64 {
	s, m := 0, 0
	for _, x := range milestones {
		s += x
		m = max(m, x)
	}
	return int64(min(s, (s-m)*2+1))
}
