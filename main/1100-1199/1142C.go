package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type vec42 struct{ x, y int64 }

func (a vec42) sub(b vec42) vec42 { return vec42{a.x - b.x, a.y - b.y} }
func (a vec42) det(b vec42) int64 { return a.x*b.y - a.y*b.x }

func CF1142C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	ps := make([]vec42, n)
	for i := range ps {
		Fscan(in, &ps[i].x, &ps[i].y)
		ps[i].y -= ps[i].x * ps[i].x
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
	q := []vec42{}
	for i, p := range ps {
		if i+1 < n && p.x == ps[i+1].x {
			continue
		}
		for len(q) > 1 && q[len(q)-1].sub(q[len(q)-2]).det(p.sub(q[len(q)-1])) >= 0 {
			q = q[:len(q)-1]
		}
		q = append(q, p)
	}
	Fprint(out, len(q)-1)
}

//func main() { CF1142C(os.Stdin, os.Stdout) }
