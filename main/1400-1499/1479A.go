package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

type (
	input79 struct{ n int }
	req79   struct{ i int }
	resp79  struct{ v int }
	guess79 struct{ ans int }
)

// github.com/EndlessCheng/codeforces-go
func CF1479A(in input79, Q func(req79) resp79) (gs guess79) {
	q := func(i int) int { return Q(req79{i}).v }
	n := in.n
	if n == 1 || q(1) < q(2) {
		gs.ans = 1
		return
	}
	if q(n-1) > q(n) {
		gs.ans = n
		return
	}
	res := sort.Search(n, func(i int) bool {
		if i < 2 || gs.ans > 0 {
			return false
		}
		x, y, z := q(i-1), q(i), q(i+1)
		if x > y && y < z {
			gs.ans = i
			return false
		}
		return x < y && y < z
	})
	if gs.ans == 0 {
		gs.ans = res
	}
	return
}

func ioq79() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	Q := func(req req79) (resp resp79) {
		Fprintln(out, "?", req.i)
		out.Flush()
		Fscan(in, &resp.v)
		return
	}

	d := input79{}
	Fscan(in, &d.n)
	gs := CF1479A(d, Q)
	Fprint(out, "! ", gs.ans)
	out.Flush()
}

//func main() { ioq79() }
