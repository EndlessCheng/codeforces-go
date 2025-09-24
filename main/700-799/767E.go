package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf767E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, coin, w, sum int
	Fscan(in, &n, &coin)
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}

	ans := make([]pair67, n)
	h := hp67{}
	for i, c := range c {
		Fscan(in, &w)
		r := c % 100
		ans[i] = pair67{c / 100, r}
		if r == 0 {
			continue
		}
		coin -= r
		heap.Push(&h, pair67{w * (100 - r), i})
		if coin < 0 {
			p := heap.Pop(&h).(pair67)
			sum += p.v
			ans[p.i].v++
			ans[p.i].i = 0
			coin += 100
		}
	}

	Fprintln(out, sum)
	for _, p := range ans {
		Fprintln(out, p.v, p.i)
	}
}

//func main() { cf767E(bufio.NewReader(os.Stdin), os.Stdout) }
type pair67 struct{ v, i int }
type hp67 []pair67
func (h hp67) Len() int           { return len(h) }
func (h hp67) Less(i, j int) bool { return h[i].v < h[j].v }
func (h hp67) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp67) Push(v any)        { *h = append(*h, v.(pair67)) }
func (h *hp67) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
