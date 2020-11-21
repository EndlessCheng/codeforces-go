package testutil

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestRG_Permutation(t *testing.T) {
	assert := assert.New(t)
	rg := NewRandGenerator()
	min, max := 3, 7
	n := max - min + 1
	p := rg.Permutation(n, min, max)
	assert.Len(p, n)
	sort.Ints(p)
	for i, v := range p {
		assert.Equal(min+i, v)
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
