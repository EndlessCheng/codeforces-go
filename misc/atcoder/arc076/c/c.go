package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

type event struct{ p, i, d int }
type events []event

func (p events) Len() int           { return len(p) }
func (p events) Less(i, j int) bool { return p[i].p < p[j].p }
func (p events) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var r, c, n, x1, y1, x2, y2 int
	border := func(x, y int) bool { return x == 0 || x == r || y == 0 || y == c }
	getP := func(x, y int) int {
		if y == 0 || x == r {
			return x + y
		}
		return 2*r + 2*c - x - y
	}
	es := events{}
	Fscan(in, &r, &c, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &x1, &y1, &x2, &y2)
		if border(x1, y1) && border(x2, y2) {
			p1, p2 := getP(x1, y1), getP(x2, y2)
			if p1 > p2 {
				p1, p2 = p2, p1
			}
			es = append(es, event{p1, i, 1}, event{p2, i, -1})
		}
	}
	sort.Sort(es)
	s := []int{}
	for _, e := range es {
		if e.d == 1 {
			s = append(s, e.i)
		} else {
			if s[len(s)-1] != e.i {
				Fprint(_w, "NO")
				return
			}
			s = s[:len(s)-1]
		}
	}
	Fprint(_w, "YES")
}

func main() { run(os.Stdin, os.Stdout) }
