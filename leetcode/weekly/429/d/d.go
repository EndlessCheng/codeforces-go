package main

import (
	"container/heap"
	"sort"
)

// https://space.bilibili.com/206214
func minLength(s string, numOps int) int {
	n := len(s)
	cnt := 0
	for i, b := range s {
		cnt += (int(b) ^ i) & 1
	}
	if min(cnt, n-cnt) <= numOps {
		return 1
	}

	type pair struct{ k, seg int }
	h := make([][]pair, n+1)
	k := 0
	for i := 0; i < n; i++ {
		k++
		// 到达连续相同子串的末尾
		if i == n-1 || s[i] != s[i+1] {
			h[k] = append(h[k], pair{k, 1})
			k = 0
		}
	}

	i := n
	for range numOps {
		for len(h[i]) == 0 {
			i--
		}
		if i == 2 {
			return 2
		}
		p := h[i][len(h[i])-1]
		h[i] = h[i][:len(h[i])-1]
		p.seg++
		maxSeg := p.k / p.seg
		h[maxSeg] = append(h[maxSeg], p)
	}

	for len(h[i]) == 0 {
		i--
	}
	return i
}

func minLengthHeap(s string, numOps int) int {
	n := len(s)
	cnt := 0
	for i, b := range s {
		cnt += (int(b) ^ i) & 1
	}
	if min(cnt, n-cnt) <= numOps {
		return 1
	}

	h := hp{}
	k := 0
	for i := 0; i < n; i++ {
		k++
		// 到达连续相同子串的末尾
		if i == n-1 || s[i] != s[i+1] {
			h = append(h, tuple{k, k, 1})
			k = 0
		}
	}
	heap.Init(&h)

	for ; numOps > 0 && h[0].maxSeg > 2; numOps-- {
		h[0].seg++
		h[0].maxSeg = h[0].k / h[0].seg // 重新分割
		heap.Fix(&h, 0)
	}
	return h[0].maxSeg
}

type tuple struct{ maxSeg, k, seg int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].maxSeg > h[j].maxSeg } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func minLengthBS(s string, numOps int) int {
	n := len(s)
	return 1 + sort.Search(n-1, func(m int) bool {
		m++
		cnt := 0
		if m == 1 {
			// 改成 0101...
			for i, b := range s {
				// 如果 s[i] 和 i 的奇偶性不同，cnt 加一
				cnt += (int(b) ^ i) & 1
			}
			// n-cnt 表示改成 1010...
			cnt = min(cnt, n-cnt)
		} else {
			k := 0
			for i := range n {
				k++
				// 到达连续相同子串的末尾
				if i == n-1 || s[i] != s[i+1] {
					cnt += k / (m + 1)
					k = 0
				}
			}
		}
		return cnt <= numOps
	})
}
