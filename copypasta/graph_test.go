package copypasta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_graph(t *testing.T) {
	n := 10
	g := newGraph(n, 0)
	g.addBoth(1, 2, 1)
	g.addBoth(2, 3, 1)
	g.addBoth(3, 4, 1)
	g.addBoth(3, 5, 1)
	g.addBoth(5, 6, 1)

	calc := func(start int) (anotherStart int, maxPath int) {
		const inf = 1e9
		dist := make([]int, n+1)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		g.resetStates()
		g.bfs(start, func(from, to int, weight int) {
			dist[to] = dist[from] + weight
		})
		for v := range dist {
			if dist[v] != inf && dist[v] > maxPath {
				maxPath = dist[v]
				anotherStart = v
			}
		}
		return
	}
	s0 := 3
	s1, _ := calc(s0)
	s2, ans := calc(s1)
	t.Log(s0, s1, s2, ans)
	assert.Equal(t, ans, 4)
}

func Test_graph_shortestPaths(t *testing.T) {
	defer t.Skip()
	n := 6
	g := newGraph(n, 0)
	g.addBoth(1, 2, 1)
	g.addBoth(1, 3, 1)
	g.addBoth(3, 4, 100)
	g.addBoth(4, 5, 1)
	g.addBoth(5, 6, 1)
	g.addBoth(3, 6, 1)
	dist, parents := g.shortestPaths(1)
	t.Log(dist[1:])
	t.Log(parents[1:])
}

func Test_graph_mstKruskal(t *testing.T) {
	n := 6
	g := newGraph(n, 6)
	// 只需添加一条边
	g.add(1, 2, 1)
	g.add(1, 3, 1)
	g.add(3, 4, 100)
	g.add(4, 5, 1)
	g.add(5, 6, 1)
	g.add(3, 6, 1)
	sum := g.mstKruskal()
	assert.EqualValues(t, sum, 5)

	g = newGraph(n, 6)
	// 只需添加一条边
	g.add(1, 2, 1)
	g.add(1, 3, 2)
	g.add(3, 4, 100)
	g.add(4, 5, 3)
	g.add(5, 6, 4)
	g.add(3, 6, 5)
	sum = g.mstKruskal()
	assert.EqualValues(t, sum, 15)
}

func Test_directedGraph_topSort(t *testing.T) {
	g := newDirectedGraph(6, 0)
	g.add(1, 2, 1)
	g.add(2, 3, 1)
	g.add(3, 4, 1)
	g.add(3, 5, 1)
	g.add(5, 6, 1)
	order, acyclic := g.topSort()
	t.Log(order)
	assert.True(t, acyclic)

	g = newDirectedGraph(6, 0)
	g.add(1, 2, 1)
	g.add(2, 3, 1)
	g.add(3, 4, 1)
	g.add(3, 5, 1)
	g.add(5, 6, 1)
	g.add(6, 3, 1)
	order, acyclic = g.topSort()
	t.Log(order)
	assert.False(t, acyclic)
}
