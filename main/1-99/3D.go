package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf3D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var s []byte
	var l, r, ans, cnt int
	Fscan(in, &s)
	h := hp03{}
	for i, c := range s {
		if c == '(' {
			cnt++
			continue
		}
		if c == '?' {
			s[i] = ')'
			Fscan(in, &l, &r)
			ans += r
			heap.Push(&h, pair03{l - r, i})
		}
		if cnt > 0 {
			cnt--
			continue
		}
		if len(h) == 0 {
			Fprint(out, -1)
			return
		}
		cnt++
		p := heap.Pop(&h).(pair03)
		ans += p.v
		s[p.i] = '('
	}
	if cnt > 0 {
		Fprint(out, -1)
		return
	}
	Fprintln(out, ans)
	Fprintf(out, "%s", s)
}

//func main() { cf3D(os.Stdin, os.Stdout) }
type pair03 struct{ v, i int }
type hp03 []pair03
func (h hp03) Len() int           { return len(h) }
func (h hp03) Less(i, j int) bool { return h[i].v < h[j].v }
func (h hp03) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp03) Push(v any)        { *h = append(*h, v.(pair03)) }
func (h *hp03) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
