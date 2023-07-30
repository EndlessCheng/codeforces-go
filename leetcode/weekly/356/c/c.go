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
	for i := min(len(s), len(t)); ; i-- {
		// 枚举：s 的后 i 个字母和 t 的前 i 个字母是一样的
		if s[len(s)-i:] == t[:i] {
			return s + t[i:]
		}
	}
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
