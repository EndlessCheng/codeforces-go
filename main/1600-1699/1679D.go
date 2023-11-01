package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1679D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, w int
	Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v-1] = append(g[v-1], w-1)
	}

	ans := sort.Search(1e9+1, func(mx int) bool {
		deg := make([]int, n)
		left := n
		for v, ws := range g {
			if a[v] > mx {
				left--
				continue
			}
			for _, w := range ws {
				if a[w] <= mx {
					deg[w]++
				}
			}
		}
		q := []int{}
		for i, d := range deg {
			if d == 0 && a[i] <= mx {
				q = append(q, i)
			}
		}
		f := make([]int, n)
		mxF := 0
		for len(q) > 0 {
			left--
			v := q[0]
			q = q[1:]
			f[v]++
			mxF = max(mxF, f[v])
			for _, w := range g[v] {
				if a[w] > mx {
					continue
				}
				f[w] = f[v]
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		return left > 0 || mxF >= k
	})
	if ans > 1e9 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1679D(os.Stdin, os.Stdout) }
