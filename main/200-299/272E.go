package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF272E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]byte, n)
	q := make([]int, n)
	inQ := make([]bool, n)
	for i := range ans {
		ans[i] = '0'
		q[i] = i
		inQ[i] = true
	}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		inQ[v] = false
		c := byte(0)
		for _, w := range g[v] {
			c += ans[v] ^ ans[w] ^ 1 // v w 同部
		}
		if c > 1 {
			ans[v] ^= 1
			for _, w := range g[v] {
				if !inQ[w] {
					inQ[w] = true
					q = append(q, w)
				}
			}
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF272E(os.Stdin, os.Stdout) }
