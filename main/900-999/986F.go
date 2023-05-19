package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF986F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 31622776
	primes := make([]int, 0, 1951957)
	np := [mx + 1]bool{}
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > mx {
				break
			}
			np[p*i] = true
			if i%p == 0 {
				break
			}
		}
	}
	pow := func(x, n, mod int64) int64 {
		x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	var T int
	var n, k int64
	Fscan(in, &T)
	type query struct {
		n int64
		i int
	}
	qs := map[int64][]query{}
	for i := 0; i < T; i++ {
		Fscan(in, &n, &k)
		qs[k] = append(qs[k], query{n, i})
	}
	ans := make([]bool, T)
	for k, qs := range qs {
		ps := []int64{}
		x := k
		for _, p := range primes {
			p := int64(p)
			if p > x {
				break
			}
			if x%p == 0 {
				for x /= p; x%p == 0; x /= p {
				}
				ps = append(ps, p)
			}
		}
		if x > 1 {
			ps = append(ps, x)
		}
		if len(ps) == 0 { // k = 1
		} else if len(ps) == 1 {
			for _, q := range qs {
				ans[q.i] = q.n%ps[0] == 0
			}
		} else if len(ps) == 2 {
			x, y := ps[0], ps[1]
			for _, q := range qs {
				t := q.n % x * pow(y, x-2, x) % x
				ans[q.i] = t*y <= q.n
			}
		} else {
			dis := make([]int64, ps[0])
			for i := range dis {
				dis[i] = math.MaxInt64
			}
			dis[0] = 0
			h := hp86{{}}
			for len(h) > 0 {
				top := h.pop()
				v := top.v
				if top.dis > dis[v] {
					continue
				}
				for _, p := range ps[1:] {
					w := (int64(v) + p) % ps[0]
					if newD := dis[v] + p; newD < dis[w] {
						dis[w] = newD
						h.push(pair86{int(w), newD})
					}
				}
			}
			for _, q := range qs {
				ans[q.i] = dis[q.n%ps[0]] <= q.n
			}
		}
	}
	for _, b := range ans {
		if b {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF986F(os.Stdin, os.Stdout) }

type pair86 struct {
	v   int
	dis int64
}
type hp86 []pair86

func (h hp86) Len() int              { return len(h) }
func (h hp86) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp86) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp86) Push(v interface{})   { *h = append(*h, v.(pair86)) }
func (h *hp86) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp86) push(v pair86)        { heap.Push(h, v) }
func (h *hp86) pop() pair86          { return heap.Pop(h).(pair86) }
