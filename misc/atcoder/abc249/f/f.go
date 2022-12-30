package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, s int
	Fscan(in, &n, &k)
	a := make([]struct{ t, y int }, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].t, &a[i].y)
	}
	ans := int(-1e18)
	h := hp{}
	for i := n; i >= 0; i-- {
		q := a[i]
		y := q.y
		if q.t != 2 {
			ans = max(ans, y+s)
			if k == 0 {
				break
			}
			k-- // 必须 skip
		} else if y >= 0 {
			s += y
		} else {
			heap.Push(&h, -y) // skip
		}
		if h.Len() > k {
			s -= heap.Pop(&h).(int) // 选一个最大的负数不 skip
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
type hp struct{ sort.IntSlice }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func max(a, b int) int { if b > a { return b }; return a }
