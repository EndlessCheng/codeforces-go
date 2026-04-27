package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg14 []struct{ l, r, gcd int }

func (t seg14) maintain(o int) {
	t[o].gcd = gcd14(t[o<<1].gcd, t[o<<1|1].gcd)
}

func (t seg14) build(in io.Reader, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		Fscan(in, &t[o].gcd)
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg14) update(o, i, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.gcd = v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t.maintain(o)
}

func (t seg14) query(o, l, r, v int) int {
	if t[o].gcd%v == 0 {
		return 0
	}
	if t[o].l == t[o].r {
		return 1
	}
	m := (t[o].l + t[o].r) >> 1
	cnt := 0
	if l <= m {
		cnt += t.query(o<<1, l, r, v)
	}
	if cnt < 2 && r > m {
		cnt += t.query(o<<1|1, l, r, v)
	}
	return cnt
}

func cf914D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, v int
	Fscan(in, &n)
	t := make(seg14, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 1, n)

	Fscan(in, &q)
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 2 {
			t.update(1, l, r)
			continue
		}
		Fscan(in, &v)
		if t.query(1, l, r, v) < 2 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf914D(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd14(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
