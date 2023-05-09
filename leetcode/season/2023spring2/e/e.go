package main

// https://space.bilibili.com/206214
var k int

type seg []struct {
	l, r int
	to   [2]int
}

func (t seg) set(o, val int) {
	t[o].to[1] = t[o].to[0] ^ val
}

func (t seg) maintain(o int) {
	a, b, c := t[o<<1].to, t[o<<1|1].to, [2]int{}
	for i := 0; i < k; i++ {
		c[0] |= b[a[0]>>i&1] >> i & 1 << i
		c[1] |= b[a[1]>>i&1] >> i & 1 << i
	}
	t[o].to = c
}

func (t seg) build(a []int, k, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].to[0] = 1<<k - 1
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, k, o<<1, l, m)
	t.build(a, k, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func getNandResult(K int, arr []int, operations [][]int) (ans int) {
	k = K
	t := make(seg, len(arr)*4)
	t.build(arr, k, 1, 1, len(arr))
	for _, op := range operations {
		if op[0] == 0 {
			t.update(1, op[1]+1, op[2])
			continue
		}
		to := t[1].to
		x, y := op[1], op[2]
		for i := 0; i < k; i++ {
			var res int
			y := y >> i & 1
			y1 := to[y] >> i & 1 // 穿过 arr 一次
			if y1 == y { // 不变
				res = y
			} else if x == 1 || to[y1]>>i&1 == y1 {
				// 只穿过一次，或者穿过两次和穿过一次相同
				res = y1
			} else {
				res = y ^ x%2 // 奇变偶不变
			}
			ans ^= res << i
		}
	}
	return
}
