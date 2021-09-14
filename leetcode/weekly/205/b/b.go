package main

// github.com/EndlessCheng/codeforces-go
func numTriplets(a []int, b []int) (ans int) {
	f := func(a, b []int) {
		cnt := map[int]int{}
		for _, v := range a {
			cnt[v*v]++
		}
		for i, v := range b {
			for j := i + 1; j < len(b); j++ {
				ans += cnt[v*b[j]]
			}
		}
	}
	f(a, b)
	f(b, a)
	return
}
