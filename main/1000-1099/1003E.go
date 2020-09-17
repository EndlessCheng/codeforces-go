package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1003E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type edge struct{ v, w int }
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, d, k int
	Fscan(in, &n, &d, &k)
	if d > n-1 || k == 1 && d > 1 {
		Fprint(out, "NO")
		return
	}
	es := []edge{}
	cur := d + 2
	var f func(v, l int)
	f = func(v, l int) {
		if l == 0 {
			return
		}
		i := 1
		if v <= d {
			i = 2
		}
		for ; i < k && len(es) < n-1; i++ {
			es = append(es, edge{v, cur})
			cur++
			f(cur-1, l-1)
		}
	}
	for i := 1; i <= d; i++ {
		es = append(es, edge{i, i + 1})
	}
	for i := 2; i <= d; i++ {
		f(i, min(i-1, d+1-i))
	}
	if len(es) < n-1 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for _, e := range es {
		Fprintln(out, e.v, e.w)
	}
}

//func main() { CF1003E(os.Stdin, os.Stdout) }
