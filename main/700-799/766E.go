package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF766E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	// 也可以 DFS 20 次，这样可以把空间减少为 1/20
	ans := int64(0)
	var f func(v, fa int) [20][2]int
	f = func(v, fa int) (cnt [20][2]int) {
		av := a[v]
		ans += int64(av)
		for i := 0; i < 20; i++ {
			cnt[i][av>>i&1] = 1
		}
		for _, w := range g[v] {
			if w != fa {
				for i, p := range f(w, v) {
					ans += (int64(cnt[i][0])*int64(p[1]) + int64(cnt[i][1])*int64(p[0])) * int64(1<<i)
					cnt[i][av>>i&1] += p[0]
					cnt[i][av>>i&1^1] += p[1]
				}
			}
		}
		return
	}
	f(0, -1)
	Fprint(out, ans)
}

//func main() { CF766E(os.Stdin, os.Stdout) }
