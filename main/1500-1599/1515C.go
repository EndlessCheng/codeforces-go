package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1515C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, m, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &x)
		a := make([]struct{ v, i int }, n)
		for i := range a {
			Fscan(in, &a[i].v)
			a[i].i = i
		}
		sort.Slice(a, func(i, j int) bool { return a[i].v < a[j].v })
		ans := make([]interface{}, n)
		for i, p := range a {
			ans[p.i] = i%m + 1
		}
		Fprintln(out, "YES")
		Fprintln(out, ans...)
	}
}

//func main() { CF1515C(os.Stdin, os.Stdout) }
