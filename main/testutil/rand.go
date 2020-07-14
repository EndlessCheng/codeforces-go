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

func (r *RG) String() string {
	return r.sb.String()
}

func (r *RG) NewLine() {
	r.sb.WriteByte('\n')
}

func (r *RG) Int(min, max int) int {
	v := min + rand.Intn(max-min+1)
	r.sb.WriteString(strconv.Itoa(v))
	r.sb.WriteByte(' ')
	return v
}

func (r *RG) Float(min, max float64) float64 {
	const precision = 6
	v := min + rand.Float64()*(max-min)
	r.sb.WriteString(strconv.FormatFloat(v, 'f', precision, 64))
	r.sb.WriteByte(' ')
	return v
}

func (r *RG) Slice(size, min, max int) []int {
	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, r.Int(min, max))
	}
	r.NewLine()
	return a
}

// 长度为 max-min+1
func (r *RG) Permutation(min, max int) []int {
	size := max - min + 1
	p := make(sort.IntSlice, 0, size)
	for i := min; i <= max; i++ {
		p = append(p, i)
	}
	rand.Shuffle(size, p.Swap)
	for _, v := range p {
		r.sb.WriteString(strconv.Itoa(v) + " ")
	}
	r.NewLine()
	return p
}

// todo: weighted
func (r *RG) Graph(n, m int, directed bool) (edges [][2]int) {
	// TODO
	// 无自环重边
	return
}

func (r *RG) DAG(n, m int) (edges [][2]int) {
	// TODO
	return
}

func (r *RG) Tree(n int, directed bool) (edges [][2]int) {
	// TODO
	return
}
