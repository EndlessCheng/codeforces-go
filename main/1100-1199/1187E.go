package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1187E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	size := make([]int, n+1)
	sum := make([]int64, n+1)
	var f func(v, fa int) (int, int64)
	f = func(v, fa int) (sz int, s int64) {
		sz = 1
		for _, w := range g[v] {
			if w != fa {
				subSz, subSum := f(w, v)
				sz += subSz
				s += subSum
			}
		}
		size[v] = sz
		s += int64(sz)
		sum[v] = s
		return
	}
	f(1, 0)

	ans := int64(0)
	var reroot func(v, fa int, othersSum int64)
	reroot = func(v, fa int, othersSum int64) {
		for _, w := range g[v] {
			if w != fa {
				// 涂黑 w 后，「父子树」的贡献为涂黑父节点的贡献 n-size[w]，加上涂黑父节点的非 w 子树的贡献 sum[v]-int64(size[v])-sum[w]
				reroot(w, v, othersSum+int64(n-size[w])+sum[v]-int64(size[v])-sum[w])
			}
		}
		if s := int64(n) + othersSum + sum[v] - int64(size[v]); s > ans {
			ans = s
		}
	}
	reroot(1, 0, 0)
	Fprint(_w, ans)
}

//func main() { CF1187E(os.Stdin, os.Stdout) }
