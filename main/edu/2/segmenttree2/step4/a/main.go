package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct {
	l, r                int
	setTodo, addTodo, s int64
}

func (t seg) maintain(o int) {
	t[o].s = t[o<<1].s + t[o<<1|1].s
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].setTodo = l, r, -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

// 如果区间已经有了赋值标记：
// - 赋值操作 更新 赋值标记
// - 加操作 更新 赋值标记
// 如果区间已经有了加标记：
// - 赋值操作 更新 赋值标记，清空 加标记
// - 加操作 更新 加标记
// 我们发现，赋值标记和加标记并不会共存
func (t seg) doSet(o int, v int64) {
	to := &t[o]
	to.setTodo = v
	to.addTodo = 0
	to.s = int64(to.r-to.l+1) * v
}
func (t seg) doAdd(o int, v int64) {
	to := &t[o]
	if to.setTodo < 0 {
		to.addTodo += v
	} else {
		to.setTodo += v
	}
	to.s += int64(to.r-to.l+1) * v
}
func (t seg) spread(o int) {
	if v := t[o].setTodo; v >= 0 {
		t.doSet(o<<1, v)
		t.doSet(o<<1|1, v)
		t[o].setTodo = -1
	} else if v := t[o].addTodo; v != 0 {
		t.doAdd(o<<1, v)
		t.doAdd(o<<1|1, v)
		t[o].addTodo = 0
	}
}

func (t seg) update(o, l, r, v int, set bool) {
	if l <= t[o].l && t[o].r <= r {
		if set {
			t.doSet(o, int64(v))
		} else {
			t.doAdd(o, int64(v))
		}
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v, set)
	}
	if m < r {
		t.update(o<<1|1, l, r, v, set)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].s
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op < 3 {
			Fscan(in, &v)
			t.update(1, l+1, r, v, op == 1)
		} else {
			Fprintln(out, t.query(1, l+1, r))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
