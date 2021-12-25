package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp51 struct{ sort.IntSlice }

func (h *hp51) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp51) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func CF1251E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ m, p int }, n)
		for i := range a {
			Fscan(in, &a[i].m, &a[i].p)
		}
		// 按 (mi,pi) 排序，然后把 (i,mi) 画在平面直角坐标系上
		// 初始时，在 y=x 直线下方的点都可以视作是「免费」的，如果有不能免费的点，应考虑从最后一个不能免费的到末尾这段中的最小 pi，然后将 y=x 抬高成 y=x+1 继续比较
		// 维护最小 pi 可以用最小堆
		sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.m < b.m || a.m == b.m && a.p < b.p })
		ans := int64(0)
		for i, buy, h := n-1, 0, new(hp51); i >= 0; i-- {
			heap.Push(h, a[i].p)
			if a[i].m > i+buy {
				buy++
				ans += int64(heap.Pop(h).(int))
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1251E2(os.Stdin, os.Stdout) }
