package main

import "container/heap"

// https://space.bilibili.com/206214
type EventManager struct {
	idToPriority map[int]int
	h            *hp
}

func Constructor(events [][]int) EventManager {
	n := len(events)
	idToPriority := make(map[int]int, n) // 预分配空间
	h := make(hp, n)
	for i, e := range events {
		id, priority := e[0], e[1]
		idToPriority[id] = priority
		h[i] = event{priority, id}
	}
	heap.Init(&h)
	return EventManager{idToPriority, &h}
}

func (m EventManager) UpdatePriority(eventId, newPriority int) {
	m.idToPriority[eventId] = newPriority
	heap.Push(m.h, event{newPriority, eventId})
}

func (m EventManager) PollHighest() int {
	for m.h.Len() > 0 {
		e := heap.Pop(m.h).(event)
		if m.idToPriority[e.id] == e.priority {
			delete(m.idToPriority, e.id)
			return e.id
		}
		// else 货不对板，继续找下一个
	}
	return -1
}

type event struct{ priority, id int }
type hp []event

func (h hp) Len() int      { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].priority > h[j].priority || h[i].priority == h[j].priority && h[i].id < h[j].id
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(event)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
