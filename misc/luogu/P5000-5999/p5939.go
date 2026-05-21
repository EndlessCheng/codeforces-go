package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func p5939(in io.Reader, out io.Writer) {
	var n, x, y int
	Fscan(in, &n)
	type pair struct{ x, y int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &x, &y)
		a[i] = pair{x + y, y - x}
	}
	slices.SortFunc(a, func(a, b pair) int { return cmp.Or(a.x-b.x, b.y-a.y) })

	f := []int{}
	for _, p := range a {
		v := p.y
		j := sort.SearchInts(f, v)
		if j < len(f) {
			f[j] = v
		} else {
			f = append(f, v)
		}
	}
	Fprint(out, len(f))
}

//func main() { p5939(bufio.NewReader(os.Stdin), os.Stdout) }
