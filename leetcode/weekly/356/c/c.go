package main

import "strings"

// https://space.bilibili.com/206214
func merge(s, t string) string {
	// 先特判完全包含的情况
	if strings.Contains(s, t) {
		return s
	}
	if strings.Contains(t, s) {
		return t
	}

	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	match := calcMaxMatchLengths(t+"#"+s)
	return s + t[match[len(match)-1]:]
}

func minimumString(a, b, c string) (ans string) {
	arr := []string{a, b, c}
	// 枚举 arr 的全排列
	for _, p := range [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}} {
		s := merge(merge(arr[p[0]], arr[p[1]]), arr[p[2]])
		if ans == "" || len(s) < len(ans) || len(s) == len(ans) && s < ans {
			ans = s
		}
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
