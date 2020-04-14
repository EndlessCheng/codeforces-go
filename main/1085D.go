package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1085D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, s, v, w int
	Fscan(in, &n, &s)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	c := 0
	for _, vs := range g[1:] {
		if len(vs) == 1 {
			c++
		}
	}
	Fprintf(_w, "%.18f", float64(2*s)/float64(c))
}

//func main() { CF1085D(os.Stdin, os.Stdout) }
