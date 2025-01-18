package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]struct{ x, y int }, n)
		for i := range a {
			Fscan(in, &a[i].x)
		}
		for i := range a {
			Fscan(in, &a[i].y)
		}
		sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

		h := hp{make([]int, k)}
		s := 0
		for i, p := range a[:k] {
			h.IntSlice[i] = p.y
			s += p.y
		}
		heap.Init(&h)
		ans := a[k-1].x * s
		for _, p := range a[k:] {
			if p.y < h.IntSlice[0] {
				s -= h.IntSlice[0] - p.y
				ans = min(ans, p.x*s)
				h.IntSlice[0] = p.y
				heap.Fix(&h, 0)
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
