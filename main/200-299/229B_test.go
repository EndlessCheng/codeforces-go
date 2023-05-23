package main

import (
	"container/heap"
	"fmt"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"sort"
	"testing"
)

// https://codeforces.com/problemset/problem/229/B
// https://codeforces.com/problemset/status/229/problem/B
func TestCF229B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 6
1 2 2
1 3 3
1 4 8
2 3 4
2 4 5
3 4 3
0
1 3
2 3 4
0
outputCopy
7
inputCopy
3 1
1 2 3
0
1 3
0
outputCopy
-1
inputCopy
2 1
1 2 3
1 0
0
outputCopy
4
inputCopy
3 2
1 2 2
2 3 1
1 3
1 2
0
outputCopy
4
inputCopy
2 1
1 2 3
0
1 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF229B)
}

func TestCompare(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 3)
		m := rg.Int(0, n*(n-1)/2)
		rg.NewLine()
		rg.GraphWeightedEdges(n, m, 1, 1, 3, false)
		for i := 0; i < n; i++ {
			k := rg.Int(0, 1)
			rg.IntSliceOrdered(k, 0, 3, true, true)
		}
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, sol229b, CF229B)
}

type Edge struct {
	From int
	To   int
	Dist int
	Cap  int
	Flow int
}

type Graph struct {
	NNode int
	NEdge int
	Edges []Edge
	Nodes [][]int
}

func NewGraph(nNode, nEdge int) *Graph {
	g := &Graph{
		NNode: nNode,
		NEdge: nEdge,
		Edges: make([]Edge, 0, nEdge),
		Nodes: make([][]int, nNode, nNode),
	}
	return g
}

func (g *Graph) AddEdge(from, to, dist, cap, flow int) {
	g.Nodes[from] = append(g.Nodes[from], len(g.Edges))
	g.Edges = append(g.Edges, Edge{from, to, dist, cap, flow})
	g.Nodes[to] = append(g.Nodes[to], len(g.Edges))
	g.Edges = append(g.Edges, Edge{to, from, dist, cap, flow})
}

func (g *Graph) Show() {
	fmt.Printf("n node = %d, n edge = %d\n", g.NNode, g.NEdge)
	for i, cur := range g.Nodes {
		fmt.Printf("[%d]", i)
		for _, next := range cur {
			fmt.Printf(" %d", g.Edges[next].To)
		}
		fmt.Println()
	}
}

type QueueStack struct {
	data []interface{}
}

func (q *QueueStack) size() int {
	return len(q.data)
}

func (q *QueueStack) push(v interface{}) {
	q.data = append(q.data, v)
}

func (q *QueueStack) popFront() interface{} {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func (q *QueueStack) pop() interface{} {
	last := len(q.data) - 1
	v := q.data[last]
	q.data = q.data[:last]
	return v
}

var arrival [][]int

func findNextTiming(at, earlist int) int {
	a := arrival[at]
	i := sort.Search(len(a), func(i int) bool {
		return a[i] >= earlist
	})
	for i < len(a) && a[i] == earlist {
		i++
		earlist++
	}
	return earlist
}

func Spfa(g *Graph, from, to int) int {
	n := g.NNode
	dist := make([]int, n, n)
	inQueue := make([]bool, n, n)
	dist[0] = 0
	for i := 1; i < n; i++ {
		dist[i] = -1
	}

	q := &QueueStack{}
	q.push(from)
	inQueue[from] = true
	for q.size() > 0 {
		cur := q.pop().(int)
		inQueue[cur] = false

		nextStart := findNextTiming(cur, dist[cur])
		for _, ie := range g.Nodes[cur] {
			e := g.Edges[ie]
			fromCur := nextStart + e.Dist
			if dist[e.To] == -1 || fromCur < dist[e.To] {
				dist[e.To] = fromCur
				if !inQueue[e.To] {
					q.push(e.To)
					inQueue[e.To] = true
				}
			}
		}
	}

	return dist[to]
}

type PQType struct {
	node int
	dist int
}

type PriorityQueue []PQType

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(PQType))
}

func (pq *PriorityQueue) Pop() interface{} {
	last := len(*pq) - 1
	old := *pq
	x := old[last]
	*pq = old[:last]
	return x
}

func (pq PriorityQueue) Show() {
	for _, x := range pq {
		fmt.Printf("%v <<< ", x)
	}
	fmt.Println()
}

func Dijkstra(g *Graph, from, to int) int {
	n := g.NNode
	done := make([]bool, n, n)

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, PQType{from, 0})

	ans := -1
	for pq.Len() > 0 {
		x := heap.Pop(&pq).(PQType)
		if done[x.node] {
			continue
		}
		done[x.node] = true
		if x.node == to {
			ans = x.dist
			break
		}
		nextStart := findNextTiming(x.node, x.dist)
		for _, ie := range g.Nodes[x.node] {
			e := g.Edges[ie]
			if !done[e.To] {
				x = PQType{e.To, e.Dist + nextStart}
				heap.Push(&pq, x)
			}
		}
	}

	return ans
}

func sol229b(in io.Reader, out io.Writer) {
	readInt := func() (v int) {
		Fscan(in, &v)
		return
	}

	n, m := readInt(), readInt()
	g := NewGraph(n, m)

	for i := 0; i < m; i++ {
		x, y, d := readInt(), readInt(), readInt()
		g.AddEdge(x-1, y-1, d, 0, 0)
	}

	arrival = make([][]int, n, n)
	for i := 0; i < n; i++ {
		nArrival := readInt()
		arrival[i] = make([]int, nArrival, nArrival)
		for j := 0; j < nArrival; j++ {
			arrival[i][j] = readInt()
		}
	}

	//ans := Spfa(g, 0, n - 1)
	ans := Dijkstra(g, 0, n-1)
	Fprint(out, ans)
}
