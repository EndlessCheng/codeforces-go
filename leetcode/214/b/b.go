package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minDeletions(s string) (ans int) {
	c := make(sort.IntSlice, 26)
	for _, b := range s {
		c[b-'a']++
	}
	sort.Sort(sort.Reverse(c))
	vis := map[int]bool{}
o:
	for _, v := range c {
		if !vis[v] {
			vis[v] = true
			continue
		}
		for j := v - 1; j > 0; j-- {
			if !vis[j] {
				ans += v - j
				vis[j] = true
				continue o
			}
		}
		ans += v
	}
	return
}
