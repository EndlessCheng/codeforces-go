package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1837F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		suf := make([]int, n+1)
		ans := sort.Search(k*1e9, func(mx int) bool {
			h := hp37{}
			s := mx
			upd := func(v int) {
				if s >= v {
					s -= v
					heap.Push(&h, v)
				} else if h.Len() > 0 && v < h.IntSlice[0] {
					s += h.IntSlice[0] - v
					h.IntSlice[0] = v
					heap.Fix(&h, 0)
				}
			}
			for i := n - 1; i >= 0; i-- {
				upd(a[i])
				if h.Len() >= k {
					return true
				}
				suf[i] = h.Len()
			}

			h = hp37{}
			s = mx
			for i, v := range a {
				upd(v)
				if h.Len()+suf[i+1] >= k {
					return true
				}
			}
			return false
		})
		Fprintln(out, ans)
	}
}

//func main() { cf1837F(os.Stdin, os.Stdout) }
type hp37 struct{ sort.IntSlice }
func (h hp37) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp37) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp37) Pop() (_ any)         { return }
