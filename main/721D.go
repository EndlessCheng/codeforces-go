package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

type pair struct {
	v int64
	i int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].v < h[j].v }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// github.com/EndlessCheng/codeforces-go
func CF721D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, k, x int64
	Fscan(in, &n, &k, &x)
	sign, minI := int64(1), 0
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] < 0 {
			sign = -sign
		}
		if abs(a[i]) < abs(a[minI]) {
			minI = i
		}
	}

	if sign > 0 {
		if abs(a[minI]) >= k*x {
			for i, v := range a {
				if i == minI {
					if v < 0 {
						v += k * x
					} else {
						v -= k * x
					}
				}
				Fprint(out, v, " ")
			}
			return
		}
		c := abs(a[minI])/x + 1
		if a[minI] < 0 {
			a[minI] += c * x
		} else {
			a[minI] -= c * x
		}
		k -= c
	}
	h := make(hp, n)
	for i, v := range a {
		h[i] = pair{abs(v), i}
	}
	heap.Init(&h)
	for ; k > 0; k-- {
		h[0].v += x
		heap.Fix(&h, 0)
	}
	ans := make([]interface{}, n)
	for _, p := range h {
		if a[p.i] < 0 {
			ans[p.i] = -p.v
		} else {
			ans[p.i] = p.v
		}
	}
	Fprint(out, ans...)
}

//func main() { CF721D(os.Stdin, os.Stdout) }
