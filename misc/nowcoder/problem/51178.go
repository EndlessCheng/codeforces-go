package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func nc51178(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var n, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &w, &v)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var f func(int, int) (int, int)
	f = func(v, fa int) (notChosen, chosen int) {
		chosen = a[v]
		for _, w := range g[v] {
			if w != fa {
				nc, c := f(w, v)
				notChosen += max(nc, c)
				chosen += nc
			}
		}
		return
	}
	nc, c := f(0, -1)
	Fprint(out, max(nc, c))
}

func main() { nc51178(os.Stdin, os.Stdout) }
