package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1701D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a01 = make([]struct{ l, r, i int }, n)
		for i := range a01 {
			Fscan(in, &v)
			a01[i].l = (i+1)/(v+1) + 1
			if v > 0 {
				a01[i].r = (i + 1) / v
			} else {
				a01[i].r = n
			}
			a01[i].i = i
		}
		sort.Slice(a01, func(i, j int) bool { return a01[i].l < a01[j].l })
		ans := make([]int, n)
		h := hp01{}
		for i, j := 1, 0; i <= n; i++ {
			for ; j < n && a01[j].l == i; j++ {
				heap.Push(&h, j)
			}
			k := heap.Pop(&h).(int)
			ans[a01[k].i] = i
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1701D(os.Stdin, os.Stdout) }

var a01 []struct{ l, r, i int }
type hp01 struct{ sort.IntSlice }
func (h hp01) Less(i, j int) bool  { return a01[h.IntSlice[i]].r < a01[h.IntSlice[j]].r }
func (h *hp01) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp01) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
