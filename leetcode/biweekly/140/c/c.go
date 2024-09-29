package main

// https://space.bilibili.com/206214
func validSequence(s, t string) []int {
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	suf[n] = m
	for i, j := n-1, m-1; i >= 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = j + 1
	}

	ans := make([]int, m)
	j := 0
	changed := false // 是否修改过
	for i := range s {
		if s[i] == t[j] || !changed && suf[i+1] <= j+1 {
			if s[i] != t[j] {
				changed = true
			}
			ans[j] = i
			j++
			if j == m {
				return ans
			}
		}
	}
	return nil
}
