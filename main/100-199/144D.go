package _00_199

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
)

type hPair144 struct{ x, y int }
type pairHeap144 []hPair144

func (h pairHeap144) Len() int              { return len(h) }
func (h pairHeap144) Less(i, j int) bool    { return h[i].x < h[j].x || h[i].x == h[j].x && h[i].y < h[j].y }
func (h pairHeap144) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap144) Push(v interface{})   { *h = append(*h, v.(hPair144)) }
func (h *pairHeap144) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

// github.com/EndlessCheng/codeforces-go
func Sol144D(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, m, st := read(), read(), read()-1
	type neighbor struct{ to, weight int }
	g := make([][]neighbor, n)
	for i := 0; i < m; i++ {
		v, w, weight := read()-1, read()-1, read()
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}
	l := read()

	const inf int = 1e9 + 1
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	h := &pairHeap144{}
	Push(h, hPair144{0, st})
	for h.Len() > 0 {
		p := Pop(h).(hPair144)
		d, v := p.x, p.y
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.weight; newD < dist[w] {
				dist[w] = newD
				Push(h, hPair144{newD, w})
			}
		}
	}

	ans := 0
	for _, d := range dist {
		if d == l {
			ans++
		}
	}
	for v, edges := range g {
		for _, e := range edges {
			w := e.to
			if w < v {
				continue
			}
			dv, dw, size := dist[v], dist[w], e.weight
			// imagine an axis from v to w
			posv := l - dv
			posw := size - (l - dw)
			if posv > posw {
				continue
			}
			if 0 < posv && posv < size {
				ans++
			}
			if 0 < posw && posw < size && posw != posv {
				ans++
			}
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol144D(os.Stdin, os.Stdout)
//}
