package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type node14 struct {
	lo, ro  *node14
	l, r, s int
}

func (o *node14) update(i int) {
	o.s++
	if o.l+1 == o.r {
		return
	}
	m := (o.l + o.r) >> 1
	ls := m - o.l
	if o.lo != nil {
		ls -= o.lo.s
	}
	if i < ls {
		if o.lo == nil {
			o.lo = &node14{l: o.l, r: m}
		}
		o.lo.update(i)
	} else {
		if o.ro == nil {
			o.ro = &node14{l: m, r: o.r}
		}
		o.ro.update(i - ls)
	}
}

func (o *node14) query(r int) int {
	if o == nil || r <= o.l {
		return 0
	}
	if o.r <= r {
		return o.s
	}
	return o.lo.query(r) + o.ro.query(r)
}

func cf1614E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_001

	seg := &node14{r: mod}
	var n, ans, t, k, x, base int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &t, &k)
		if t >= base {
			seg.update(t - base)
		} else {
			base--
		}
		if t > base {
			seg.update(t - 1 - base)
			base++
		}
		for ; k > 0; k-- {
			Fscan(in, &x)
			x = (x + ans) % mod
			ans = base + x - seg.query(x)
			Fprintln(out, ans)
		}
	}
}

//func main() { cf1614E(os.Stdin, os.Stdout) }
