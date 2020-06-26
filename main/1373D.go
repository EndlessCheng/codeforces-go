package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1373D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	f := func(a []int64) int64 {
		if len(a) == 0 {
			return 0
		}
		curS, maxS := a[0], a[0]
		for _, v := range a[1:] {
			curS = max(curS+v, v)
			maxS = max(maxS, curS)
		}
		return max(maxS, 0)
	}
	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		s := int64(0)
		a := make([]int64, n)
		var b, c []int64
		for i := range a {
			Fscan(in, &a[i])
			if i&1 == 0 {
				s += a[i]
				if i > 0 {
					c = append(c, a[i-1]-a[i])
				}
			} else {
				b = append(b, a[i]-a[i-1])
			}
		}
		Fprintln(out, s+max(f(b), f(c)))
	}
}

//func main() { CF1373D(os.Stdin, os.Stdout) }
