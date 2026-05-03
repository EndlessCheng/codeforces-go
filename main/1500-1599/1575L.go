package main

import (
	"cmp"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1575L(in io.Reader, out io.Writer) {
	var n, x int
	Fscan(in, &n)
	a := [][2]int{}
	for i := 1; i <= n; i++ {
		Fscan(in, &x)
		if i >= x {
			a = append(a, [2]int{x, i - x})
		}
	}

	slices.SortFunc(a, func(a, b [2]int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	g := []int{}
	for _, e := range a {
		h := e[1]
		j := sort.SearchInts(g, h+1)
		if j < len(g) {
			g[j] = h
		} else {
			g = append(g, h)
		}
	}
	Fprint(out, len(g))
}

//func main() { cf1575L(bufio.NewReader(os.Stdin), os.Stdout) }
