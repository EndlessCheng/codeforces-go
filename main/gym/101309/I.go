package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func runI(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, c int
	Fscan(in, &n, &m)
	type nb struct{ to, color int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &c)
		v--
		w--
		g[v] = append(g[v], nb{w, c})
		g[w] = append(g[w], nb{v, c})
	}
	st, end := 0, n-1

	const inf int = 1e9
	dis := make([]int, n)
	for i := range dis {
		dis[i] = inf
	}
	dis[end] = 0
	q := []int{end}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			if w := e.to; dis[v]+1 < dis[w] {
				dis[w] = dis[v] + 1
				q = append(q, w)
			}
		}
	}
	Fprintln(out, dis[st])

	next := []int{st}
	inN := make([]bool, n)
	inN[st] = true
	for loop := dis[st]; loop > 0; loop-- {
		minC := inf
		tmp := next
		next = nil
		for _, v := range tmp {
			for _, e := range g[v] {
				if w, c := e.to, e.color; dis[w] == dis[v]-1 {
					if c < minC {
						for _, w := range next {
							inN[w] = false
						}
						minC, next, inN[w] = c, []int{w}, true
					} else if c == minC && !inN[w] {
						next = append(next, w)
						inN[w] = true
					}
				}
			}
		}
		Fprint(out, minC, " ")
	}
}

//func main() { r, _ := os.Open("ideal.in"); w, _ := os.Create("ideal.out"); runI(r, w) }
