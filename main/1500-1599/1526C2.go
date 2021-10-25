package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp26 struct{ sort.IntSlice }

func (h *hp26) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp26) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp26) push(v int)         { heap.Push(h, v) }
func (h hp26) empty() bool         { return len(h.IntSlice) == 0 }
func (h hp26) top() int            { return h.IntSlice[0] }

func CF1526C2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, cntP int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] >= 0 {
			cntP++
		}
	}

	s := int64(0)
	h := hp26{}
	for _, v := range a {
		if v >= 0 {
			s += int64(v)
		} else {
			if s+int64(v) >= 0 {
				h.push(v)
				s += int64(v)
			} else if h.Len() > 0 && v > h.top() {
				s -= int64(h.top())
				s += int64(v)
				h.IntSlice[0] = v
				heap.Fix(&h, 0)
			}
		}
	}
	Fprint(out, cntP+h.Len())
}

//func main() { CF1526C2(os.Stdin, os.Stdout) }
