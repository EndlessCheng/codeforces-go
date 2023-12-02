package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3594(_r io.Reader, out io.Writer) {
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, lim, d := r(), r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}

	ans := d
	sd := 0
	for _, v := range a[:d] {
		sd += v
	}
	type pair struct{ l, s int }
	q := []pair{{0, sd}}
	s := sd
	left := 0
	for i := d; i < n; i++ {
		sd += a[i] - a[i-d]
		for len(q) > 0 && sd >= q[len(q)-1].s {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i - d + 1, sd})

		s += a[i]
		for s-q[0].s > lim {
			s -= a[left]
			left++
			for q[0].l < left {
				q = q[1:]
			}
		}

		ans = max(ans, i-left+1)
	}
	Fprint(out, ans)
}

//func main() { p3594(os.Stdin, os.Stdout) }
