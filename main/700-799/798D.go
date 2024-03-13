package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf798D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, s, se int
	Fscan(in, &n)
	a := make([]struct{ x, y, i int }, n)
	for i := range a {
		Fscan(in, &a[i].x)
		a[i].i = i + 1
	}
	for i := range a {
		Fscan(in, &a[i].y)
		s += a[i].y
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	for i := 0; i < n; i += 2 {
		se += a[i].y
	}
	if n%2 == 0 {
		se += a[n-1].y
	}
	Fprintln(out, n/2+1)
	if se*2 < s {
		for i := 1; i < n; i += 2 {
			Fprint(out, a[i].i, " ")
		}
		Fprintln(out, a[n-2+n%2].i)
	} else {
		for i := 0; i < n; i += 2 {
			Fprint(out, a[i].i, " ")
		}
		if n%2 == 0 {
			Fprint(out, a[n-1].i)
		}
		Fprintln(out)
	}
}

//func main() { cf798D(os.Stdin, os.Stdout) }
