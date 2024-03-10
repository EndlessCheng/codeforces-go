package main

import "strings"

// https://space.bilibili.com/206214
func shortestSubstrings(arr []string) []string {
	ans := make([]string, len(arr))
	for i, s := range arr {
		m := len(s)
		res := ""
		for size := 1; size <= m && res == ""; size++ {
		next:
			for k := size; k <= m; k++ {
				sub := s[k-size : k]
				if res != "" && sub >= res {
					continue
				}
				for j, t := range arr {
					if j != i && strings.Contains(t, sub) {
						continue next
					}
				}
				res = sub
			}
		}
		ans[i] = res
	}
	return ans
}
