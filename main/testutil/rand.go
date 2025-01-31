package testutil

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

const (
	Digits = "0123456789"
	Upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lower  = "abcdefghijklmnopqrstuvwxyz"
)

func NewRandGenerator() *RG {
	return NewRandGeneratorWithSeed(1)
}

func NewRandGeneratorWithSeed(seed int64) *RG {
	rand.Seed(seed)
	return &RG{&strings.Builder{}}
}

type RG struct {
	sb *strings.Builder
}

func (r *RG) Clear() {
	r.sb.Reset()
}

// for random string, see Str
func (r *RG) String() string {
	return r.sb.String()
}

func (r *RG) Space() {
	r.sb.WriteByte(' ')
}

func (r *RG) NewLine() {
	r.sb.WriteByte('\n')
}

func (r *RG) Byte(b byte) {
	r.sb.WriteByte(b)
}

func (r *RG) Bytes(s string) {
	r.sb.WriteString(s)
}

func (r *RG) One() {
	r.sb.WriteString("1\n")
}

func (r *RG) _int(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func (r *RG) IntOnly(min, max int) int {
	return r._int(min, max)
}

// Int generates a random int in range [min, max]
func (r *RG) Int(min, max int) int {
	v := r._int(min, max)
	r.sb.WriteString(strconv.Itoa(v))
	r.Space()
	return v
}

// Float generates a random float in range [min, max] with a fixed precision
func (r *RG) Float(min, max float64, precision int) float64 {
	v := min + rand.Float64()*(max-min)
	r.sb.WriteString(strconv.FormatFloat(v, 'f', precision, 64))
	r.Space()
	return v
}

// Str generates a random string with length in range [minLen, maxLen] and its chars in range [min, max]
func (r *RG) Str(minLen, maxLen int, min, max byte) string {
	l := r._int(minLen, maxLen)
	sb := &strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(byte(r._int(int(min), int(max))))
	}
	s := sb.String()
	r.sb.WriteString(s)
	r.Space()
	return s
}

// StrInSet generates a random string with length in range [minLen, maxLen] and its chars in chars
func (r *RG) StrInSet(minLen, maxLen int, chars string) string {
	l := r._int(minLen, maxLen)
	sb := &strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		sb.WriteByte(chars[rand.Intn(len(chars))])
	}
	s := sb.String()
	r.sb.WriteString(s)
	r.Space()
	return s
}

func (r *RG) intSlice(size int, min, max int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = r._int(min, max)
	}
	return a
}

func (r *RG) intSliceInSet(size int, set []int) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = set[rand.Intn(len(set))]
	}
	return a
}

// IntSlice generates a random int slice with a fixed size and its values in range [min, max]
func (r *RG) IntSlice(size int, min, max int) []int {
	a := r.intSlice(size, min, max)
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

func (r *RG) IntSliceInSet(size int, set []int) []int {
	a := r.intSliceInSet(size, set)
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

// IntSliceOrdered generates a random ordered int slice with a fixed size and its values in range [min, max]
func (r *RG) IntSliceOrdered(size int, min, max int, inc, unique bool) []int {
	var a []int
	if unique {
		a = r.uniqueSlice(size, min, max)
	} else {
		a = r.intSlice(size, min, max)
	}
	if inc {
		sort.Ints(a)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
	}
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

// IntSliceDiffNeighbor generates a random int slice with a fixed size and its values in range [min, max]
// It's guaranteed that a[i] != a[i+1] for all i in [0, size-2]
func (r *RG) IntSliceDiffNeighbor(size int, min, max int) []int {
	if min >= max {
		panic("min must less than max")
	}
	a := make([]int, size)
	a[0] = r._int(min, max)
	for i := 1; i < size; i++ {
		for {
			a[i] = r._int(min, max)
			if a[i] != a[i-1] {
				break
			}
		}
	}
	for _, v := range a {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return a
}

// IntMatrix generates a random int matrix with fixed row and col and its values in range [min, max]
func (r *RG) IntMatrix(row, col int, min, max int) [][]int {
	a := make([][]int, row)
	for i := range a {
		a[i] = r.intSlice(col, min, max)
	}
	for _, row := range a {
		for _, v := range row {
			r.sb.WriteString(strconv.Itoa(v))
			r.Space()
		}
		r.NewLine()
	}
	return a
}

func (r *RG) IntMatrixInSet(row, col int, set []int) [][]int {
	a := make([][]int, row)
	for i := range a {
		a[i] = r.intSliceInSet(col, set)
	}
	for _, row := range a {
		for _, v := range row {
			r.sb.WriteString(strconv.Itoa(v))
			r.Space()
		}
		r.NewLine()
	}
	return a
}

// FloatSlice generates a random float slice with a fixed size and its values in range [min, max]
func (r *RG) FloatSlice(size int, min, max float64, precision int) []float64 {
	a := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, r.Float(min, max, precision))
	}
	r.NewLine()
	return a
}

// TODO: O(size) 做法 https://mivik.blog.luogu.org/the-art-of-randomness
func (r *RG) uniqueSlice(size int, min, max int) []int {
	if size > max-min+1 {
		panic("size is too large")
	}
	p := rand.Perm(max - min + 1)[:size]
	for i := range p {
		p[i] += min
	}
	return p
}

// UniqueSlice generates a int slice with a fixed size and all ints are unique within range [min, max]
func (r *RG) UniqueSlice(size int, min, max int) []int {
	p := r.uniqueSlice(size, min, max)
	for _, v := range p {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return p
}

func (r *RG) IntSliceUnique(size int, min, max int) []int {
	return r.UniqueSlice(size, min, max)
}

// Permutation generates a random permutation with a fixed size and its values in range [min, max]
func (r *RG) Permutation(min, max int) []int {
	size := max - min + 1
	return r.UniqueSlice(size, min, max)
}

// UniquePoints generates n unique int points, x in range [minX, maxX], y in range [minY, maxY]
func (r *RG) UniquePoints(n, minX, maxX, minY, maxY int) (points [][2]int) {
	points = make([][2]int, n)
	has := map[[2]int]bool{}
	for i := 0; i < n; i++ {
		for {
			if p := [2]int{r._int(minX, maxX), r._int(minY, maxY)}; !has[p] {
				has[p] = true
				points[i] = p
				break
			}
		}
	}
	for _, p := range points {
		r.sb.WriteString(fmt.Sprintln(p[0], p[1]))
	}
	return
}

// 随机二叉树
// 可用于生成随机表达式等
// https://en.wikipedia.org/wiki/Random_binary_tree
// 参考：https://stackoverflow.com/questions/56873764/how-to-randomly-generate-a-binary-tree-given-the-node-number
// http://www.cs.otago.ac.nz/staffpriv/mike/Papers/RandomGeneration/RandomBinaryTrees.pdf
// https://mivik.blog.luogu.org/the-art-of-randomness
// 这里只是简单地实现，不保证均匀分布
// 若某个儿子不存在，对应的值为 -1
func (r *RG) binaryTree(n int) (children [][2]int) {
	children = make([][2]int, n)
	idx := 0
	var f func(int) int
	f = func(size int) int {
		if size == 0 {
			return -1
		}
		root := idx
		idx++
		leftSize := rand.Intn(size) // [0, size-1]
		rightSize := size - 1 - leftSize
		children[root] = [2]int{f(leftSize), f(rightSize)}
		return root
	}
	f(n)
	return
}

// 随机二叉树
// st 仅用于调整输出为 st-index
func (r *RG) BinaryTree(n, st int) (children [][2]int) {
	children = r.binaryTree(n)
	for v, ch := range children {
		for _, w := range ch {
			if w != -1 {
				r.sb.WriteString(fmt.Sprintln(v+st, w+st))
			}
		}
	}
	return
}

// 随机父节点
// 期望树高 https://blog.csdn.net/EI_Captain/article/details/109910307
// 更严格的随机树见 https://mivik.blog.luogu.org/the-art-of-randomness
// 叶子节点个数：期望 n/2 个
// On the number of leaves in a random recursive tree https://projecteuclid.org/journals/brazilian-journal-of-probability-and-statistics/volume-29/issue-4/On-the-number-of-leaves-in-a-random-recursive-tree/10.1214/14-BJPS252.pdf
func (r *RG) treeEdges(n, st int) (edges [][2]int) {
	edges = make([][2]int, 0, n-1)
	for i := 1; i < n; i++ {
		// v < w
		v := st + rand.Intn(i)
		w := st + i
		edges = append(edges, [2]int{v, w})
	}
	return
}

// TreeEdges generates a tree with n nodes, st-index, and v<w for each edge v-w.
// TODO: support set max degree limit
func (r *RG) TreeEdges(n, st int) (edges [][2]int) {
	edges = r.treeEdges(n, st)
	for _, e := range edges {
		r.sb.WriteString(fmt.Sprintln(e[0], e[1]))
	}
	return
}

// TreeWeightedEdges generates a tree with n nodes, st-index, edge weights in range [minWeight, maxWeight]
func (r *RG) TreeWeightedEdges(n, st, minWeight, maxWeight int) (edges [][3]int) {
	edges = make([][3]int, n-1)
	for i, e := range r.treeEdges(n, st) {
		weight := r._int(minWeight, maxWeight)
		r.sb.WriteString(fmt.Sprintln(e[0], e[1], weight))
		edges[i] = [3]int{e[0], e[1], weight}
	}
	return
}

// todo https://codeforces.com/blog/entry/77970
func (r *RG) graphEdges(n, m, st int, directed bool) (edges [][2]int) {
	//if m < n-1 {
	//	panic("m is too small")
	//}
	if m > n*(n-1)/2 { // 64-bit, no worry about overflow
		panic("m is too large")
	}

	edges = r.treeEdges(n, st)
	if m < n-1 {
		edges = edges[:m]
	}

	if m >= n {
		has := make([]map[int]bool, n)
		for i := range has {
			has[i] = map[int]bool{}
		}
		for _, e := range edges {
			// v < w
			v, w := e[0]-st, e[1]-st
			has[v][w] = true
		}
		for i := n - 1; i < m; i++ {
			for {
				// v < w
				v := r._int(0, n-2)
				w := r._int(v+1, n-1)
				if !has[v][w] { // todo 对于稠密图，这样做可能会导致运行时间较长，此时可以考虑生成补图，然后转化到原图
					has[v][w] = true
					v += st
					w += st
					edges = append(edges, [2]int{v, w})
					break
				}
			}
		}
	}

	if directed {
		for i := range edges {
			if rand.Intn(2) == 0 {
				edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
			}
		}
	}

	return
}

// TreeEdges generates a graph with n nodes, m edges, st-index, without self-loops and multiple edges
// TIPS: pass directed=false to generate a DAG.
func (r *RG) GraphEdges(n, m, st int, directed bool) (edges [][2]int) {
	edges = r.graphEdges(n, m, st, directed)
	for _, e := range edges {
		r.sb.WriteString(fmt.Sprintln(e[0], e[1]))
	}
	return
}

// TreeEdges generates a graph with n nodes, m edges, st-index, without self-loops and multiple edges, edge weights in range [minWeight, maxWeight]
// TIPS: pass directed=false to generate a DAG.
func (r *RG) GraphWeightedEdges(n, m, st, minWeight, maxWeight int, directed bool) (edges [][3]int) {
	edges = make([][3]int, m)
	for i, e := range r.graphEdges(n, m, st, directed) {
		weight := r._int(minWeight, maxWeight)
		r.sb.WriteString(fmt.Sprintln(e[0], e[1], weight))
		edges[i] = [3]int{e[0], e[1], weight}
	}
	return
}

func (r *RG) GraphMatrix(n int, directed bool) (g [][]byte) {
	upper := n * (n - 1) / 2
	m := r._int(0, upper)
	edges := r.graphEdges(n, m, 0, directed)
	g = make([][]byte, n)
	for i := range g {
		g[i] = bytes.Repeat([]byte{'0'}, n)
	}
	for _, e := range edges {
		g[e[0]][e[1]] = '1'
		if !directed {
			g[e[1]][e[0]] = '1'
		}
	}
	for _, row := range g {
		r.sb.WriteString(string(row))
		r.NewLine()
	}
	return
}

// Graph Hack SPFA
// GraphHackSPFA generates a undirected grid graph with n nodes, st-index, without self-loops and multiple edges, edge weights in range [minWeight, maxWeight]
// The return `edges` contains about 2n edges.
//
// For example, a 10 nodes 2 row grid graph looks like this:
// 1-2-3-4-5
// | | | | |
// 6-7-8-9-10
//
// And an 11 nodes 2 row grid graph looks like this:
// 1-2-3-4-5
// | | | | |
// 6-7-8-9-10-11
//
// All weights of vertical edges are 1, the others are random.
//
// In practice, set row to 6 will make SPFA to run in worst case (about n^2/10 relaxations).
// (Random weights are in range [1,1e5])
//      n    avg relax (100 runs)  Dijkstra's algorithm (for comparing)
//    5e3    2'418'393                            5'493
//    1e4    9'674'877                           18'323
//    2e4   38'586'596                           36'658
//    3e4   87'033'045                           54'992
//    1e5  966'319'883 (10 runs)                183'320
//
// Reference:
// https://blog.csdn.net/qq_45721135/article/details/102472101
// https://blog.csdn.net/jinzhao1994/article/details/38311561
// https://www.zhihu.com/question/292283275
// https://www.zhihu.com/question/268382638
func (r *RG) GraphHackSPFA(n, row, st, minWeight, maxWeight int, dir bool) (edges [][3]int) {
	rowLen := n / row
	m := row*(rowLen-1) + (row-1)*rowLen + n%row

	edges = make([][3]int, 0, m)
	for i := range row - 1 {
		for j := 1 + i*rowLen; j < (i+1)*rowLen; j++ {
			weight := r._int(minWeight, maxWeight)
			edges = append(edges, [3]int{j - 1, j, weight})
		}
		for j := i * rowLen; j < (i+1)*rowLen; j++ {
			edges = append(edges, [3]int{j, j + rowLen, 1})
		}
	}
	for j := 1 + (row-1)*rowLen; j < n; j++ {
		weight := r._int(minWeight, maxWeight)
		edges = append(edges, [3]int{j - 1, j, weight})
	}

	// 有向图添加双向边
	if dir {
		for _, e := range edges {
			edges = append(edges, [3]int{e[1], e[0], e[2]})
		}
	}

	rand.Shuffle(len(edges), func(i, j int) { edges[i], edges[j] = edges[j], edges[i] })

	// add st
	if st != 0 {
		for i := range edges {
			edges[i][0] += st
			edges[i][1] += st
		}
	}

	for _, e := range edges {
		r.sb.WriteString(fmt.Sprintln(e[0], e[1], e[1]))
	}
	return
}

// 菊花图
// 要求：m 和 n 数量级一样
// 要求：0 是起点
// 0 --- 1 --- end (菊花)
//   \-- 2 --/
//   \-- 3 --/
// 注：如果没有卡掉，可以试试重边（但感觉没多大意思，毕竟去个重就绕过去了）
func (r *RG) GraphHackBellmanFord(n, m, st, minWeight, maxWeight int) (edges [][3]int) {
	edges = make([][3]int, 0, m)

	end := n/3 + 1 // 假定 n=m 的写法，如果 m >= 3n/2 可以改成 n/2+1
	for i := 1; i < end; i++ {
		// 注意不能设为 1，因为 Go 和 Java 的堆只比较边权，不比较节点，这可能会导致小的 maxWeight - i*2 跑到前面去
		edges = append(edges, [3]int{0, i, i + minWeight})
		//edges = append(edges, [3]int{i, end, maxWeight - min(i, end-i)*2}) // 可以卡掉倒着遍历邻居的写法
		edges = append(edges, [3]int{i, end, maxWeight - i*2})
	}

	for i := end + 1; len(edges) < m; i++ {
		edges = append(edges, [3]int{end, i, minWeight})
	}

	// add st
	if st != 0 {
		for i := range edges {
			edges[i][0] += st
			edges[i][1] += st
		}
	}

	for _, e := range edges {
		r.sb.WriteString(fmt.Sprintln(e[0], e[1], e[1]))
	}
	return
}

//func genHackExample() {
//	sb := &strings.Builder{}
//	defer func() { writeString(sb.String()) }()
//	Print := func(a ...interface{}) { sb.WriteString(Sprint(a...)) }
//	Printf := func(format string, a ...interface{}) { sb.WriteString(Sprintf(format, a...)) }
//	Println := func(a ...interface{}) { sb.WriteString(Sprintln(a...)) }
//	_, _, _ = Print, Printf, Println
//
//	n := 25000
//	edges := rg.GraphHackSPFA(n, 6, 0, 1, 1e5)
//	Println(n)
//	Print("[")
//	for i, e := range edges {
//		if i > 0 {
//			Print(",")
//		}
//		Print("[",e[0],",",e[1],",",e[2],"]")
//	}
//	Print("]")
//}

// https://leetcode.cn/problems/find-edges-in-shortest-paths/
func simpleQueueExample(n int, edges [][]int) (ans []bool) {
	f := func(st int) []int {
		type nb struct{ v, w int }
		g := make([][]nb, n)
		for _, e := range edges {
			v, w, wt := e[0], e[1], e[2]
			g[v] = append(g[v], nb{w, wt})
			g[w] = append(g[w], nb{v, wt})
		}
		d := make([]int, n)
		for i := range d {
			d[i] = 1e18
		}
		d[st] = 0
		q := []nb{{st, 0}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.w > d[p.v] {
				continue
			}
			for _, e := range g[p.v] {
				y, w := e.v, e.w
				nw := p.w + w
				if d[y] > nw {
					d[y] = nw
					q = append(q, nb{y, nw})
				}
			}
		}
		return d
	}
	d1 := f(0)
	d2 := f(n - 1)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		ans = append(ans, d1[v]+wt+d2[w] == d1[n-1] || d1[w]+wt+d2[v] == d1[n-1])
	}
	return
}

// TODO: 随机括号序列 https://mivik.blog.luogu.org/the-art-of-randomness
