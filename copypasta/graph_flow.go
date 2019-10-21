package copypasta

type edge struct {
	to  int
	cap int
	rev int // 反向边
}

type flowGraph struct {
	size  int
	edges [][]edge
	level []int // 顶点到源点的距离标号
	iter  []int // 当前弧，在其之前的边已经没有用了
}

func newFlowGraph(size int) *flowGraph {
	return &flowGraph{
		size:  size,
		edges: make([][]edge, size+1),
		level: make([]int, size+1),
		iter:  make([]int, size+1),
	}
}

func (g *flowGraph) addEdge(from, to int, cap int) {
	g.edges[from] = append(g.edges[from], edge{to, cap, len(g.edges[to])})
	g.edges[to] = append(g.edges[to], edge{from, 0, len(g.edges[from]) - 1})
}

// BFS 地计算从源点出发的距离标号
func (g *flowGraph) bfs(v int) {
	for i := range g.level {
		g.level[i] = -1
	}
	g.level[v] = 0
	q := []int{v}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, e := range g.edges[v] {
			w, cap := e.to, e.cap
			if cap > 0 && g.level[w] < 0 {
				g.level[w] = g.level[v] + 1
				q = append(q, w)
			}
		}
	}
}

func (*flowGraph) min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// DFS 地寻找增广路
func (g *flowGraph) dfs(v, t int, f int) int {
	if v == t {
		return f
	}
	for i := g.iter[v]; i < len(g.edges[v]); i++ {
		e := &g.edges[v][i]
		if e.cap > 0 && g.level[v] < g.level[e.to] {
			if d := g.dfs(e.to, t, g.min(f, e.cap)); d > 0 {
				e.cap -= d
				g.edges[e.to][e.rev].cap += d
				return d
			}
		}
		g.iter[v]++
	}
	return 0
}

// Dinic's algorithm
func (g *flowGraph) maxFlow(s, t int) (flow int) {
	const inf int = 1e8
	for {
		g.bfs(s)
		if g.level[t] < 0 {
			return
		}
		g.iter = make([]int, g.size+1)
		for {
			if f := g.dfs(s, t, inf); f > 0 {
				flow += f
			} else {
				break
			}
		}
	}
}
