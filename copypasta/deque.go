package copypasta

// 双端队列
// 另一种实现是 make 个两倍大小的 slice，然后用两个下标 s t 模拟

// l-1,...1,0,0,1...,r-1
// int 可以替换成自己想要的类型
type deque struct{ l, r []int }

func (q deque) empty() bool  { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque) pushL(v int) { q.l = append(q.l, v) }
func (q *deque) pushR(v int) { q.r = append(q.r, v) }
func (q *deque) popL() (v int) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}
func (q *deque) popR() (v int) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q deque) len() int { return len(q.l) + len(q.r) }
func (q deque) topL() int {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}
func (q deque) topR() int {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}
