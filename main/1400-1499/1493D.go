package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1493D(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
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

	const mx int = 2e5
	lpf := [mx + 1]int{}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	n, q := r(), r()
	cnts := make([]map[int]int, n)
	for i := range cnts {
		cnts[i] = map[int]int{}
	}
	cc := [mx + 1][]int{}
	g := int64(1)
	add := func(i, v int) {
		cnt := cnts[i]
		for v > 1 {
			for p := lpf[v]; lpf[v] == p; v /= p {
				c := cnt[p]
				cnt[p]++
				if len(cc[p]) <= c {
					cc[p] = append(cc[p], 0)
				}
				if cc[p][c]++; cc[p][c] == n {
					g = g * int64(p) % (1e9 + 7)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		add(i, r())
	}
	for ; q > 0; q-- {
		add(r()-1, r())
		Fprintln(out, g)
	}
}

//func main() { CF1493D(os.Stdin, os.Stdout) }
