package copypasta

type neighbor struct {
	vertex int
	weight int
}

type graph struct {
	edges   [][]neighbor
	visited []bool
}

func newGraph(size int) *graph {
	return &graph{
		edges:   make([][]neighbor, size+1),
		visited: make([]bool, size+1),
	}
}

func (g *graph) addBoth(from, to int, weight int) {
	g.edges[from] = append(g.edges[from], neighbor{to, weight})
	if from != to {
		g.edges[to] = append(g.edges[to], neighbor{from, weight})
	}
}

func (g *graph) reset() {
	g.visited = make([]bool, len(g.edges))
}

func (g *graph) dfs(v int, do func(from, to int, weight int)) {
	g.visited[v] = true
	for _, e := range g.edges[v] {
		w, weight := e.vertex, e.weight
		if !g.visited[w] {
			do(v, w, weight)
			g.dfs(w, do)
		}
	}
}

// ShortestPaths computes the shortest paths from v to all other vertices.
// Only edges with non-negative weights are included.
// The number parent[w] is the predecessor of w on a shortest path from v to w,
// or -1 if none exists.
// The number dist[w] equals the length of a shortest path from v to w,
// or is -1 if w cannot be reached.
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func (g *graph) shortestPaths(v int) (parent []int, dist []int) {
	dist = make([]int, len(g.edges[v]))
	parent = make([]int, len(g.edges[v]))
	for i := range dist {
		dist[i], parent[i] = -1, -1
	}
	dist[v] = 0

	// Dijkstra's algorithm
	queue := emptyPrioQueue(dist)
	queue.Push(v)
	for queue.Len() > 0 {
		v := queue.Pop()
		for _, e := range g.edges[v] {
			w, d := e.vertex, e.weight
			if d < 0 {
				continue
			}
			alt := dist[v] + d
			switch {
			case dist[w] == -1:
				dist[w], parent[w] = alt, v
				queue.Push(w)
			case alt < dist[w]:
				dist[w], parent[w] = alt, v
				queue.Fix(w)
			}
		}
	}
	return
}
