package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1733D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var T, x, y int64
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &x, &y, &s, &t)
		a := []int{}
		for i, c := range s {
			if c != t[i] {
				a = append(a, i)
			}
		}
		m := len(a)
		if m&1 > 0 {
			Fprintln(out, -1)
		} else if m == 0 || y <= x {
			if m == 2 && a[0]+1 == a[1] {
				Fprintln(out, min(y*2, x))
			} else {
				Fprintln(out, int64(m/2)*y)
			}
		} else {
			f, g := int64(0), y
			for i := 1; i < m; i++ {
				f, g = g, min(g+y, f+int64(a[i]-a[i-1])*x*2)
			}
			Fprintln(out, g/2)
		}
	}
}

//func main() { CF1733D2(os.Stdin, os.Stdout) }
