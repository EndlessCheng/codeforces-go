package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1208E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n, w, m int
	Fscan(in, &n, &w)
	d := make([]int64, w+1)
	for ; n > 0; n-- {
		Fscan(in, &m)
		a := make([]int64, m)
		for i := range a {
			Fscan(in, &a[i])
		}
		if m*2 <= w {
			pre, suf := int64(0), int64(0)
			for i, v := range a {
				pre = max(pre, v)
				d[i] += pre
				d[i+1] -= pre
				suf = max(suf, a[m-1-i])
				d[w-1-i] += suf
				d[w-i] -= suf
			}
			d[m] += suf
			d[w-m] -= suf
		} else {
			sz := w - m + 1
			q := []int{}
			for i, v := range a {
				for len(q) > 0 && v >= a[q[len(q)-1]] {
					q = q[:len(q)-1]
				}
				q = append(q, i)
				if q[0] <= i-sz {
					q = q[1:]
				}
				mx := a[q[0]]
				if mx < 0 && i < w-m {
					mx = 0
				}
				d[i] += mx
				d[i+1] -= mx
			}
			suf := int64(0)
			for i := 0; i < w-m; i++ {
				suf = max(suf, a[m-1-i])
				d[w-1-i] += suf
				d[w-i] -= suf
			}
		}
	}
	s := int64(0)
	for _, v := range d[:w] {
		s += v
		Fprint(out, s, " ")
	}
}

//func main() { CF1208E(os.Stdin, os.Stdout) }
