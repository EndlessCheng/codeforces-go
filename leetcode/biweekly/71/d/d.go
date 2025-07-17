package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func minimumDifference(nums []int) int64 {
	m := len(nums)
	n := m / 3
	minH := minHeap{nums[m-n:]}
	heap.Init(&minH)
	sum := 0
	for _, v := range nums[m-n:] {
		sum += v
	}

	sufMax := make([]int, m-n+1) // 后缀最大和
	sufMax[m-n] = sum
	for i := m - n - 1; i >= n; i-- {
		if v := nums[i]; v > minH.IntSlice[0] {
			sum += v - minH.IntSlice[0]
			minH.IntSlice[0] = v
			heap.Fix(&minH, 0)
		}
		sufMax[i] = sum
	}

	maxH := maxHeap{nums[:n]}
	heap.Init(&maxH)
	preMin := 0 // 前缀最小和
	for _, v := range nums[:n] {
		preMin += v
	}

	ans := preMin - sufMax[n]
	for i := n; i < m-n; i++ {
		if v := nums[i]; v < maxH.IntSlice[0] {
			preMin += v - maxH.IntSlice[0]
			maxH.IntSlice[0] = v
			heap.Fix(&maxH, 0)
		}
		ans = min(ans, preMin-sufMax[i+1])
	}
	return int64(ans)
}

type minHeap struct{ sort.IntSlice }
func (minHeap) Push(any) {}
func (minHeap) Pop() (_ any) { return }

type maxHeap struct{ sort.IntSlice }
func (h maxHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (maxHeap) Push(any) {}
func (maxHeap) Pop() (_ any) { return }
