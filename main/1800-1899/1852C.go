package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1852C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		h := hp52{}
		ans, pre := 0, 0
		for range n {
			Fscan(in, &v)
			d := v%k - pre
			if d < 0 {
				heap.Push(&h, d)
			} else if h.Len() == 0 || h.IntSlice[0]+k >= d {
				ans += d
			} else {
				ans += h.IntSlice[0] + k
				h.IntSlice[0] = d - k
				heap.Fix(&h, 0)
			}
			pre = v % k
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1852C(bufio.NewReader(os.Stdin), os.Stdout) }
type hp52 struct{ sort.IntSlice }
func (h *hp52) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp52) Pop() (_ any)  { return }
