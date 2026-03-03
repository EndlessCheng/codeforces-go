package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2122C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type pair struct{ x, y, i int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
			a[i].i = i + 1
		}

		slices.SortFunc(a, func(a, b pair) int { return a.x - b.x })
		slices.SortFunc(a[:n/2], func(a, b pair) int { return a.y - b.y })
		slices.SortFunc(a[n/2:], func(a, b pair) int { return b.y - a.y })

		for i := range n / 2 {
			Fprintln(out, a[i].i, a[n/2+i].i)
		}
	}
}

//func main() { cf2122C(bufio.NewReader(os.Stdin), os.Stdout) }
