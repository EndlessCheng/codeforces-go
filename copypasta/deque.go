package copypasta

// 双端队列
// 用两个 slice 头对头拼在一起实现
// 在知道数据量的情况下，也可以直接创建个两倍数据量大小的 slice，然后用两个下标表示头尾，初始化在 slice 正中
// 应用见 graph.go 中的 01 最短路

// l-1,...1,0,0,1...,r-1
type Deque struct{ l, r []interface{} }

func (q Deque) Empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

func (q Deque) Size() int {
	return len(q.l) + len(q.r)
}

func (q *Deque) PushFront(v interface{}) {
	q.l = append(q.l, v)
}

func (q *Deque) PushBack(v interface{}) {
	q.r = append(q.r, v)
}

func (q *Deque) PopFront() (v interface{}) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func (q *Deque) PopBack() (v interface{}) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

func (q Deque) Front() interface{} {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	}
	return q.r[0]
}

func (q Deque) Back() interface{} {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	}
	return q.l[0]
}

// 0 <= i < q.Size()
func (q Deque) Get(i int) interface{} {
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}
