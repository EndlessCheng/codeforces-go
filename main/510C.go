package main

import (
	"bufio"
	. "fmt"
	"io"
)

type graph struct {
	size     int
	edges    [][]int
	inDegree []int
}

func newGraph(size int) *graph {
	return &graph{
		size:     size,
		edges:    make([][]int, size),
		inDegree: make([]int, size),
	}
}

func (g *graph) add(from, to int) {
	g.edges[from] = append(g.edges[from], to)
	g.inDegree[to]++
}

func (g *graph) topSort() (order []int, acyclic bool) {
	queue := []int{}
	for i := 0; i < g.size; i++ {
		if g.inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	var v int
	for len(queue) > 0 {
		v, queue = queue[0], queue[1:]
		order = append(order, v)
		for _, w := range g.edges[v] {
			g.inDegree[w]--
			if g.inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	return order, len(order) == g.size
}

// github.com/EndlessCheng/codeforces-go
func Sol510C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	g := newGraph(26)
	var n int
	Fscan(in, &n)
	var prev, cur string
	for ; n > 0; n-- {
		Fscan(in, &cur)
		minL := len(prev)
		if len(cur) < minL {
			minL = len(cur)
		}
		ok := false
		for i := 0; i < minL; i++ {
			if v, w := prev[i]-'a', cur[i]-'a'; v != w {
				g.add(int(v), int(w))
				ok = true
				break
			}
		}
		if !ok && len(prev) > len(cur) {
			Fprint(out, "Impossible")
			return
		}
		prev = cur
	}

	order, acyclic := g.topSort()
	if !acyclic {
		Fprint(out, "Impossible")
		return
	}
	done := make([]bool, 26)
	for _, c := range order {
		Fprintf(out, "%c", 'a'+c)
		done[c] = true
	}
	for i, d := range done {
		if !d {
			Fprintf(out, "%c", 'a'+i)
		}
	}
}

//func main() {
//	Sol510C(os.Stdin, os.Stdout)
//}
