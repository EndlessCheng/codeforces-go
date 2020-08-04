package __99

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
)

type pr20 struct {
	d int64
	v int
}
type hp20 []pr20

func (h hp20) Len() int              { return len(h) }
func (h hp20) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp20) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp20) Push(v interface{})   { *h = append(*h, v.(pr20)) }
func (h *hp20) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

// github.com/EndlessCheng/codeforces-go
func CF20C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type edge struct {
		to int
		w  int64
	}

	var n, m, v, w int
	var weight int64
	Fscan(in, &n, &m)
	g := make([][]edge, n+1)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &weight)
		g[v] = append(g[v], edge{w, weight})
		g[w] = append(g[w], edge{v, weight})
	}

	const inf int64 = 1e18
	dist := make([]int64, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[1] = 0
	fa := make([]int, n+1)
	h := hp20{{0, 1}}
	for len(h) > 0 {
		p := Pop(&h).(pr20)
		d, v := p.d, p.v
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := d + e.w; newD < dist[w] {
				dist[w] = newD
				fa[w] = v
				Push(&h, pr20{newD, w})
			}
		}
	}
	if dist[n] == inf {
		Fprint(out, -1)
		return
	}
	path := []int{}
	for x := n; x > 0; x = fa[x] {
		path = append(path, x)
	}
	for i := len(path) - 1; i >= 0; i-- {
		Fprint(out, path[i], " ")
	}
}

//func main() { CF20C(os.Stdin, os.Stdout) }
