package main

import "math"

// https://space.bilibili.com/206214
func stringIndices(wordsContainer, wordsQuery []string) []int {
	type node struct {
		son     [26]*node
		minL, i int
	}
	root := &node{minL: math.MaxInt}

	for idx, s := range wordsContainer {
		l := len(s)
		cur := root
		if l < cur.minL {
			cur.minL, cur.i = l, idx
		}
		for i := len(s) - 1; i >= 0; i-- {
			b := s[i] - 'a'
			if cur.son[b] == nil {
				cur.son[b] = &node{minL: math.MaxInt}
			}
			cur = cur.son[b]
			if l < cur.minL {
				cur.minL, cur.i = l, idx
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	for idx, s := range wordsQuery {
		cur := root
		for i := len(s) - 1; i >= 0 && cur.son[s[i]-'a'] != nil; i-- {
			cur = cur.son[s[i]-'a']
		}
		ans[idx] = cur.i
	}
	return ans
}
