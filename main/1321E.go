package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go

type pair struct{ val, cost int }
type node struct{ l, r, max, todo int }
type seg []node

func (seg) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) _pushUp(o int) { t[o].max = t.max(t[o<<1].max, t[o<<1|1].max) }

func (t seg) _build(a []pair, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].max = -a[l-1].cost
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _spread(o int) {
	if add := t[o].todo; add != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.max += add
		ro.max += add
		lo.todo += add
		ro.todo += add
		t[o].todo = 0
	}
}

func (t seg) _update(o, l, r, add int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].max += add
		t[o].todo += add
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._update(o<<1, l, r, add)
	}
	if m < r {
		t._update(o<<1|1, l, r, add)
	}
	t._pushUp(o)
}

func (t seg) init(a []pair)        { t._build(a, 1, 1, len(a)) }
func (t seg) update(l, r, val int) { t._update(1, l, r, val) }
func (t seg) maxAll() int          { return t[1].max }

func CF1321E(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m, p := read(), read(), read()
	a := make([]pair, n)
	for i := range a {
		a[i] = pair{read(), read()}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].val < a[j].val })
	b := make([]pair, m)
	for i := range b {
		b[i] = pair{read(), read()}
	}
	sort.Slice(b, func(i, j int) bool { return b[i].val < b[j].val })
	t := make(seg, 4*m)
	t.init(b)
	type monster struct{ x, y, coins int }
	monsters := make([]monster, p)
	for i := range monsters {
		monsters[i] = monster{read(), read(), read()}
	}
	sort.Slice(monsters, func(i, j int) bool { return monsters[i].x < monsters[j].x })

	ans := int(-2e9)
	i := 0
	for _, weapon := range a {
		for ; i < p; i++ {
			mst := monsters[i]
			if mst.x >= weapon.val {
				break
			}
			if minArmourI := sort.Search(m, func(i int) bool { return b[i].val > mst.y }); minArmourI < m {
				t.update(minArmourI+1, m, mst.coins)
			}
		}
		if profit := t.maxAll() - weapon.cost; profit > ans {
			ans = profit
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1321E(os.Stdin, os.Stdout) }
