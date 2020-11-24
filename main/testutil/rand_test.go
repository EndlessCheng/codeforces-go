package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestRG_Permutation(t *testing.T) {
	rg := NewRandGenerator()
	min, max := 3, 7
	n := max - min + 1
	p := rg.Permutation(n, min, max)
	assert.Len(t, p, n)
	sort.Ints(p)
	for i, v := range p {
		assert.Equal(t, min+i, v)
	}
}

func TestRG_TreeEdges(t *testing.T) {
	rg := NewRandGenerator()
	n, st := 10, 1
	es := rg.TreeEdges(n, st)
	g := make([][]int, n)
	for _, e := range es {
		v, w := e[0]-st, e[1]-st
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	cc := 0
	var f func(v, fa int)
	f = func(v, fa int) {
		cc++
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
	}
	f(0, -1)
	assert.Equal(t, n, cc)
}

func TestRG_TreeWeightedEdges(t *testing.T) {
	rg := NewRandGenerator()
	n, st := 10, 1
	mi, mx := 0, 1
	es := rg.TreeWeightedEdges(n, st, mi, mx)
	g := make([][]int, n)
	for _, e := range es {
		assert.True(t, mi <= e[2] && e[2] <= mx)
		v, w := e[0]-st, e[1]-st
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	cc := 0
	var f func(v, fa int)
	f = func(v, fa int) {
		cc++
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
	}
	f(0, -1)
	assert.Equal(t, n, cc)
}

func TestRG_GraphEdges(t *testing.T) {
	rg := NewRandGenerator()
	n := 10
	m := n * (n - 1) / 2
	st := 1
	edges := rg.GraphEdges(n, m, st, false)
	for _, e := range edges {
		fmt.Println(e[0], e[1])
	}

	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}
	for _, e := range edges {
		v, w := e[0]-st, e[1]-st
		g[v][w] = true
		g[w][v] = true
	}
	for i, row := range g {
		assert.False(t, row[i])
		for j, has := range row {
			if j != i {
				assert.True(t, has)
			}
		}
	}
}
