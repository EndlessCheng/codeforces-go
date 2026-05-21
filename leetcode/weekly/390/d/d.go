package main

import (
	"math"
	"runtime/debug"
)

// https://space.bilibili.com/206214
func init() { debug.SetGCPercent(-1) }

func stringIndices(wordsContainer, wordsQuery []string) []int {
	type node struct {
		son       [26]*node
		minLen    int // 子树中的最短字符串的长度
		bestIndex int // 子树中的最短字符串的下标
	}
	root := &node{minLen: math.MaxInt}

	for i, s := range wordsContainer {
		l := len(s)
		if l < root.minLen {
			root.minLen = l
			root.bestIndex = i
		}

		// 把 reverse(s) 插入字典树
		cur := root
		for j := l - 1; j >= 0; j-- {
			b := s[j] - 'a'
			if cur.son[b] == nil {
				cur.son[b] = &node{minLen: math.MaxInt}
			}
			cur = cur.son[b]
			// 维护 cur 子树中的最短字符串的长度及其下标
			// 由于我们是按照 i 从小到大的顺序遍历，字符串长度相同时不更新 bestIndex
			if l < cur.minLen {
				cur.minLen = l
				cur.bestIndex = i
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	for i, s := range wordsQuery {
		cur := root
		for j := len(s) - 1; j >= 0 && cur.son[s[j]-'a'] != nil; j-- {
			cur = cur.son[s[j]-'a']
		}
		ans[i] = cur.bestIndex
	}
	return ans
}
