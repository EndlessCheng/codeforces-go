package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp83 struct{ sort.IntSlice }

func (h *hp83) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp83) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp83) push(v int)         { heap.Push(h, v) }
func (h *hp83) pop() int           { return heap.Pop(h).(int) }

func CF1283F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	d := make([]int, n+1)
	a := make([]int, n-1)
	for i := range a {
		Fscan(in, &a[i])
		d[a[i]]++
	}
	Fprintln(out, a[0])
	h := &hp83{}
	for i := 1; i <= n; i++ {
		if d[i] == 0 {
			h.push(i)
		}
	}
	for i := n - 2; i >= 0; i-- {
		v := a[i]
		Fprintln(out, v, h.pop())
		if d[v]--; d[v] == 0 {
			h.push(v)
		}
	}
}

//func main() { CF1283F(os.Stdin, os.Stdout) }
