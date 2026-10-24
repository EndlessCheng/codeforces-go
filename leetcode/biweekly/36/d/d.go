package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func busiestServers(k int, arrival, load []int) (ans []int) {
	// 小优化：如果 k >= len(arrival)，我们只会用到前 len(arrival) 台服务器，且每台服务器恰好使用一次
	if k >= len(arrival) {
		ans = make([]int, len(arrival))
		for i := range ans {
			ans[i] = i
		}
		return
	}

	// 初始化空闲服务器堆，保存空闲服务器的虚拟编号
	idle := hp{make([]int, k)}
	for i := range k {
		idle.IntSlice[i] = i
	}

	busy := hp2{}         // 繁忙服务器堆，保存 (服务器完成请求的时刻, 服务器编号)
	cnt := make([]int, k) // cnt[i] 统计服务器 i 处理的请求数
	maxCnt := 0

	for i, t := range arrival {
		// 把已经完成请求的服务器从 busy 移至 idle
		for len(busy) > 0 && busy[0].t <= t {
			j := heap.Pop(&busy).(pair).i
			// 把服务器 j 在 [i,i+k) 中的同余值 virtual 作为该服务器的虚拟编号
			virtual := i + (j-i%k+k)%k // +k 保证取模结果非负
			heap.Push(&idle, virtual)
		}

		// 没有空闲服务器，舍弃该请求
		if idle.Len() == 0 {
			continue
		}

		j := heap.Pop(&idle).(int) % k
		heap.Push(&busy, pair{t + load[i], j})

		cnt[j]++
		if cnt[j] > maxCnt { // j 是最繁忙的服务器
			maxCnt = cnt[j]
			ans = []int{j}
		} else if cnt[j] == maxCnt { // j 是最繁忙的服务器之一
			ans = append(ans, j)
		}
	}

	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type pair struct{ t, i int }
type hp2 []pair

func (h hp2) Len() int           { return len(h) }
func (h hp2) Less(i, j int) bool { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
