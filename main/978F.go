package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF978F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w int
	Fscan(in, &n, &m)
	a := make([]int, n)
	g := make([][]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], a[w])
		g[w] = append(g[w], a[v])
	}
	for _, vs := range g {
		sort.Ints(vs)
	}
	b := make([]int, n)
	copy(b, a)
	sort.Ints(b)
	for i, v := range a {
		Fprint(out, sort.SearchInts(b, v)-sort.SearchInts(g[i], v), " ")
	}
}

//func main() { CF978F(os.Stdin, os.Stdout) }
