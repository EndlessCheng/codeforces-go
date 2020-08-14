package testutil

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func NewRandGenerator() *RG {
	return &RG{&strings.Builder{}}
}

func NewRandGeneratorS(seed int64) *RG {
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

// Int generates a random int in range [min, max]
func (r *RG) Int(min, max int) int {
	v := min + rand.Intn(max-min+1)
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
	l := minLen + rand.Intn(maxLen-minLen+1)
	sb := &strings.Builder{}
	sb.Grow(l)
	for i := 0; i < l; i++ {
		c := min + byte(rand.Intn(int(max-min+1)))
		sb.WriteByte(c)
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

// FloatSlice generates a random float slice with a fixed size and its values in range [min, max]
func (r *RG) FloatSlice(size int, min, max float64, precision int) []float64 {
	a := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, r.Float(min, max, precision))
	}
	r.NewLine()
	return a
}

// Permutation generates a random permutation with a fixed size max-min+1 and its values in range [min, max]
func (r *RG) Permutation(min, max int) []int {
	size := max - min + 1
	p := make(sort.IntSlice, 0, size)
	for i := min; i <= max; i++ {
		p = append(p, i)
	}
	rand.Shuffle(size, p.Swap)
	for _, v := range p {
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
	}
	r.NewLine()
	return p
}

// TreeEdges generates a tree with n nodes, st-index
// TODO: max degree
func (r *RG) TreeEdges(n, st int) (edges [][2]int) {
	// random labels
	//labels := make(sort.IntSlice, n)
	//for i := range labels {
	//	labels[i] = i
	//}
	//rand.Shuffle(n, labels.Swap)

	edges = make([][2]int, 0, n-1)
	for i := 1; i < n; i++ {
		//v := st + labels[i]
		//w := st + labels[rand.Intn(i)]
		v := st + i
		w := st + rand.Intn(i)
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
		r.sb.WriteString(strconv.Itoa(w))
		r.NewLine()
		edges = append(edges, [2]int{v, w})
	}
	return
}

// TreeWeightedEdges generates a tree with n nodes, st-index, edge weights in range [minWeight, maxWeight]
func (r *RG) TreeWeightedEdges(n, st, minWeight, maxWeight int) (edges [][3]int) {
	// random labels
	labels := make(sort.IntSlice, n)
	for i := range labels {
		labels[i] = i
	}
	rand.Shuffle(n, labels.Swap)

	edges = make([][3]int, 0, n-1)
	for i := 1; i < n; i++ {
		v := st + labels[i]
		w := st + labels[rand.Intn(i)]
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
		r.sb.WriteString(strconv.Itoa(v))
		r.Space()
		weight := r.Int(minWeight, maxWeight)
		r.NewLine()
		edges = append(edges, [3]int{v, w, weight})
	}
	return
}

// todo: weighted
func (r *RG) GraphEdges(n, m int, directed bool) (edges [][2]int) {
	// TODO
	// 无自环重边
	return
}

func (r *RG) DAGEdges(n, m int) (edges [][2]int) {
	// TODO
	return
}
