package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1985F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, left, n, ans int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &left, &n)
		h := make(hp85, n)
		for i := range h {
			Fscan(in, &h[i].d)
		}
		for i := range h {
			Fscan(in, &h[i].cd)
		}
		for left > 0 {
			ans = h[0].t
			left -= h[0].d
			h[0].t += h[0].cd
			heap.Fix(&h, 0)
		}
		Fprintln(out, ans+1)
	}
}

//func main() { cf1985F(bufio.NewReader(os.Stdin), os.Stdout) }
type hp85 []struct{ t, d, cd int }
func (h hp85) Len() int           { return len(h) }
func (h hp85) Less(i, j int) bool { return h[i].t < h[j].t }
func (h hp85) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp85) Push(any)             {}
func (hp85) Pop() (_ any)         { return }
