package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1358D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}

	var n int
	var x, ans, s, s2 int64
	Fscan(in, &n, &x)
	d := make([]int64, n)
	for i := range d {
		Fscan(in, &d[i])
	}
	d = append(d, d...)
	l := 0
	for _, v := range d {
		s += v
		s2 += (v + 1) * v / 2
		for s > x {
			s -= d[l]
			s2 -= (d[l] + 1) * d[l] / 2
			l++
		}
		if l > 0 {
			c := x - s
			ans = max(ans, s2+(d[l-1]*2-c+1)*c/2)
		} else {
			ans = max(ans, s2)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1358D(os.Stdin, os.Stdout) }
