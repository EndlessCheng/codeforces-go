package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1510G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		fa := make([]int, n)
		g := make([][]int, n)
		dep := make([]int, n)
		mxD := 0
		for i := 1; i < n; i++ {
			Fscan(in, &fa[i])
			fa[i]--
			g[fa[i]] = append(g[fa[i]], i)
			dep[i] = dep[fa[i]] + 1
			if dep[i] > mxD {
				mxD = dep[i]
			}
		}
		if k <= mxD {
			mxD = k - 1
		}
		end := 1
		for i, d := range dep {
			if d == mxD {
				end = i
				break
			}
		}
		onPath := make([]bool, n)
		for v := end; v > 0; v = fa[v] {
			onPath[v] = true
		}

		Fprintln(out, 2*(k-1)-mxD)
		k -= mxD
		var f func(int)
		f = func(v int) {
			Fprint(out, v+1, " ")
			if !onPath[v] {
				k--
			}
			for _, w := range g[v] {
				if onPath[w] {
					defer f(w)
				} else if k > 0 {
					f(w)
					Fprint(out, v+1, " ")
				}
			}
		}
		f(0)
		Fprintln(out)
	}
}

//func main() { CF1510G(os.Stdin, os.Stdout) }
