package copypasta

import "testing"

func Test_flowGraph_maxFlow(t *testing.T) {
	g := newFlowGraph(5)
	g.addEdge(0, 1, 10)
	g.addEdge(0, 2, 2)
	g.addEdge(1, 2, 6)
	g.addEdge(1, 3, 6)
	g.addEdge(2, 4, 5)
	g.addEdge(3, 2, 3)
	g.addEdge(3, 4, 8)
	t.Log(g.maxFlow(0, 4))
}
