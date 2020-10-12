package main

import (
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1427D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	ans := [][]interface{}{}
	p := make([]int, n+2)
	for !sort.IntsAreSorted(a) {
		for i, v := range a {
			p[v] = i
		}
		v := 1
		for ; p[v] < p[v+1]; v++ {
		}
		w := v + 1
		for ; p[w]+1 == p[w+1]; w++ {
		}
		i, j, k := p[v+1], p[w]+1, p[v]+1
		sz := []interface{}{}
		if i > 0 {
			sz = append(sz, i)
		}
		sz = append(sz, j-i, k-j)
		if k < n {
			sz = append(sz, n-k)
		}
		ans = append(ans, sz)
		a = append(append(append(a[k:], a[j:k]...), a[i:j]...), a[:i]...)
	}
	Fprintln(out, len(ans))
	for _, sz := range ans {
		Fprint(out, len(sz), " ")
		Fprintln(out, sz...)
	}
}

//func main() { CF1427D(os.Stdin, os.Stdout) }
