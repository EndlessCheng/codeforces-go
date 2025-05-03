package main

import (
	. "fmt"
	"io"
	"math/big"
	"sort"
)

// https://github.com/EndlessCheng
type vec60 struct{ x, y int }
func (a vec60) sub(b vec60) vec60 { return vec60{a.x - b.x, a.y - b.y} }
func (a vec60) dot(b vec60) int   { return a.x*b.x + a.y*b.y }
func (a vec60) detCmp(b vec60) int {
	v := new(big.Int).Mul(big.NewInt(int64(a.x)), big.NewInt(int64(b.y)))
	w := new(big.Int).Mul(big.NewInt(int64(a.y)), big.NewInt(int64(b.x)))
	return v.Cmp(w)
}

func cf660F(in io.Reader, out io.Writer) {
	var n, v, s, s2, ans int
	Fscan(in, &n)
	q := []vec60{{}}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		s += v
		s2 += v * i

		p := vec60{-s, 1}
		j := sort.Search(len(q)-1, func(j int) bool { return p.dot(q[j]) > p.dot(q[j+1]) })
		ans = max(ans, p.dot(q[j])+s2)

		p = vec60{i, s*i - s2}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).detCmp(p.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	Fprint(out, ans)
}

//func main() { cf660F(bufio.NewReader(os.Stdin), os.Stdout) }
