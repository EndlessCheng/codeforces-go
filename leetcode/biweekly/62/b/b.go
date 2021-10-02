package main

// github.com/EndlessCheng/codeforces-go
func numOfPairs(nums []string, target string) (ans int) {
	cnt := map[string]int{}
	for _, s := range nums {
		cnt[s]++
	}
	for i, n := 1, len(target); i < n; i++ {
		p, s := target[:i], target[i:] // 枚举所有前缀+后缀
		if p != s {
			ans += cnt[p] * cnt[s]
		} else {
			ans += cnt[p] * (cnt[p] - 1)
		}
	}
	return
}
