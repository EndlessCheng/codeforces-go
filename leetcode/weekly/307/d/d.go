package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func kSum(nums []int, k int) int64 {
	sum, total := 0, 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
			total += x
		} else {
			total -= x
			nums[i] = -x
		}
	}
	slices.Sort(nums)

	kthS := sort.Search(total, func(sumLimit int) bool {
		cnt := 1 // 空子序列算一个
		var dfs func(int, int)
		dfs = func(i, s int) {
			if cnt == k || i == len(nums) || s+nums[i] > sumLimit {
				return
			}
			cnt++               // s + nums[i] <= sumLimit
			dfs(i+1, s+nums[i]) // 选
			dfs(i+1, s)         // 不选
		}
		dfs(0, 0)
		return cnt == k // 找到 k 个元素和不超过 sumLimit 的子序列
	})
	return int64(sum - kthS)
}

func kSum2(nums []int, k int) int64 {
	n := len(nums)
	sum := 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
		} else {
			nums[i] = -x
		}
	}
	slices.Sort(nums)

	h := hp{{0, 0}} // 空子序列
	for ; k > 1; k-- {
		p := heap.Pop(&h).(pair)
		i := p.i
		if i < n {
			// 在子序列的末尾添加 nums[i]
			heap.Push(&h, pair{p.sum + nums[i], i + 1}) // 下一个添加/替换的元素下标为 i+1
			if i > 0 { // 替换子序列的末尾元素为 nums[i]
				heap.Push(&h, pair{p.sum + nums[i] - nums[i-1], i + 1})
			}
		}
	}
	return int64(sum - h[0].sum)
}

type pair struct{ sum, i int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
