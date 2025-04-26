package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	deg := make([]int, n)
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		g[v-1] = append(g[v-1], w-1)
		deg[w-1]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	f := make([]int, n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		f[v]++
		ans = max(ans, f[v])
		for _, w := range g[v] {
			f[w] = max(f[w], f[v])
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	Fprint(out, ans-1)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
