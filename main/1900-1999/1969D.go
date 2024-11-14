package main

import (
	"container/heap"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1969D(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		type pair struct{ x, y int }
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
		}
		slices.SortFunc(a, func(a, b pair) int { return b.y - a.y })

		s := 0
		h := hp69{make([]int, k)}
		for i, p := range a[:k] {
			h.IntSlice[i] = p.x
			s -= p.x
		}
		heap.Init(&h)

		for _, p := range a[k:] {
			s += max(p.y-p.x, 0)
		}
		ans := max(s, 0)
		for _, p := range a[k:] {
			s -= max(p.y-p.x, 0)
			if k > 0 && p.x < h.IntSlice[0] {
				s += h.IntSlice[0] - p.x
				ans = max(ans, s)
				h.IntSlice[0] = p.x
				heap.Fix(&h, 0)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1969D(bufio.NewReader(os.Stdin), os.Stdout) }
type hp69 struct{ sort.IntSlice }
func (h hp69) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp69) Push(any)             {}
func (hp69) Pop() (_ any)         { return }
