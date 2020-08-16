package main

import (
	"bufio"
	. "fmt"
	"io"
)

type vec1299 struct{ x, y int64 }

func (a vec1299) sub(b vec1299) vec1299 { return vec1299{a.x - b.x, a.y - b.y} }
func (a vec1299) det(b vec1299) int64   { return a.x*b.y - a.y*b.x }

// github.com/EndlessCheng/codeforces-go
func CF1299C(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := read()
	s := make([]int64, n+1)
	ps := make([]vec1299, 1, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + int64(read())
		p := vec1299{int64(i), s[i]}
		for {
			sz := len(ps)
			if sz <= 1 || ps[sz-1].sub(ps[sz-2]).det(p.sub(ps[sz-1])) > 0 {
				break
			}
			ps = ps[:sz-1]
		}
		ps = append(ps, p)
	}
	sz := len(ps)
	for i := 1; i < sz; i++ {
		a, b := ps[i-1], ps[i]
		l := b.x - a.x
		avg := float64(b.y-a.y) / float64(l)
		for ; l > 0; l-- {
			Fprintf(out, "%.9f\n", avg)
		}
	}
}

//func main() { CF1299C(os.Stdin, os.Stdout) }
