package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair28 struct{ x, c int }
type hp28 []pair28

func f28(x, c int) int64          { l := int64(x / c); return int64(c-x%c)*l*l + int64(x%c)*(l+1)*(l+1) }
func dec28(p pair28) int64        { return f28(p.x, p.c) - f28(p.x, p.c+1) }
func (h hp28) Len() int           { return len(h) }
func (h hp28) Less(i, j int) bool { return dec28(h[i]) > dec28(h[j]) }
func (h hp28) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp28) Push(interface{})     {}
func (hp28) Pop() (_ interface{}) { return }

func CF1428E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v int
	var ans int64
	Fscan(in, &n, &m)
	h := make(hp28, n)
	for i := range h {
		Fscan(in, &v)
		ans += int64(v) * int64(v)
		h[i] = pair28{v, 1}
	}
	heap.Init(&h)
	for i := n; i < m; i++ {
		ans -= dec28(h[0])
		h[0].c++
		heap.Fix(&h, 0)
	}
	Fprint(out, ans)
}

//func main() { CF1428E(os.Stdin, os.Stdout) }
