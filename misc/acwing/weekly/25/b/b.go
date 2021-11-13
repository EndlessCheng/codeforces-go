package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]bool, n+1)
	for i := range g {
		g[i] = make([]bool, n+1)
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		g[v][w] = true
		g[w][v] = true
	}

	dis := make([]int, n+1)
	q := []int{1}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		if v == n {
			Fprint(out, dis[n])
			return
		}
		for w := 2; w <= n; w++ {
			if dis[w] == 0 && g[v][w] != g[1][n] {
				q = append(q, w)
				dis[w] = dis[v] + 1
			}
		}
	}
	Fprint(out, -1)
}

func main() { run(os.Stdin, os.Stdout) }
