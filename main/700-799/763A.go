package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go

// 巧妙的写法
func CF763A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, cnt int
	Fscan(in, &n)
	es := make([][2]int, n-1)
	for i := range es {
		Fscan(in, &es[i][0], &es[i][1])
	}
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}
	diff := make([]int, n)
	for _, e := range es {
		if v, w := e[0]-1, e[1]-1; c[v] != c[w] {
			diff[v]++
			diff[w]++
			cnt++
		}
	}
	for i, d := range diff {
		if d == cnt {
			Fprint(out, "YES\n", i+1)
			return
		}
	}
	Fprint(out, "NO")
}

// 换根 DP 写法
func CF763Adp(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}

	same := make([]bool, n)
	for i := range same {
		same[i] = true
	}
	var f func(v, fa int) bool
	f = func(v, fa int) bool {
		for _, w := range g[v] {
			if w != fa {
				if !f(w, v) || c[w] != c[v] {
					same[v] = false
				}
			}
		}
		return same[v]
	}
	f(0, -1)

	ans := -1
	var f2 func(v, fa int)
	f2 = func(v, fa int) {
		cntDiff, allSame := 0, fa < 0 || c[v] == c[fa]
		for _, w := range g[v] {
			if w != fa {
				if !same[w] {
					cntDiff++
				} else if c[w] != c[v] {
					allSame = false
				}
			}
		}
		if cntDiff == 0 {
			ans = v
			return
		}
		if !allSame || cntDiff > 1 {
			return
		}
		for _, w := range g[v] {
			if w != fa {
				if !same[w] {
					f2(w, v)
					break
				}
			}
		}
	}
	f2(0, -1)
	if ans < 0 {
		Fprint(out, "NO")
	} else {
		Fprint(out, "YES\n", ans+1)
	}
}

//func main() { CF763A(os.Stdin, os.Stdout) }
