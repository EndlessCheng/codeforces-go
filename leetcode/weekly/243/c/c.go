package main

import "container/heap"

// github.com/EndlessCheng/codeforces-go
func assignTasks(servers, tasks []int) []int {
	// 初始化空闲服务器堆，保存 (服务器权重, 服务器下标)
	idle := make(hp, len(servers))
	for i, w := range servers {
		idle[i] = pair{w, i}
	}
	heap.Init(&idle)

	busy := hp{} // 繁忙服务器堆，保存 (服务器完成任务的时刻, 服务器下标)
	now := 0     // 当前时刻

	// 把已经完成任务的服务器从 busy 移至 idle
	release := func() {
		for len(busy) > 0 && busy[0].t <= now {
			i := heap.Pop(&busy).(pair).i
			heap.Push(&idle, pair{servers[i], i})
		}
	}

	// 依次处理每个任务
	for i, t := range tasks {
		now = max(now, i)
		release()
		if len(idle) == 0 {
			now = busy[0].t // 跳到 busy 的堆顶时刻，从而保证我们有空闲服务器
			release()
		}
		j := heap.Pop(&idle).(pair).i
		heap.Push(&busy, pair{now + t, j})
		tasks[i] = j // 直接把答案记在 tasks 中
	}
	return tasks
}

type pair struct{ t, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.t < b.t || a.t == b.t && a.i < b.i }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
