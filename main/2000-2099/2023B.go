package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2023B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] += s[i-1]
		}
		g := make([]int, n)
		for i := range g {
			Fscan(in, &g[i])
			g[i]--
		}

		ans := 0
		dis := make([]int, n)
		for i := range dis {
			dis[i] = 1e18
		}
		dis[0] = 0
		h := hp23{{}}
		for len(h) > 0 {
			top := heap.Pop(&h).(pair23)
			v := top.v
			d := top.dis
			if d > dis[v] {
				continue
			}
			ans = max(ans, s[v+1]-d)
			if v > 0 && d < dis[v-1] {
				dis[v-1] = d
				heap.Push(&h, pair23{d, v - 1})
			}
			if w := g[v]; w > v {
				d += s[v+1] - s[v]
				if d < dis[w] {
					dis[w] = d
					heap.Push(&h, pair23{d, w})
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2023B(bufio.NewReader(os.Stdin), os.Stdout) }

type pair23 struct{ dis, v int }
type hp23 []pair23
func (h hp23) Len() int           { return len(h) }
func (h hp23) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp23) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp23) Push(v any)        { *h = append(*h, v.(pair23)) }
func (h *hp23) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
