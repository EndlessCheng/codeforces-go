package main

// github.com/EndlessCheng/codeforces-go
func countPairs(a []int) (ans int) {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	for v, c := range cnt {
		for i := 0; i < 22; i++ {
			w := 1<<i - v
			if v < w {
				ans += c * cnt[w]
			} else if v == w {
				ans += c * (c - 1) / 2
			}
		}
	}
	return ans % (1e9 + 7)
}
