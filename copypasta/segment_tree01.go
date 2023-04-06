package copypasta

// 0-1 线段树
// 支持区间翻转比特、单比特加减等
// 某些情况下可作为 Bitset 的代替品
// LC2569 https://leetcode.cn/problems/handling-sum-queries-after-update/
// https://codeforces.com/contest/1705/problem/E
// https://codeforces.com/problemset/problem/877/E
type seg01 []struct {
	l, r int
	ones int // EXTRA: 1 的个数
	// -1: [l,r] 内全为 0
	//  1: [l,r] 内全为 1
	//  0: [l,r] 内有 0 有 1
	state int8
	flip  bool // lazy tag
}

// 见 buildWithBinary
func newSeg01(a string) seg01 {
	t := make(seg01, 4*len(a))
	t.buildWithBinary(a, 1, 1, len(a))
	return t
}

func (t seg01) maintain(o int) {
	lo, ro := &t[o<<1], &t[o<<1|1]
	if lo.state < 0 && ro.state < 0 {
		t[o].state = -1
	} else if lo.state > 0 && ro.state > 0 {
		t[o].state = 1
	} else {
		t[o].state = 0
	}
	t[o].ones = lo.ones + ro.ones
}

func (t seg01) build(o, l, r int) {
	t[o].l, t[o].r, t[o].state = l, r, -1 // 初始全为 0，故设置为 -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// a 从左到右是二进制从低到高
// a 的下标从 0 开始（虽然整个 0-1 线段树都是从 1 开始的）
func (t seg01) buildWithBinary(a string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		if a[l-1] == '0' {
			t[o].state = -1
		} else { // '1'
			t[o].state = 1
			t[o].ones = 1
		}
		return
	}
	m := (l + r) >> 1
	t.buildWithBinary(a, o<<1, l, m)
	t.buildWithBinary(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg01) doFlip(O int) {
	o := &t[O]
	o.state = -o.state
	o.ones = o.r - o.l + 1 - o.ones
	o.flip = !o.flip
}

func (t seg01) spread(o int) {
	if t[o].flip {
		t.doFlip(o << 1)
		t.doFlip(o<<1 | 1)
		t[o].flip = false
	}
}

// 将 [l,r] 内的 0 置为 1，1 置为 0
// o=1, 1<=l<=r<=n
func (t seg01) flip(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.doFlip(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.flip(o<<1, l, r)
	}
	if m < r {
		t.flip(o<<1|1, l, r)
	}
	t.maintain(o)
}

// 返回下标 >= l 的第一个 0 的位置，不存在时返回 -1
// o=1, l>=1
func (t seg01) next0(o, l int) int {
	if t[o].l == t[o].r {
		if t[o].state < 0 {
			return t[o].l
		}
		return -1
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m && t[o<<1].state <= 0 {
		if p := t.next0(o<<1, l); p > 0 {
			return p
		}
	}
	return t.next0(o<<1|1, l)
}

// 返回下标 >= l 的第一个 1 的位置，不存在时返回 -1
// o=1, l>=1
func (t seg01) next1(o, l int) int {
	if t[o].l == t[o].r {
		if t[o].state > 0 {
			return t[o].l
		}
		return -1
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m && t[o<<1].state >= 0 {
		if p := t.next1(o<<1, l); p > 0 {
			return p
		}
	}
	return t.next1(o<<1|1, l)
}

// 返回第 k 个 1 的位置
// 必须满足 k <= t[1].ones
// o=1, k>=1
func (t seg01) kth1(o, k int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	if k <= t[o<<1].ones {
		return t.kth1(o<<1, k)
	}
	return t.kth1(o<<1|1, k-t[o<<1].ones)
}

// 返回最后一个 1 的位置（类似 bits.Len）
// 如果题目不保证此时一定有 1，则要特判下：如果 t[1].state < 0 则没有 1
// o=1
func (t seg01) lastIndex1(o int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	if t[o<<1|1].state >= 0 {
		return t.lastIndex1(o<<1 | 1)
	}
	return t.lastIndex1(o << 1)
}

// += 1<<i，模拟进位
// i>=1
func (t seg01) add(i int) { t.flip(1, i, t.next0(1, i)) }

// -= 1<<i，模拟借位
// i>=1
func (t seg01) sub(i int) { t.flip(1, i, t.next1(1, i)) }

// 返回 [l,r] 内 1 的个数
// o=1, 1<=l<=r<=n
func (t seg01) onesCount(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].ones
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.onesCount(o<<1, l, r)
	}
	if m < l {
		return t.onesCount(o<<1|1, l, r)
	}
	return t.onesCount(o<<1, l, r) + t.onesCount(o<<1|1, l, r)
}

// 一些等价含义
func (t seg01) getBit(i int) bool  { return t.onesCount(1, i, i) == 1 }
func (t seg01) all0(l, r int) bool { return t.onesCount(1, l, r) == 0 }
func (t seg01) all1(l, r int) bool { return t.onesCount(1, l, r) == r-l+1 }
func (t seg01) index0() int        { return t.next0(1, 1) }
func (t seg01) index1() int        { return t.next1(1, 1) }
func (t seg01) trailingZeros() int { return t.index1() }
func (t seg01) len() int {
	if t[1].state < 0 {
		return 0
	}
	return t.lastIndex1(1)
}

// todo setRange resetRange
