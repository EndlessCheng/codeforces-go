package main

import (
	"cmp"
	"container/heap"
)

// https://space.bilibili.com/206214
type pair struct{ priority, userId int }

type TaskManager struct {
	h  *hp          // (priority, taskId, userId)
	mp map[int]pair // taskId -> (priority, userId)
}

func Constructor(tasks [][]int) TaskManager {
	h := hp{}
	mp := map[int]pair{}
	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		mp[taskId] = pair{priority, userId}
		h = append(h, tuple{priority, taskId, userId})
	}
	heap.Init(&h)
	return TaskManager{&h, mp}
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	tm.mp[taskId] = pair{priority, userId}
	heap.Push(tm.h, tuple{priority, taskId, userId})
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	// 懒修改
	tm.Add(tm.mp[taskId].userId, taskId, newPriority)
}

func (tm *TaskManager) Rmv(taskId int) {
	// 懒删除
	tm.mp[taskId] = pair{-1, -1}
}

func (tm *TaskManager) ExecTop() int {
	for tm.h.Len() > 0 {
		top := heap.Pop(tm.h).(tuple)
		priority, taskId, userId := top.priority, top.taskId, top.userId
		// 如果货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
		if tm.mp[taskId] == (pair{priority, userId}) {
			tm.Rmv(taskId)
			return userId
		}
	}
	return -1
}

type tuple struct{ priority, taskId, userId int }
type hp []tuple

func (h hp) Len() int      { return len(h) }
func (h hp) Less(i, j int) bool {
	return cmp.Or(h[i].priority-h[j].priority, h[i].taskId-h[j].taskId) > 0
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
