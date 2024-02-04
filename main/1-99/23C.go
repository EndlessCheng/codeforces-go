package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf23C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		n = n*2 - 1
		a := make([]struct{ x, y, i int }, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
			s += a[i].y
			a[i].i = i + 1
		}
		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

		se := 0
		for i := 0; i < n; i += 2 {
			se += a[i].y
		}
		Fprintln(out, "YES")
		if se*2 < s {
			for i := 1; i < n; i += 2 {
				Fprint(out, a[i].i, " ")
			}
			Fprintln(out, a[n-1].i)
		} else {
			for i := 0; i < n; i += 2 {
				Fprint(out, a[i].i, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf23C(os.Stdin, os.Stdout) }
