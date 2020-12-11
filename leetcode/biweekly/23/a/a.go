package main

// github.com/EndlessCheng/codeforces-go
func countLargestGroup(n int) (ans int) {
	cnt := map[int]int{}
	for i := 1; i <= n; i++ {
		s := 0
		for v := i; v > 0; v /= 10 {
			s += v % 10
		}
		cnt[s]++
	}
	mx := 0
	for _, c := range cnt {
		if c > mx {
			mx, ans = c, 1
		} else if c == mx {
			ans++
		}
	}
	return
}
