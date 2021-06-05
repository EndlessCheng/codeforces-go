package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1239D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			if v != w {
				g[v-1] = append(g[v-1], w-1)
			}
		}
		sccID := make([]int, n)
		sid := 0
		c0 := 0
		dfn := make([]int, n)
		dfsClock := 0
		s := []int{}
		inS := make([]bool, n)
		var f func(int) int
		f = func(v int) int {
			dfsClock++
			dfn[v] = dfsClock
			lowV := dfsClock
			s = append(s, v)
			inS[v] = true
			for _, w := range g[v] {
				if dfn[w] == 0 {
					lowV = min(lowV, f(w))
				} else if inS[w] {
					lowV = min(lowV, dfn[w])
				}
			}
			if dfn[v] == lowV {
				for {
					w := s[len(s)-1]
					s = s[:len(s)-1]
					inS[w] = false
					sccID[w] = sid
					if sid == 0 {
						c0++
					}
					if w == v {
						break
					}
				}
				sid++
			}
			return lowV
		}
		for v, ts := range dfn {
			if ts == 0 {
				f(v)
			}
		}
		if sid == 1 {
			Fprintln(out, "No")
			continue
		}
		Fprintln(out, "Yes")
		Fprintln(out, c0, n-c0)
		for i, id := range sccID {
			if id == 0 {
				Fprint(out, i+1, " ")
			}
		}
		Fprintln(out)
		for i, id := range sccID {
			if id > 0 {
				Fprint(out, i+1, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1239D(os.Stdin, os.Stdout) }
