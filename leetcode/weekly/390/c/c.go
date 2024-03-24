package main

import (
	"container/heap"
	"github.com/emirpasic/gods/trees/redblacktree"
)

// https://space.bilibili.com/206214
func mostFrequentIDs(nums []int, freq []int) []int64 {
	ans := make([]int64, len(nums))
	cnt := make(map[int]int)
	h := hp{}
	heap.Init(&h)
	for i, x := range nums {
		cnt[x] += freq[i]
		heap.Push(&h, pair{cnt[x], x})
		for h[0].c != cnt[h[0].x] { // 堆顶保存的数据已经发生变化
			heap.Pop(&h) // 删除
		}
		ans[i] = int64(h[0].c)
	}
	return ans
}

type pair struct{ c, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].c > h[j].c } // 最大堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func mostFrequentIDs2(nums, freq []int) []int64 {
	cnt := map[int]int{}
	t := redblacktree.New[int, int]()
	ans := make([]int64, len(nums))
	for i, x := range nums {
		// 减少一次 cnt[x] 的出现次数
		node := t.GetNode(cnt[x])
		if node != nil {
			node.Value--
			if node.Value == 0 {
				t.Remove(node.Key)
			}
		}

		cnt[x] += freq[i]

		// 增加一次 cnt[x] 的出现次数
		node = t.GetNode(cnt[x])
		if node == nil {
			t.Put(cnt[x], 1)
		} else {
			node.Value++
		}
		ans[i] = int64(t.Right().Key)
	}
	return ans
}
