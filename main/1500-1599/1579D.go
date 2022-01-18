package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair79 struct{ v, i int }
type hp79 []pair79

func (h hp79) Len() int            { return len(h) }
func (h hp79) Less(i, j int) bool  { return h[i].v > h[j].v }
func (h hp79) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp79) Push(v interface{}) { *h = append(*h, v.(pair79)) }
func (h *hp79) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func CF1579D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		h := make(hp79, n)
		for i := range h {
			Fscan(in, &h[i].v)
			h[i].i = i + 1
		}
		heap.Init(&h)
		ans := [][2]int{}
		for {
			p, q := heap.Pop(&h).(pair79), heap.Pop(&h).(pair79)
			if q.v == 0 {
				break
			}
			ans = append(ans, [2]int{p.i, q.i})
			p.v--
			q.v--
			heap.Push(&h, p)
			heap.Push(&h, q)
		}
		Fprintln(out, len(ans))
		for _, p := range ans {
			Fprintln(out, p[0], p[1])
		}
	}
}

//func main() { CF1579D(os.Stdin, os.Stdout) }
