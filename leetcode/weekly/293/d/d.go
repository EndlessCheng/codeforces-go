package main

// github.com/EndlessCheng/codeforces-go
type CountIntervals struct {
	left, right *CountIntervals
	l, r, sum   int
}

func Constructor() CountIntervals { return CountIntervals{l: 1, r: 1e9} }

func (o *CountIntervals) Add(l, r int) {
	if o.sum == o.r-o.l+1 { return } // o 已被完整覆盖，无需执行任何操作
	if l <= o.l && o.r <= r { // 当前节点已被区间 [l,r] 完整覆盖，不再继续递归
		o.sum = o.r - o.l + 1
		return
	}
	mid := (o.l + o.r) >> 1
	if o.left == nil { o.left = &CountIntervals{l: o.l, r: mid} } // 动态开点
	if o.right == nil { o.right = &CountIntervals{l: mid + 1, r: o.r} } // 动态开点
	if l <= mid { o.left.Add(l, r)}
	if mid < r { o.right.Add(l, r) }
	o.sum = o.left.sum + o.right.sum
}

func (o *CountIntervals) Count() int { return o.sum }
