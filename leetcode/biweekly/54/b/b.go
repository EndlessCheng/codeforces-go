package main

// github.com/EndlessCheng/codeforces-go
func chalkReplacer(a []int, k int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	k %= s
	for i, v := range a {
		if k < v {
			return i
		}
		k -= v
	}
	panic(1)
}
