package main

import (
	"container/heap"
	"sort"
)

// 事件扫描线+堆

// github.com/EndlessCheng/codeforces-go
func smallestChair(times [][]int, targetFriend int) int {
	// 按时间顺序，记录每个到达事件和离开事件相对应的朋友编号
	events := make([][2][]int, 1e5+1)
	for i, t := range times {
		l, r := t[0], t[1]
		events[l][1] = append(events[l][1], i) // 朋友到达
		events[r][0] = append(events[r][0], i) // 朋友离开
	}

	// 初始化未被占据的椅子
	n := len(times)
	unoccupied := hp{make([]int, n)}
	for i := range unoccupied.IntSlice {
		unoccupied.IntSlice[i] = i
	}

	// 按时间顺序扫描每个事件
	belong := make([]int, n)
	for _, e := range events {
		for _, id := range e[0] { // 朋友离开
			heap.Push(&unoccupied, belong[id]) // 返还椅子
		}
		for _, id := range e[1] { // 朋友到达
			belong[id] = heap.Pop(&unoccupied).(int) // 记录占据该椅子的朋友编号
			if id == targetFriend {
				return belong[id]
			}
		}
	}
	return 0
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
