package main

// github.com/EndlessCheng/codeforces-go
type data struct {
	pre, suf, max, size int
	lch, rch            byte
}

type seg []struct {
	l, r int
	data
}

func (t seg) set(o int, ch byte) {
	t[o].lch = ch
	t[o].rch = ch
}

func (t seg) merge(a, b data) data {
	d := data{a.pre, b.suf, max(a.max, b.max), a.size + b.size, a.lch, b.rch}
	if a.rch == b.lch { // 两个区间中间字符相同，可以合并
		if a.suf == a.size {
			d.pre += b.pre
		}
		if b.pre == b.size {
			d.suf += a.suf
		}
		if a.suf != a.size && b.pre != b.size {
			d.max = max(d.max, a.suf+b.pre)
		}
		d.max = max(d.max, max(d.pre, d.suf))
	}
	return d
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].data = t.merge(lo.data, ro.data)
}

func (t seg) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].pre = 1
		t[o].suf = 1
		t[o].max = 1
		t[o].size = 1
		t.set(o, s[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int, ch byte) {
	if t[o].l == t[o].r {
		t.set(o, ch)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, ch)
	} else {
		t.update(o<<1|1, i, ch)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lv := t.query(o<<1, l, r)
	rv := t.query(o<<1|1, l, r)
	return t.merge(lv, rv)
}

func longestRepeating(s, queryCharacters string, queryIndices []int) []int {
	n := len(s)
	t := make(seg, n*4)
	t.build(s, 1, 1, n)
	ans := make([]int, len(queryIndices))
	for i, index := range queryIndices {
		t.update(1, index+1, queryCharacters[i])
		ans[i] = t[1].max
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a}
