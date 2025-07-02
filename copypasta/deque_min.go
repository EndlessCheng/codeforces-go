package copypasta

import "math"

// 前置内容：最小栈
// LC155 https://leetcode.cn/problems/min-stack/
type minStPair struct{ val, preMin int }

type minStack []minStPair

func newMinStack() minStack {
	// 这里的 0 用不到
	return minStack{{0, math.MaxInt}} // 栈底哨兵
}

func (st minStack) min() int {
	return st[len(st)-1].preMin
}

func (st minStack) empty() bool {
	return len(st) == 1
}

func (st *minStack) push(v int) {
	*st = append(*st, minStPair{v, min(st.min(), v)})
}

func (st *minStack) pop() int {
	v := (*st)[len(*st)-1].val
	*st = (*st)[:len(*st)-1]
	return v
}

func (st minStack) top() int {
	return st[len(st)-1].val
}

// 最小双端队列
// 用两个最小栈底对底
type minDeque struct{ l, r minStack }

func newMinDeque() minDeque {
	return minDeque{newMinStack(), newMinStack()}
}

// 时间复杂度：均摊 O(1)    
// 定义势能函数为 abs(l.size() - r.size())
// 一开始势能是 0
// 左右入队最多把势能增大 1
// 而 rebalance 会把势能减少至 0 或者 1
func (q *minDeque) rebalance() {
	if q.r.empty() {
		q.l, q.r = q.r, q.l
		defer func() { q.l, q.r = q.r, q.l }()
	}

	m := len(q.r) / 2
	for i := m; i > 0; i-- { // 注意 q.r[0] 是哨兵
		q.l.push(q.r[i].val)
	}
	// 重新计算后半段的前缀最小值
	t := q.r[m+1:]
	q.r = q.r[:1]
	for _, p := range t {
		q.r.push(p.val)
	}
}

// 如果 q 是空的，返回 math.MaxInt
func (q minDeque) min() int {
	return min(q.l.min(), q.r.min())
}

func (q minDeque) empty() bool {
	return q.l.empty() && q.r.empty()
}

func (q minDeque) size() int {
	return len(q.l) + len(q.r) - 2 // 减去栈底哨兵
}

func (q *minDeque) pushFront(v int) {
	q.l.push(v)
}

func (q *minDeque) pushBack(v int) {
	q.r.push(v)
}

func (q *minDeque) popFront() int {
	if q.l.empty() {
		q.rebalance()
	}
	return q.l.pop()
}

func (q *minDeque) popBack() int {
	if q.r.empty() {
		q.rebalance()
	}
	return q.r.pop()
}

func (q *minDeque) front() int {
	if q.l.empty() {
		q.rebalance()
	}
	return q.l.top()
}

func (q *minDeque) back() int {
	if q.r.empty() {
		q.rebalance()
	}
	return q.r.top()
}
