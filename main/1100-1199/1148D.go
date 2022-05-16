package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1148D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x, y int
	Fscan(in, &n)
	type pair struct{ x, i int }
	var a, b []pair
	for i := 1; i <= n; i++ {
		Fscan(in, &x, &y)
		if x < y {
			a = append(a, pair{x, i})
		} else {
			b = append(b, pair{x, i})
		}
	}
	if len(a) > len(b) {
		sort.Slice(a, func(i, j int) bool { return a[i].x > a[j].x })
	} else {
		sort.Slice(b, func(i, j int) bool { return b[i].x < b[j].x })
		a = b
	}
	Fprintln(out, len(a))
	for _, p := range a {
		Fprint(out, p.i, " ")
	}
}

//func main() { CF1148D(os.Stdin, os.Stdout) }
