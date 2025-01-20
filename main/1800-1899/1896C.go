package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1896C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, x int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		type pair struct{ v, i int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].v)
			a[i].i = i
		}
		slices.SortFunc(a, func(a, b pair) int { return a.v - b.v })
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		slices.Sort(b)

		for i, v := range b[x:] {
			if a[i].v > v {
				Fprintln(out, "NO")
				continue o
			}
		}
		for i, p := range a[n-x:] {
			if p.v <= b[i] {
				Fprintln(out, "NO")
				continue o
			}
		}

		Fprintln(out, "YES")
		b = append(b[x:], b[:x]...)
		ans := make([]any, n)
		for i, v := range b {
			ans[a[i].i] = v
		}
		Fprintln(out, ans...)
	}
}

//func main() { cf1896C(bufio.NewReader(os.Stdin), os.Stdout) }
