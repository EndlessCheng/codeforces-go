package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1172B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod = 998244353
	const mx = 2e5
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	ans := int64(n)
	for _, vs := range g {
		ans = ans * F[len(vs)] % mod
	}
	Fprint(out, ans)
}

//func main() { CF1172B(os.Stdin, os.Stdout) }
