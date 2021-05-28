package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp39 struct{ sort.IntSlice }

func (h *hp39) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp39) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp39) push(v int)         { heap.Push(h, v) }
func (h *hp39) pop() int           { return heap.Pop(h).(int) }

func CF1239C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var t int64
	Fscan(in, &n, &t)
	a := make([]struct{ t, i int }, n)
	for i := range a {
		Fscan(in, &a[i].t)
		a[i].i = i
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.t < b.t || a.t == b.t && a.i < b.i })

	ans := make([]interface{}, n)
	q, w := []int{}, hp39{}
	for left, i, cur := n, 0, int64(0); left > 0; left-- {
		if len(q) == 0 {
			if len(w.IntSlice) > 0 {
				q = append(q, w.pop()) // 立刻进队
			} else {
				cur = int64(a[i].t)
			}
		}
		cur += t
		for ; i < n && int64(a[i].t) <= cur; i++ {
			if i := a[i].i; len(q) == 0 || i < q[len(q)-1] {
				q = append(q, i)
			} else {
				w.push(i)
			}
		}
		ans[q[0]] = cur
		q = q[1:]
	}
	Fprintln(out, ans...)
}

//func main() { CF1239C(os.Stdin, os.Stdout) }
