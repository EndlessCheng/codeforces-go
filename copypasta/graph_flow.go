package copypasta

// TODO: 待整理

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
	//used  []bool
}

func newFlowGraph(size int) *flowGraph {
	return &flowGraph{
		size:  size,
		edges: make([][]edge, size+1),
		level: make([]int, size+1),
		iter:  make([]int, size+1),
		//used:  make([]bool, size+1),
	}
}

func (g *flowGraph) add(from, to int, cap int) {
	g.edges[from] = append(g.edges[from], edge{to, cap, len(g.edges[to])})
	g.edges[to] = append(g.edges[to], edge{from, 0, len(g.edges[from]) - 1})
}

func (g *flowGraph) addBoth(from, to int, cap int) {
	g.add(from, to, cap)
	if from != to {
		g.add(to, from, cap)
	}
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
		if e.cap > 0 && g.level[e.to] > g.level[v] {
			if d := g.dfs(e.to, t, g.min(f, e.cap)); d > 0 {
				e.cap -= d
				g.edges[e.to][e.rev].cap += d
				return d
			}
		}
		g.iter[v]++ // 当前弧优化（避免对没有用的边进行多次检查）
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

//// DFS 地寻找增广路
//func (g *flowGraph) dfs2(v, t int, f int) int {
//	if v == t {
//		return f
//	}
//	g.used[v] = true
//	for i := range g.edges[v] {
//		e := &g.edges[v][i]
//		if !g.used[e.to] && e.cap > 0 {
//			if d := g.dfs2(e.to, t, g.min(f, e.cap)); d > 0 {
//				e.cap -= d
//				g.edges[e.to][e.rev].cap += d
//				return d
//			}
//		}
//	}
//	return 0
//}
//
//// Ford–Fulkerson algorithm (FFA)
//// 耗时大约是 Dinic x2.4
//func (g *flowGraph) maxFlow2(s, t int) (flow int) {
//	const inf int = 1e8
//	for {
//		g.used = make([]bool, g.size+1)
//		if f := g.dfs2(s, t, inf); f > 0 {
//			flow += f
//		} else {
//			return
//		}
//	}
//}
