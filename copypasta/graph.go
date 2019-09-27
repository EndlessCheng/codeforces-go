package copypasta

type neighbor struct {
	vertex int
	weight int
}

type graph struct {
	size    int
	edges   [][]neighbor
	visited []bool
}

func newGraph(size int) *graph {
	return &graph{
		size:    size,
		edges:   make([][]neighbor, size+1),
		visited: make([]bool, size+1),
	}
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

func (g *graph) reset() {
	g.visited = make([]bool, g.size+1)
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

func (g *graph) bfs(v int, do func(from, to int, weight int)) {
	g.visited[v] = true
	queue := []int{v}
	for len(queue) > 0 {
		v, queue = queue[0], queue[1:]
		for _, e := range g.edges[v] {
			w, weight := e.vertex, e.weight
			if !g.visited[w] {
				do(v, w, weight)
				g.visited[w] = true
				queue = append(queue, w)
			}
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
func (g *graph) shortestPaths(v int) (parents []int, dist []int) {
	dist = make([]int, len(g.edges[v]))
	parents = make([]int, len(g.edges[v]))
	for i := range dist {
		dist[i], parents[i] = -1, -1
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
				dist[w], parents[w] = alt, v
				queue.Push(w)
			case alt < dist[w]:
				dist[w], parents[w] = alt, v
				queue.Fix(w)
			}
		}
	}
	return
}

// MST computes a minimum spanning tree for each connected component
// of an undirected weighted graph.
// The forest of spanning trees is returned as a slice of parent pointers:
// parent[v] is either the parent of v in a tree,
// or -1 if v is the root of a tree.
//
// The time complexity is O(|E|⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
func (g *graph) mst() (parent []int) {
	const inf int = 0x3f3f3f3f
	parent = make([]int, g.size)
	weights := make([]int, g.size)
	for i := range parent {
		parent[i] = -1
		weights[i] = inf
	}

	// Prim's algorithm
	queue := newPrioQueue(weights)
	for queue.Len() > 0 {
		v := queue.Pop()
		for _, e := range g.edges[v] {
			w, weight := e.vertex, e.weight
			if queue.Contains(w) && weight < weights[w] {
				weights[w] = weight
				queue.Fix(w)
				parent[w] = v
			}
		}
	}
	return
}

//

type directedGraph struct {
	*graph
	outDegree []int
	inDegree  []int
}

func newDirectedGraph(size int) *directedGraph {
	return &directedGraph{
		graph:     newGraph(size),
		outDegree: make([]int, size+1),
		inDegree:  make([]int, size+1),
	}
}

func (g *directedGraph) add(from, to int, weight int) {
	g.graph.add(from, to, weight)
	g.outDegree[from]++
	g.inDegree[to]++
}

// Kahn's algorithm
func (g *directedGraph) topSort() (order []int, acyclic bool) {
	queue := []int{}
	for i := 1; i <= g.size; i++ {
		if g.inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	var v int
	for len(queue) > 0 {
		v, queue = queue[0], queue[1:]
		order = append(order, v)
		for _, e := range g.edges[v] {
			w := e.vertex
			g.inDegree[w]-- // NOTE: copy g.inDegree if reusing is needed.
			if g.inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}
	return order, len(order) == g.size
}
