package main

// github.com/EndlessCheng/codeforces-go
func maximumRequests(n int, req [][]int) (ans int) {
	f := func(sub int) (res int) {
		cnt := make([]int, n)
		for i, r := range req {
			if sub>>i&1 > 0 {
				cnt[r[0]]--
				cnt[r[1]]++
				res++
			}
		}
		for _, c := range cnt {
			if c != 0 {
				return 0
			}
		}
		return
	}
	for sub := 0; sub < 1<<len(req); sub++ {
		if res := f(sub); res > ans {
			ans = res
		}
	}
	return
}
