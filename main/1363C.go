package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1363C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var t, n, x, v, w int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &x)
		x--
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		if len(g[x]) <= 1 || n&1 == 0 {
			Fprint(out, "Ayush")
		} else {
			Fprint(out, "Ashish")
		}
	}
}

//func main() { CF1363C(os.Stdin, os.Stdout) }
