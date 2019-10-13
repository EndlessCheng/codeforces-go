package copypasta

import (
	. "container/heap"
	"sort"
)

type color int8

const (
	colorNone color = iota
	colorBlack
	colorWhite
)

type neighbor struct {
	vertex int
	weight int
}

type graph struct {
	size     int
	edgeSize int
	edges    [][]neighbor
	visited  []bool
	color    []color
}

func newGraph(size, edgeSize int) *graph {
	return &graph{
		size:     size,
		edgeSize: edgeSize,
		edges:    make([][]neighbor, size+1),
		visited:  make([]bool, size+1),
		color:    make([]color, size+1),
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

func (g *graph) resetStates() {
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
	for queue := []int{v}; len(queue) > 0; {
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

// Floyd's Algorithm
func (g *graph) allShortestPaths() [][]int {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	const inf int = 1e9

	var n, m int
	// read n m

	// 该图的邻接矩阵
	weights := make([][]int, n+1)
	for i := range weights {
		weights[i] = make([]int, n+1)
		for j := range weights[i] {
			weights[i][j] = inf
		}
	}
	for i := range weights {
		weights[i][i] = 0
	}
	for i := 0; i < m; i++ {
		var v, w, weight int
		// read v w weight
		weights[v][w] = weight
		weights[w][v] = weight
		// 或
		// weights[v][w] = min(weights[v][w], weight)
		// weights[w][v] = min(weights[w][v], weight)
	}

	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, n+1)
		copy(dist[i], weights[i])
	}
	for k := 1; k <= n; k++ { // 阶段
		for i := 1; i <= n; i++ { // 状态
			for j := 1; j <= n; j++ { // 决策
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return dist
}

// Floyd's Algorithm
func (g *graph) shortestCycle() int {
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	const inf int = 1e8 // *NOTE*

	var n, m int
	// read n m

	// 该图的邻接矩阵
	weights := make([][]int, n+1)
	for i := range weights {
		weights[i] = make([]int, n+1)
		for j := range weights[i] {
			weights[i][j] = inf
		}
	}
	for i := range weights {
		weights[i][i] = 0
	}
	for i := 0; i < m; i++ {
		var v, w, weight int
		// read v w weight
		weights[v][w] = weight
		weights[w][v] = weight
	}

	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, n+1)
		copy(dist[i], weights[i])
	}
	ans := inf
	for k := 1; k <= n; k++ { // 阶段
		for i := 1; i < k; i++ { // 状态
			for j := 1; j < i; j++ { // 决策
				ans = min(ans, dist[i][j]+weights[i][k]+weights[k][j])
			}
		}
		for i := 1; i <= n; i++ { // 状态
			for j := 1; j <= n; j++ { // 决策
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return ans
}

// ShortestPaths uses the Dijkstra's Algorithm to compute the shortest paths from `start` to all other vertices.
// The number dist[w] equals the length of a shortest path from `start` to w,
// or is `inf` if w cannot be reached.
// The number parent[w] is the predecessor of w on a shortest path from `start` to w,
// or -1 if none exists.
//
// The time complexity is O((|E| + |V|)⋅log|V|), where |E| is the number of edges
// and |V| the number of vertices in the graph.
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
	Push(h, hPair{0, start})
	for h.Len() > 0 {
		p := Pop(h).(hPair)
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
				Push(h, hPair{newDist, w})
			}
		}
	}

	// path from n to start
	//path := []int{}
	//for v := n; v != -1; v = parents[v] {
	//	path = append(path, v)
	//}
	return
}

func (g *graph) mstKruskal() (sum int64) {
	fa := make([]int, g.size+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	type edge struct {
		v, w   int
		weight int
	}
	edges := make([]edge, 0, g.edgeSize)
	for v, es := range g.edges {
		for _, e := range es {
			edges = append(edges, edge{v, e.vertex, e.weight})
		}
	}
	sort.Slice(edges, func(i, j int) bool { return edges[i].weight < edges[j].weight })
	for _, e := range edges {
		if from, to := find(e.v), find(e.w); from != to {
			sum += int64(e.weight)
			fa[from] = to
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
func (g *graph) mstPrim() (parent []int) {
	const inf int = 0x3f3f3f3f
	parent = make([]int, g.size+1)
	weights := make([]int, g.size+1)
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

func (g *graph) _isBipartite(v int) bool {
	for _, e := range g.edges[v] {
		w := e.vertex
		if g.color[w] == g.color[v] {
			return false
		}
		if g.color[w] == colorNone {
			g.color[w] = 3 - g.color[v]
			if !g._isBipartite(w) {
				return false
			}
		}
	}
	return true
}

func (g *graph) isBipartite(v int) bool {
	g.color[v] = colorBlack
	return g._isBipartite(v)
}

//

type directedGraph struct {
	*graph
	outDegree []int
	inDegree  []int
}

func newDirectedGraph(size, edgeSize int) *directedGraph {
	return &directedGraph{
		graph:     newGraph(size, edgeSize),
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
	vOrder := make([]int, g.size+1)
	for i := 1; i <= g.size; i++ {
		if g.inDegree[i] == 0 {
			queue = append(queue, i)
			vOrder[i] = 1
		}
	}
	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		order = append(order, v)
		for _, e := range g.edges[v] {
			w := e.vertex
			g.inDegree[w]-- // NOTE: copy g.inDegree if reusing is needed.
			if g.inDegree[w] == 0 {
				queue = append(queue, w)
				vOrder[w] = vOrder[v] + 1
			}
		}
	}
	return order, len(order) == g.size
}
