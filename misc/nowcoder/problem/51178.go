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
	son := make([]bool, n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &w, &v)
		son[w-1] = true
		g[v-1] = append(g[v-1], w-1)
	}

	rt := 0
	for ; rt < n && son[rt]; rt++ {
	}
	var f func(int) (int, int)
	f = func(v int) (notChosen, chosen int) {
		chosen = a[v]
		for _, w := range g[v] {
			nc, c := f(w)
			notChosen += max(nc, c)
			chosen += nc
		}
		return
	}
	nc, c := f(rt)
	Fprint(out, max(nc, c))
}

func main() { nc51178(os.Stdin, os.Stdout) }
