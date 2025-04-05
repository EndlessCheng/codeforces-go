package main

import (
	. "fmt"
	"io"
	"math/big"
)

// https://github.com/EndlessCheng
type vec19 struct{ x, y int }

func (a vec19) sub(b vec19) vec19 { return vec19{a.x - b.x, a.y - b.y} }
func (a vec19) dot(b vec19) int   { return a.x*b.x + a.y*b.y }
func (a vec19) det(b vec19) bool {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w) <= 0
}

func cf319C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}

	q := []vec19{{-b[0], 0}}
	for i := 1; i < n; i++ {
		p := vec19{-a[i], 1}
		for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
			q = q[1:]
		}
		f := p.dot(q[0])
		v := vec19{-b[i], f}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(v.sub(q[len(q)-1])) {
			q = q[:len(q)-1]
		}
		q = append(q, v)
	}
	Fprint(out, q[len(q)-1].y)
}

//func main() { cf319C(bufio.NewReader(os.Stdin), os.Stdout) }
