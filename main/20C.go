package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
	"os"
)

type pair struct {
	x int64
	y int
}
type pairHeap []pair

func (h pairHeap) Len() int              { return len(h) }
func (h pairHeap) Less(i, j int) bool    { return h[i].x < h[j].x }
func (h pairHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *pairHeap) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

type neighbor struct {
	vertex int
	weight int
}

type graph struct {
	size    int
	edges   [][]neighbor
	visited []bool
}

func (g *graph) add(from, to int, weight int) {
	g.edges[from] = append(g.edges[from], neighbor{to, weight})
}

func (g *graph) addBoth(from, to int, weight int) {
	g.add(from, to, weight)
	if from != to {
		g.add(to, from, weight)
	}
}

func (g *graph) shortestPaths(start int) (dist []int64, parents []int) {
	const inf int64 = 1e18
	dist = make([]int64, g.size+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[start] = 0
	parents = make([]int, g.size+1)
	for i := range parents {
		parents[i] = -1
	}

	h := &pairHeap{}
	Push(h, pair{0, start})
	for h.Len() > 0 {
		p := Pop(h).(pair)
		v := p.y
		if g.visited[v] {
			continue
		}
		g.visited[v] = true
		for _, e := range g.edges[v] {
			w := e.vertex
			if newDist := dist[v] + int64(e.weight); newDist < dist[w] {
				dist[w] = newDist
				parents[w] = v
				Push(h, pair{newDist, w})
			}
		}
	}
	return
}

// github.com/EndlessCheng/codeforces-go
func Sol20C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	g := &graph{
		size:    n,
		edges:   make([][]neighbor, n+1),
		visited: make([]bool, n+1),
	}
	for i := 0; i < m; i++ {
		var v, w, weight int
		Fscan(in, &v, &w, &weight)
		g.addBoth(v, w, weight)
	}

	dist, parents := g.shortestPaths(1)
	const inf int64 = 1e18
	if dist[n] == inf {
		Fprint(out, -1)
		return
	}
	path := []int{}
	for v := n; v != -1; v = parents[v] {
		path = append(path, v)
	}
	for i := len(path) - 1; i >= 0; i-- {
		Fprint(out, path[i], " ")
	}
}

func main() {
	Sol20C(os.Stdin, os.Stdout)
}
