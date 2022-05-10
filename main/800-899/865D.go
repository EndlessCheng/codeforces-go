package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp65 struct{ sort.IntSlice }
func (h *hp65) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp65) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func CF865D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	ans := int64(0)
	h := hp65{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		if h.Len() > 0 && h.IntSlice[0] < v {
			ans += int64(v - h.IntSlice[0])
			h.IntSlice[0] = v
			heap.Fix(&h, 0)
		}
		heap.Push(&h, v)
	}
	Fprint(out, ans)
}

//func main() { CF865D(os.Stdin, os.Stdout) }
