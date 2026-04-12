package main

import "math/bits"

// https://space.bilibili.com/206214
var targetGcd int

type seg []struct{ l, r, gcd int }

func (t seg) maintain(o int) {
	t[o].gcd = gcd(t[o<<1].gcd, t[o<<1|1].gcd)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		if a[l]%targetGcd == 0 {
			t[o].gcd = a[l]
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, val int) {
	cur := &t[o]
	if cur.l == cur.r {
		if val%targetGcd == 0 {
			cur.gcd = val
		} else {
			cur.gcd = 0
		}
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) int {
	if l > r {
		return 0
	}
	if l <= t[o].l && t[o].r <= r {
		return t[o].gcd
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return gcd(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func (t seg) check(n int) bool {
	for i := range n {
		if gcd(t.query(1, 0, i-1), t.query(1, i+1, n-1)) == targetGcd {
			return true
		}
	}
	return false
}

func countGoodSubseq(nums []int, p int, queries [][]int) (ans int) {
	targetGcd = p
	cntP := 0
	for _, x := range nums {
		if x%targetGcd == 0 {
			cntP++
		}
	}

	n := len(nums)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(nums, 1, 0, n-1)

	for _, q := range queries {
		i, x := q[0], q[1]

		if nums[i]%p == 0 {
			cntP--
		}
		if x%p == 0 {
			cntP++
		}
		nums[i] = x
		t.update(1, q[0], x)

		if t[1].gcd == p && (cntP < n || n > 7 || t.check(n)) {
			ans++
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
