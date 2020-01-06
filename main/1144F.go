package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1144F(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m := read(), read()
	g := make([][]int, n)
	type pair struct{ v, w int }
	edges := make([]pair, m)
	for i := 0; i < m; i++ {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		edges[i] = pair{v, w}
	}

	colors := make([]int8, n)
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		colors[v] = c
		for _, w := range g[v] {
			if colors[w] == c {
				return false
			}
			if colors[w] == 0 {
				if !f(w, 3-c) {
					return false
				}
			}
		}
		return true
	}
	for i, c := range colors {
		if c == 0 && !f(i, 1) {
			Fprintln(out, "NO")
			return
		}
	}
	Fprintln(out, "YES")
	for _, e := range edges {
		Fprint(out, colors[e.w]-1)
	}
}

//func main() {
//	Sol1144F(os.Stdin, os.Stdout)
//}
