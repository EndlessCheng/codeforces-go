package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
)

type hPair706C struct {
	x int64
	y int
}
type pairHeap706C []hPair706C

func (h pairHeap706C) Len() int              { return len(h) }
func (h pairHeap706C) Less(i, j int) bool {
	return h[i].x < h[j].x || h[i].x == h[j].x && h[i].y < h[j].y
}
func (h pairHeap706C) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap706C) Push(v interface{})   { *h = append(*h, v.(hPair706C)) }
func (h *pairHeap706C) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

// github.com/EndlessCheng/codeforces-go
func Sol706C(reader io.Reader, writer io.Writer) {
	const inf int64 = 1e18
	reverse := func(ss string) string {
		n := len(ss)
		s := make([]byte, n)
		for i := range s {
			s[i] = ss[n-1-i]
		}
		return string(s)
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	costs := make([]int, n)
	for i := range costs {
		Fscan(in, &costs[i])
	}
	type neighbor struct {
		vertex int
		weight int
	}
	g := make([][]neighbor, 2*n+2)
	var prev, rPrev, cur, rCur string
	for i, c := range costs {
		Fscan(in, &cur)
		rCur = reverse(cur)
		if prev == "" {
			g[0] = []neighbor{{1, 0}, {2, c}}
		} else {
			if prev <= cur {
				g[2*i-1] = append(g[2*i-1], neighbor{2*i + 1, 0})
			}
			if prev <= rCur {
				g[2*i-1] = append(g[2*i-1], neighbor{2*i + 2, c})
			}
			if rPrev <= cur {
				g[2*i] = append(g[2*i], neighbor{2*i + 1, 0})
			}
			if rPrev <= rCur {
				g[2*i] = append(g[2*i], neighbor{2*i + 2, c})
			}
		}
		prev, rPrev = cur, rCur
	}
	g[2*n-1] = []neighbor{{2*n + 1, 0}}
	g[2*n] = []neighbor{{2*n + 1, 0}}

	dist := make([]int64, 2*n+2)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0
	visited := make([]bool, 2*n+2)
	h := &pairHeap706C{}
	Push(h, hPair706C{0, 0})
	for h.Len() > 0 {
		p := Pop(h).(hPair706C)
		v := p.y
		if visited[v] {
			continue
		}
		visited[v] = true
		for _, e := range g[v] {
			w := e.vertex
			if newDist := dist[v] + int64(e.weight); newDist < dist[w] {
				dist[w] = newDist
				Push(h, hPair706C{newDist, w})
			}
		}
	}
	ans := dist[2*n+1]
	if ans == inf {
		ans = -1
	}
	Fprintln(out, ans)
}

//func main() {
//	Sol706C(os.Stdin, os.Stdout)
//}
