package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pr49 struct {
	v   int
	d   int64
	isT bool
}
type hp49 []pr49

func (h hp49) Len() int              { return len(h) }
func (h hp49) Less(i, j int) bool    { return h[i].d < h[j].d || h[i].d == h[j].d && !h[i].isT } // 同等距离下优先选不是火车的路
func (h hp49) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp49) Push(v interface{})   { *h = append(*h, v.(pr49)) }
func (h *hp49) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp49) push(v pr49)          { heap.Push(h, v) }
func (h *hp49) pop() pr49            { return heap.Pop(h).(pr49) }

func CF449B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, w, wt, ans int
	Fscan(in, &n, &m, &k)
	type nb struct {
		to, wt int
		isT    bool
	}
	g := make([][]nb, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt, false})
		g[w] = append(g[w], nb{v, wt, false})
	}
	for ; k > 0; k-- {
		Fscan(in, &w, &wt)
		g[0] = append(g[0], nb{w - 1, wt, true})
	}
	vis := make([]bool, n)
	q := hp49{{}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if vis[v] {
			if p.isT {
				ans++
			}
			continue
		}
		vis[v] = true
		for _, e := range g[v] {
			q.push(pr49{e.to, p.d + int64(e.wt), e.isT})
		}
	}
	Fprint(out, ans)
}

//func main() { CF449B(os.Stdin, os.Stdout) }
