package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1592C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		xor := 0
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			xor ^= a[i]
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		var f func(int, int) int
		f = func(v, fa int) (cntXor int) {
			for _, w := range g[v] {
				if w != fa {
					cntXor += f(w, v)
					a[v] ^= a[w]
				}
			}
			if cntXor > 0 {
				if a[v] == 0 { // 子树有 xor 的情况下，可以额外再弄个 xor
					cntXor++
				}
			} else if a[v] == xor {
				cntXor++
			}
			return
		}
		if xor == 0 || k > 2 && f(0, -1) > 1 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1592C(os.Stdin, os.Stdout) }
