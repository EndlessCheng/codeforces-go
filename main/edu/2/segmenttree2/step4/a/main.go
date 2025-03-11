package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type pair struct{ ch, add int }

type seg []struct {
	l, r int
	sum  int
	todo pair
}

func mergeInfo(a, b int) int {
	return a + b
}

var todoInit = pair{1e9, 0}

func mergeTodo(f, old pair) pair {
	if f.ch != 1e9 {
		return f
	}
	old.add += f.add
	return old
}

func (t seg) apply(o int, f pair) {
	cur := &t[o]

	sz := cur.r - cur.l + 1
	if f.ch == 1e9 {
		cur.sum += f.add * sz
	} else {
		cur.sum = (f.ch + f.add) * sz
	}

	cur.todo = mergeTodo(f, cur.todo)
}

func (t seg) maintain(o int) {
	t[o].sum = mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t seg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r int, f pair) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	for range q {
		Fscan(in, &op, &l, &r)
		r--
		if op == 1 {
			Fscan(in, &v)
			t.update(1, l, r, pair{v, 0})
		} else if op == 2 {
			Fscan(in, &v)
			t.update(1, l, r, pair{1e9, v})
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
