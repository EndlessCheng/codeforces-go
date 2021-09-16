package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxEqualFreq(a []int) (ans int) {
	c := map[int]int{}
	cc := make([]int, len(a)+1)
	for i, v := range a {
		c[v]++
		c := c[v]
		cc[c]++
		if cc[c]*c == i {
			ans = i + 1
		} else if cc[c]*c == i+1 && i+1 < len(a) {
			ans = i + 2
		}
	}
	return
}

func maxEqualFreq2(a []int) (ans int) {
	c := map[int]int{}
	cc := map[int]int{}
	for i, v := range a {
		if c := c[v]; cc[c] > 1 {
			cc[c]--
		} else if cc[c] == 1 {
			delete(cc, c)
		}
		c[v]++
		cc[c[v]]++
		if len(cc) > 2 {
			continue
		}
		cs := make([]int, 0, len(cc))
		for c := range cc {
			cs = append(cs, c)
		}
		sort.Ints(cs)
		if len(cc) == 1 {
			if cs[0] == 1 || cc[cs[0]] == 1 {
				ans = i + 1
			}
		} else {
			if cs[0] == 1 && cc[cs[0]] == 1 || cs[1]-cs[0] == 1 && cc[cs[1]] == 1 {
				ans = i + 1
			}
		}
	}
	return
}
