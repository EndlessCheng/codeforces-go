package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1477C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ x, y int64 }

	var n int
	Fscan(in, &n)
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	dis := func(i, j int) int64 { x, y := a[i].x-a[j].x, a[i].y-a[j].y; return x*x + y*y }

	used := make([]bool, n)
	used[0] = true
	Fprint(out, 1)
	for i, pre := 1, 0; i < n; i++ {
		k := -1
		for j, u := range used {
			if !u && (k < 0 || dis(pre, j) > dis(pre, k)) {
				k = j
			}
		}
		Fprint(out, " ", k+1)
		used[k] = true
		pre = k
	}
}

//func main() { CF1477C(os.Stdin, os.Stdout) }
