package main

// github.com/EndlessCheng/codeforces-go
func canDivideIntoSubsequences(a []int, k int) (ans bool) {
	cnt := map[int]int{}
	mx := 0
	for _, v := range a {
		if cnt[v]++; cnt[v] > mx {
			mx = cnt[v]
		}
	}
	return len(a) >= k*mx
}
