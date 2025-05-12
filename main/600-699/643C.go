package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
type vec43 struct{ x, y float64 }

func (a vec43) sub(b vec43) vec43   { return vec43{a.x - b.x, a.y - b.y} }
func (a vec43) dot(b vec43) float64 { return a.x*b.x + a.y*b.y }
func (a vec43) det(b vec43) float64 { return a.x*b.y - a.y*b.x }

func cf643C(in io.Reader, out io.Writer) {
	var n, k, v int
	Fscan(in, &n, &k)
	s := make([]int, n+1)
	ss := make([]float64, n+1)
	t := make([]float64, n+1)
	for i := range n {
		Fscan(in, &v)
		s[i+1] = s[i] + v
		ss[i+1] = ss[i] + float64(s[i+1])/float64(v)
		t[i+1] = t[i] + 1/float64(v)
	}

	f := slices.Clone(ss)
	for i := 2; i <= k; i++ {
		q := []vec43{{float64(s[i-1]), f[i-1] - ss[i-1] + float64(s[i-1])*t[i-1]}}
		for j := i; j <= n-(k-i); j++ {
			p := vec43{-t[j], 1}
			for len(q) > 1 && p.dot(q[0]) >= p.dot(q[1]) {
				q = q[1:]
			}
			oldF := f[j]
			f[j] = p.dot(q[0]) + ss[j]
			p = vec43{float64(s[j]), oldF - ss[j] + float64(s[j])*t[j]}
			for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) <= 0 {
				q = q[:len(q)-1]
			}
			q = append(q, p)
		}
	}
	Fprintf(out, "%.10f", f[n])
}

//func main() { cf643C(bufio.NewReader(os.Stdin), os.Stdout) }
