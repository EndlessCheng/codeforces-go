package main

import (
	"bufio"
	. "fmt"
	"io"
)

type graph623A struct {
	size   int
	edges  [][]bool
	degree []int
	color  []int
}

func (g *graph623A) add(from, to int) {
	g.edges[from][to] = true
	g.degree[from]++
}

func (g *graph623A) addBoth(from, to int) {
	g.add(from, to)
	if from != to {
		g.add(to, from)
	}
}

func (g *graph623A) _isBipartite(v int) bool {
	for w, e := range g.edges[v] {
		if e || w == v {
			continue
		}
		if g.color[w] == g.color[v] {
			return false
		}
		if g.color[w] == 0 {
			g.color[w] = 3 - g.color[v]
			if !g._isBipartite(w) {
				return false
			}
		}
	}
	return true
}

func (g *graph623A) isBipartite() bool {
	checked := false
	cnt := 0
	for i, deg := range g.degree {
		deg = g.size - 1 - deg
		if deg > 0 {
			if checked {
				if g.color[i] == 0 {
					return false
				}
				continue
			}
			g.color[i] = 1
			if !g._isBipartite(i) {
				return false
			}
			for w := range g.edges[i] {
				if g.color[w] == 2 {
					cnt++
				}
			}
			checked = true
		}
	}
	if cnt > 0 {
		for v, c := range g.color {
			if c == 1 {
				cntW := 0
				for w, e := range g.edges[v] {
					if e || w == v {
						continue
					}
					if g.color[w] == 2 {
						cntW++
					}
				}
				if cntW != cnt {
					return false
				}
			}
		}
	}
	return true
}

// github.com/EndlessCheng/codeforces-go
func Sol623A(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	g := &graph623A{
		size:   n,
		edges:  make([][]bool, n),
		degree: make([]int, n),
		color:  make([]int, n),
	}
	for i := range g.edges {
		g.edges[i] = make([]bool, n)
	}
	for ; m > 0; m-- {
		var v, w int
		Fscan(in, &v, &w)
		g.addBoth(v-1, w-1)
	}

	// 若反图为空或二分图，则存在
	if !g.isBipartite() {
		Fprint(out, "No")
		return
	}

	Fprintln(out, "Yes")
	for _, c := range g.color {
		Fprintf(out, "%c", "bac"[c])
	}
}

//func main() {
//	Sol623A(os.Stdin, os.Stdout)
//}
