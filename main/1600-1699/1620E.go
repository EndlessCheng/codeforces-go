package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1620E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var q, op int
	Fscan(in, &q)
	qs := make([]struct{ x, y int }, q)
	for i := range qs {
		Fscan(in, &op, &qs[i].x)
		if op == 2 {
			Fscan(in, &qs[i].y)
		}
	}

	pa := [5e5 + 1]int{}
	ans := make([]int, 0, 5e5)
	for i := q - 1; i >= 0; i-- {
		x, y := qs[i].x, qs[i].y
		if y == 0 {
			if pa[x] > 0 {
				x = pa[x]
			}
			ans = append(ans, x)
		} else {
			if pa[y] > 0 {
				y = pa[y]
			}
			pa[x] = y
		}
	}
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, ans[i], " ")
	}
}

//func main() { CF1620E(os.Stdin, os.Stdout) }
