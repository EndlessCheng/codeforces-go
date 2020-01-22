package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1118F1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	colors := make([]int8, n)
	types := [3]int{}
	for i := range colors {
		Fscan(in, &colors[i])
		types[colors[i]]++
	}
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	cnts := make([][3]int, n)
	var f func(int, int) [3]int
	f = func(v, p int) [3]int {
		cnts[v][colors[v]] = 1
		for _, w := range g[v] {
			if w != p {
				c := f(w, v)
				cnts[v][1] += c[1]
				cnts[v][2] += c[2]
			}
		}
		return cnts[v]
	}
	f(0, -1)
	ans := 0
	for _, c := range cnts[1:] {
		if c[1] == types[1] && c[2] == 0 || c[1] == 0 && c[2] == types[2] {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() {
//	CF1118F1(os.Stdin, os.Stdout)
//}
