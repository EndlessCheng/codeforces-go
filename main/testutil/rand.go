package testutil

import (
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
	return &RG{&strings.Builder{}}
}

func NewRandGeneratorWithSeed(seed int64) *RG {
	rand.Seed(seed)
	return NewRandGenerator()
}

type RG struct {
	sb *strings.Builder
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

// IntSlice generates a random int slice with a fixed size and its values in range [min, max]
func (r *RG) IntSlice(size int, min, max int) []int {
	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, r.Int(min, max))
	}
	r.NewLine()
	return a
}

func (r *RG) IntSliceOrdered(size int, min, max int, inc bool) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = r._int(min, max)
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

// FloatSlice generates a random float slice with a fixed size and its values in range [min, max]
func (r *RG) FloatSlice(size int, min, max float64, precision int) []float64 {
	a := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, r.Float(min, max, precision))
	}
	r.NewLine()
	return a
}

// UniqueSlice generates a int slice with a fixed size and all ints are unique within range [min, max]
func (r *RG) UniqueSlice(size int, min, max int) []int {
	if size > max-min+1 {
		panic("size is too large")
	}
	p := rand.Perm(max - min + 1)[:size]
	for i := range p {
		p[i] += min
	}
	for _, v := range p {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return p
}

// Permutation generates a random permutation with a fixed size and its values in range [min, max]
func (r *RG) Permutation(min, max int) []int {
	size := max - min + 1
	return r.UniqueSlice(size, min, max)
}

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

func (r *RG) graphEdges(n, m, st int, directed bool) (edges [][2]int) {
	if m < n-1 {
		panic("m is too small")
	}
	if m > n*(n-1)/2 { // 64-bit, no worry about overflow
		panic("m is too large")
	}

	edges = r.treeEdges(n, st)

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
			if !has[v][w] {
				has[v][w] = true
				v += st
				w += st
				edges = append(edges, [2]int{v, w})
				break
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
	edges = make([][3]int, n-1)
	for i, e := range r.graphEdges(n, m, st, directed) {
		weight := r._int(minWeight, maxWeight)
		r.sb.WriteString(fmt.Sprintln(e[0], e[1], weight))
		edges[i] = [3]int{e[0], e[1], weight}
	}
	return
}
