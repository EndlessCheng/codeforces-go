package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf297C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	s := make([]int, n)
	p := make([]int, n)
	for i := range s {
		Fscan(in, &s[i])
		p[i] = i
	}
	slices.SortFunc(p, func(i, j int) int { return s[i] - s[j] })

	m := (n + 2) / 3
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < m; i++ {
		a[p[i]] = i
		b[p[i]] = s[p[i]] - a[p[i]]
	}
	for i := m * 2; i < n; i++ {
		b[p[i]] = n - 1 - i
		a[p[i]] = s[p[i]] - b[p[i]]
	}
	for i := m; i < m*2 && i < n; i++ {
		b[p[i]] = s[p[i]]
	}

	Fprintln(out, "YES")
	for _, v := range a {
		Fprint(out, v, " ")
	}
	Fprintln(out)
	for _, v := range b {
		Fprint(out, v, " ")
	}
}

//func main() { cf297C(bufio.NewReader(os.Stdin), os.Stdout) }
