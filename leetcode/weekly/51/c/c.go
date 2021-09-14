package main

// 周赛出的是无向图，这里写的是有向图
// Update (2017-09-26):
//   We have overhauled the problem description + test cases and specified clearly the graph is an undirected graph.
//   For the directed graph follow up please see Redundant Connection II).
//   We apologize for any inconvenience caused.
// O(nlogn) 并查集
func findRedundantDirectedConnection(edges [][]int) (redundantEdge []int) {
	numNodes := len(edges)
	uf := newUnionFind(numNodes + 1)
	parent := make([]int, numNodes+1) // parent[i] 表示 i 的父节点
	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.find(from) == uf.find(to) { // from 和 to 已连接
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}

	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，其中之一与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to)
}
