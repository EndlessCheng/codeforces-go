package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
type data80 struct{ min, minI int }
type seg80 []struct {
	l, r int
	data80
}

func (seg80) merge(a, b data80) data80 {
	if a.min < b.min {
		return a
	}
	return b
}

func (t seg80) build(v, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].min, t[o].minI = v, l
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(v, o<<1, l, m)
	t.build(v, o<<1|1, m+1, r)
}

func (t seg80) update(o, i, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.min = v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t[o].data80 = t.merge(t[o<<1].data80, t[o<<1|1].data80)
}

func (t seg80) query(o, l, r int) data80 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data80
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.merge(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf780G(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var h, w, n, ans int
	Fscan(in, &h, &w, &n)
	type tuple struct{ u, l, r, s int }
	a := make([]tuple, n)
	for i := range a {
		Fscan(in, &a[i].u, &a[i].l, &a[i].r, &a[i].s)
	}
	slices.SortFunc(a, func(a, b tuple) int { return b.u - a.u })

	segT := make(seg80, 2<<bits.Len(uint(w-1)))
	segT.build(h+1, 1, 1, w)
	type pair struct{ h, num int }
	stk := make([][]pair, w+1)
	for i := 1; i <= w; i++ {
		stk[i] = []pair{{h + 1, 1}}
	}

	for _, t := range a {
		cnt := 0
		for {
			d := segT.query(1, t.l, t.r)
			if d.min > t.u+t.s {
				break
			}

			i := d.minI
			p := stk[i]
			cnt += p[len(p)-1].num
			p = p[:len(p)-1]
			stk[i] = p

			if len(p) > 0 {
				h = p[len(p)-1].h
			} else {
				h = 3e9
			}
			segT.update(1, i, h)
		}
		if cnt == 0 {
			continue
		}
		if t.l == 1 || t.r == w {
			cnt *= 2
		}
		if l := t.l - 1; l > 0 {
			stk[l] = append(stk[l], pair{t.u, cnt % mod})
			segT.update(1, l, t.u)
		}
		if r := t.r + 1; r <= w {
			stk[r] = append(stk[r], pair{t.u, cnt % mod})
			segT.update(1, r, t.u)
		}
	}
	for _, ps := range stk {
		for _, p := range ps {
			ans += p.num
		}
	}
	Fprint(out, ans%mod)
}

//func main() { cf780G(bufio.NewReader(os.Stdin), os.Stdout) }
