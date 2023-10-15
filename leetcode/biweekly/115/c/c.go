package main

// https://space.bilibili.com/206214
func ok(s, t string) (diff bool) {
	if len(s) != len(t) {
		return
	}
	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return
}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	f := make([]int, n)
	from := make([]int, n)
	mx := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j]) {
				f[i] = f[j]
				from[i] = j
			}
		}
		f[i]++ // 加一写在这里
		if f[i] > f[mx] {
			mx = i
		}
	}

	ans := make([]string, f[mx])
	for i := range ans {
		ans[i] = words[mx]
		mx = from[mx]
	}
	return ans
}
