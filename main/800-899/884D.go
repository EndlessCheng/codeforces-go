package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type hp84 []int64

func (h hp84) Len() int            { return len(h) }
func (h hp84) Less(i, j int) bool  { return h[i] < h[j] }
func (h hp84) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp84) Push(v interface{}) { *h = append(*h, v.(int64)) }
func (h *hp84) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp84) pop() int64         { return heap.Pop(h).(int64) }

func CF884D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var ans int64
	Fscan(in, &n)
	h := make(hp84, n, n+1)
	for i := range h {
		Fscan(in, &h[i])
	}
	if n&1 == 0 {
		h = append(h, 0)
	}
	heap.Init(&h)
	for len(h) > 2 {
		h[0] += h.pop() + h.pop()
		ans += h[0]
		heap.Fix(&h, 0)
	}
	Fprint(out, ans)
}

//func main() { CF884D(os.Stdin, os.Stdout) }
