package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

type pair1353 struct{ l, i int }
type hp1353 []pair1353

func (h hp1353) Len() int            { return len(h) }
func (h hp1353) Less(i, j int) bool  { a, b := h[i], h[j]; return a.l > b.l || a.l == b.l && a.i < b.i }
func (h hp1353) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp1353) Push(v interface{}) { *h = append(*h, v.(pair1353)) }
func (h *hp1353) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp1353) push(v pair1353)    { heap.Push(h, v) }
func (h *hp1353) pop() pair1353      { return heap.Pop(h).(pair1353) }

// github.com/EndlessCheng/codeforces-go
func CF1353D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		ans := make([]interface{}, n)
		h := hp1353{{n - 1, 0}}
		for v := 1; v <= n; v++ {
			p := h.pop()
			l, i := p.l>>1, p.i
			ans[i+l] = v
			if l > 0 {
				h.push(pair1353{l - 1, i})
			}
			if p.l > l {
				h.push(pair1353{p.l - l - 1, i + l + 1})
			}
		}
		Fprintln(out, ans...)
	}
}

//func main() { CF1353D(os.Stdin, os.Stdout) }
