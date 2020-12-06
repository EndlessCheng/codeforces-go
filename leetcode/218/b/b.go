package main

// github.com/EndlessCheng/codeforces-go
func maxOperations(a []int, k int) (ans int) {
	cnt := map[int]int{}
	for _, v := range a {
		cnt[v]++
	}
	if k%2 == 0 {
		ans = cnt[k/2] / 2
		delete(cnt, k/2)
	}
	for v, c := range cnt {
		if v*2 < k && cnt[k-v] > 0 {
			ans += min(c, cnt[k-v])
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
