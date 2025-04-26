package copypasta

import "sort"

// 双端队列
// 用两个 slice 头对头拼在一起实现
// 在知道数据量的情况下，也可以直接创建个两倍数据量大小的 slice，然后用两个下标表示头尾，初始化在 slice 正中
// 应用见 graph.go 中的 01 最短路
// https://codeforces.com/problemset/problem/1584/E 2300 单调双端队列

// l-1,...1,0,0,1...,r-1
type deque struct{ l, r []int }

func (q deque) empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q deque) size() int {
	return len(q.l) + len(q.r)
}

func (q *deque) pushFront(v int) {
	q.l = append(q.l, v)
}

func (q *deque) pushBack(v int) {
	q.r = append(q.r, v)
}

func (q *deque) popFront() (v int) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *deque) popBack() (v any) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque) front() int {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q deque) back() int {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

// 0 <= i < q.size()
func (q deque) get(i int) int {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}

// 假设 q 是有序的，二分找 >= v 的第一个数的下标
func (q deque) search(v int) int {
	if len(q.l) > 0 && v <= q.l[0] {
		// q.l 是递减的 
		i := sort.Search(len(q.l), func(i int) bool { return q.l[i] < v })
		return len(q.l) - i
	}
	return len(q.l) + sort.SearchInts(q.r, v)
}
