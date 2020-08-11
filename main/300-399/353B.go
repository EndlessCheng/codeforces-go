package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF353B(in io.Reader, out io.Writer) {
	var n, c1, c2 int
	Fscan(in, &n)
	a := make([]int, 2*n)
	pos := [101][]int{}
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}

	for _, ps := range pos {
		if len(ps) == 1 {
			c1++
		} else if len(ps) > 1 {
			c2++
		}
	}
	Fprintln(out, (c2+c1/2)*(c2+(c1+1)/2))
	tp, tp2 := 0, c1&1
	for _, ps := range pos {
		if len(ps) == 1 {
			a[ps[0]] = tp
			tp ^= 1
		} else if len(ps) > 1 {
			a[ps[0]] = 0
			a[ps[1]] = 1
			for _, p := range ps[2:] {
				a[p] = tp2
				tp2 ^= 1
			}
		}
	}
	for _, v := range a {
		Fprint(out, v+1, " ")
	}
}

//func main() { CF353B(os.Stdin, os.Stdout) }
