package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

type hp13 struct{ sort.IntSlice }

func (h hp13) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp13) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp13) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

// github.com/EndlessCheng/codeforces-go
func CF13C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	h := hp13{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		heap.Push(&h, v)
		if d := h.IntSlice[0] - v; d > 0 {
			ans += d
			h.IntSlice[0] = v
			heap.Fix(&h, 0)
		}
	}
	Fprint(out, ans)
}

//func main() { CF13C(os.Stdin, os.Stdout) }
