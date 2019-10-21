package copypasta

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_flowGraph_maxFlow(t *testing.T) {
	g := newFlowGraph(4)
	g.add(0, 1, 10)
	g.add(0, 2, 2)
	g.add(1, 2, 6)
	g.add(1, 3, 6)
	g.add(2, 4, 5)
	g.add(3, 2, 3)
	g.add(3, 4, 8)
	assert.EqualValues(t, 11, g.maxFlow(0, 4))
}
