package main

import (
	"index/suffixarray"
	"reflect"
	"sort"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func longestCommonSubpath(_ int, paths [][]int) (ans int) {
	a := []int{}
	minLen := int(1e9) // 二分右边界
	for _, p := range paths {
		minLen = min(minLen, len(p))
		a = append(a, 1e9) // 用一个不存在 paths 中的数拼接所有路径
		a = append(a, p...)
	}
	n, m := len(a), len(paths)

	// 标记每个元素属于哪条路径
	ids := make([]int, n)
	id := -1
	for i, v := range a {
		if v == 1e9 {
			id++
			ids[i] = m
		} else {
			ids[i] = id
		}
	}

	// 构建 a 的后缀数组和高度数组
	s := make([]byte, 0, n*4)
	for _, v := range a {
		s = append(s, byte(v>>24), byte(v>>16&0xff), byte(v>>8&0xff), byte(v&0xff))
	}
	_sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	sa := make([]int32, 0, n)
	for _, v := range _sa {
		if v&3 == 0 {
			sa = append(sa, v>>2)
		}
	}
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, rk := range rank {
		if h > 0 {
			h--
		}
		if rk > 0 {
			for j := int(sa[rk-1]); i+h < n && j+h < n && a[i+h] == a[j+h]; h++ {
			}
		}
		height[rk] = h
	}

	// 二分求答案
	return sort.Search(minLen, func(limit int) bool {
		limit++
		vis := make([]int, m)
		for i := 1; i < n; i++ {
			if height[i] < limit {
				continue
			}
			cnt := 0
			for st := i; i < n && height[i] >= limit; i++ {
				// 检查 sa[i] 和 sa[i-1]
				if j := ids[sa[i]]; j < m && vis[j] != st {
					vis[j] = st
					cnt++
				}
				if j := ids[sa[i-1]]; j < m && vis[j] != st {
					vis[j] = st
					cnt++
				}
			}
			// 连续 m 个属于不同路径的后缀长度均不小于 limit
			if cnt == m {
				return false
			}
		}
		return true
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
