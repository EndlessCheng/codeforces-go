package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go

// 用一个可修改堆来维护最小值（也可以用平衡树来维护）
type vi93 struct{ v, i int }
type hp93 []*vi93

func (h hp93) Len() int            { return len(h) }
func (h hp93) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h hp93) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].i = i; h[j].i = j }
func (h *hp93) Push(v interface{}) { *h = append(*h, v.(*vi93)) }
func (h *hp93) Pop() interface{}   { return nil }
func (h *hp93) push(v int) *vi93   { p := &vi93{v, len(*h)}; heap.Push(h, p); return p }

func CF1493DHeap(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e5
	lpf := [mx + 1]int{}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
	const mod int64 = 1e9 + 7
	pow := func(x int64, n int) int64 {
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var n, q, i, v int
	Fscan(in, &n, &q)
	hs := [mx + 1]hp93{}
	ps := make([]map[int]*vi93, n)
	for i := range ps {
		ps[i] = map[int]*vi93{}
		Fscan(in, &v)
		for v > 1 {
			p := lpf[v]
			e := 1
			for v /= p; lpf[v] == p; v /= p {
				e++
			}
			ps[i][p] = hs[p].push(e)
		}
	}
	g := int64(1)
	for p, h := range hs[:] {
		if len(h) == n {
			g = g * pow(int64(p), h[0].v) % mod
		}
	}
	for ; q > 0; q-- {
		Fscan(in, &i, &v)
		for v > 1 {
			p := lpf[v]
			e := 1
			for v /= p; lpf[v] == p; v /= p {
				e++
			}
			h := hs[p]
			before := 0
			if len(h) == n {
				before = h[0].v
			}
			if vi := ps[i-1][p]; vi == nil {
				ps[i-1][p] = h.push(e)
				hs[p] = h
			} else {
				vi.v += e
				heap.Fix(&h, vi.i)
			}
			if len(h) == n {
				g = g * pow(int64(p), h[0].v-before) % mod
			}
		}
		Fprintln(out, g)
	}
}

//func main() { CF1493D(os.Stdin, os.Stdout) }
