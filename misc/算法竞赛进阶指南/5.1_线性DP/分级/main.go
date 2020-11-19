package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://www.luogu.com.cn/problem/P2893

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

type h2 struct{ sort.IntSlice }

func (h *h2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *h2) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans, ans2 int
	Fscan(in, &n)
	h, q := hp{}, h2{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		heap.Push(&h, v)
		if d := h.IntSlice[0] - v; d > 0 {
			ans += d
			h.IntSlice[0] = v
			heap.Fix(&h, 0)
		}
		heap.Push(&q, v)
		if d := v - q.IntSlice[0]; d > 0 {
			ans2 += d
			q.IntSlice[0] = v
			heap.Fix(&q, 0)
		}
	}
	Fprint(out, min(ans, ans2))
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
