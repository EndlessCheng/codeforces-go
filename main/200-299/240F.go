package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg40 []struct {
	l, r int
	cnt  [26]int
	todo byte
}

const todoInit40 byte = 26

func (seg40) mergeInfo(a, b [26]int) [26]int {
	for i := range a {
		a[i] += b[i]
	}
	return a
}

func (t seg40) do(O int, v byte) {
	o := &t[O]
	o.cnt = [26]int{}
	o.cnt[v] = o.r - o.l + 1
	o.todo = v
}

func (t seg40) spread(o int) {
	if v := t[o].todo; v != todoInit40 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit40
	}
}

func (t seg40) build(a []byte, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit40
	if l == r {
		t[o].cnt[a[l-1]-'a'] = 1
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg40) maintain(o int) {
	t[o].cnt = t.mergeInfo(t[o<<1].cnt, t[o<<1|1].cnt)
}

func (t seg40) update(o, l, r int, v byte) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg40) query(o, l, r int) [26]int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].cnt
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func (t seg40) spreadAll(o int, s []byte) {
	if t[o].l == t[o].r {
		for i, c := range t[o].cnt {
			if c > 0 {
				s[t[o].l-1] = 'a' + byte(i)
				break
			}
		}
		return
	}
	t.spread(o)
	t.spreadAll(o<<1, s)
	t.spreadAll(o<<1|1, s)
}

func cf240F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, l, r int
	var s []byte
	Fscan(in, &n, &m, &s)
	t := make(seg40, 2<<bits.Len(uint(n-1)))
	t.build(s, 1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		cnt := t.query(1, l, r)
		odd := 0
		for _, c := range cnt {
			odd += c % 2
		}
		if odd > 1 {
			continue
		}
		for i, c := range cnt {
			if c == 0 {
				continue
			}
			h := c / 2
			if h > 0 {
				t.update(1, l, l+h-1, byte(i))
				t.update(1, r-h+1, r, byte(i))
				l += h
				r -= h
			}
			if c%2 > 0 {
				m := (l + r) / 2
				t.update(1, m, m, byte(i))
			}
		}
	}
	t.spreadAll(1, s)
	Fprintf(out, "%s", s)
}

//func main() { r, _ := os.Open("input.txt"); w, _ := os.Create("output.txt"); cf240F(r, w) }
