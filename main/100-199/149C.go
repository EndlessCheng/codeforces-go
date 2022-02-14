package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF149C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]struct{ v, i int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })
	Fprintln(out, (n+1)/2)
	for i := 0; i < n; i += 2 {
		Fprint(out, a[i].i, " ")
	}
	Fprintln(out)
	Fprintln(out, n/2)
	for i := 1; i < n; i += 2 {
		Fprint(out, a[i].i, " ")
	}
}

//func main() { CF149C(os.Stdin, os.Stdout) }
