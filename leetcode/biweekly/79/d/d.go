package main

// https://space.bilibili.com/206214/dynamic
type seg []struct{ l, r, min, sum int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// 将 idx 上的元素值增加 val
func (t seg) add(o, idx, val int) {
	if t[o].l == t[o].r {
		t[o].min += val
		t[o].sum += val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if idx <= m {
		t.add(o<<1, idx, val)
	} else {
		t.add(o<<1|1, idx, val)
	}
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].min = min(lo.min, ro.min)
	t[o].sum = lo.sum + ro.sum
}

// 返回区间 [l,r] 内的元素和
func (t seg) querySum(o, l, r int) (sum int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		sum += t.querySum(o<<1, l, r)
	}
	if r > m {
		sum += t.querySum(o<<1|1, l, r)
	}
	return
}

// 查询 [1,r] 上 <= val 的最靠左的位置
// 不存在时返回 0
func (t seg) queryFirstIdx(o, r, val int) int {
	if t[o].min > val { // 说明整个区间的元素值都大于 val
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	m := (t[o].l + t[o].r) >> 1
	if t[o<<1].min <= val { // 看看左半部分
		return t.queryFirstIdx(o<<1, r, val)
	}
	if m < r { // 看看右半部分
		return t.queryFirstIdx(o<<1|1, r, val)
	}
	return 0
}

type BookMyShow struct {
	seg
	m, i int
}

func Constructor(n, m int) BookMyShow {
	t := make(seg, n*4)
	t.build(1, 1, n)
	return BookMyShow{t, m, 1}
}

func (t BookMyShow) Gather(k, maxRow int) []int {
	i := t.queryFirstIdx(1, maxRow+1, t.m-k)
	if i == 0 { // 不存在
		return nil
	}
	seats := t.querySum(1, i, i)
	t.add(1, i, k) // 占据 k 个座位
	return []int{i - 1, seats}
}

func (t *BookMyShow) Scatter(k, maxRow int) bool {
	if (maxRow+1)*t.m-t.querySum(1, 1, maxRow+1) < k { // 剩余座位不足 k 个
		return false
	}
	// 从第一个没有坐满的排开始占座
	for ; ; t.i++ {
		leftSeats := t.m - t.querySum(1, t.i, t.i)
		if k <= leftSeats { // 剩余人数不够坐后面的排
			t.add(1, t.i, k)
			return true
		}
		k -= leftSeats
		t.add(1, t.i, leftSeats)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
