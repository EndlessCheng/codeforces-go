package main

import (
	"cmp"
	"container/heap"
)

// https://space.bilibili.com/206214
type pair struct{ priority, userId int }

type TaskManager struct {
	mp map[int]pair // taskId -> (priority, userId)
	h  *hp          // (priority, taskId, userId)
}

func Constructor(tasks [][]int) TaskManager {
	n := len(tasks)
	mp := make(map[int]pair, n) // 预分配空间
	h := make(hp, n)
	for i, t := range tasks {
		userId, taskId, priority := t[0], t[1], t[2]
		mp[taskId] = pair{priority, userId}
		h[i] = tuple{priority, taskId, userId}
	}
	heap.Init(&h)
	return TaskManager{mp, &h}
}

func (t *TaskManager) Add(userId, taskId, priority int) {
	t.mp[taskId] = pair{priority, userId}
	heap.Push(t.h, tuple{priority, taskId, userId})
}

func (t *TaskManager) Edit(taskId, newPriority int) {
	// 懒修改
	t.Add(t.mp[taskId].userId, taskId, newPriority)
}

func (t *TaskManager) Rmv(taskId int) {
	// 懒删除
	t.mp[taskId] = pair{-1, -1}
}

func (t *TaskManager) ExecTop() int {
	for t.h.Len() > 0 {
		top := heap.Pop(t.h).(tuple)
		priority, taskId, userId := top.priority, top.taskId, top.userId
		if t.mp[taskId] == (pair{priority, userId}) {
			t.Rmv(taskId)
			return userId
		}
		// else 货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
	}
	return -1
}

type tuple struct{ priority, taskId, userId int }
type hp []tuple

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return cmp.Or(h[i].priority-h[j].priority, h[i].taskId-h[j].taskId) > 0
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
