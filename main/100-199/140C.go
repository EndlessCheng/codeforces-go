package _00_199

import (
	"bufio"
	. "fmt"
	"io"
	. "container/heap"
	"sort"
)

type pair140C struct {
	r, cnt int
}
type pairHeap140C []pair140C

func (h pairHeap140C) Len() int              { return len(h) }
func (h pairHeap140C) Less(i, j int) bool    { return h[i].cnt > h[j].cnt }
func (h pairHeap140C) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *pairHeap140C) Push(v interface{})   { *h = append(*h, v.(pair140C)) }
func (h *pairHeap140C) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

// github.com/EndlessCheng/codeforces-go
func Sol140C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	rCnt := map[int]int{}
	var n, r int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &r)
		rCnt[r]++
	}

	_h := make(pairHeap140C, 0, len(rCnt))
	for r, cnt := range rCnt {
		_h = append(_h, pair140C{r, cnt})
	}
	h := &_h
	Init(h)

	ans := [][3]int{}
	for len(_h) >= 3 {
		ps := [3]pair140C{Pop(h).(pair140C), Pop(h).(pair140C), Pop(h).(pair140C)}
		ans = append(ans, [3]int{ps[0].r, ps[1].r, ps[2].r})
		for _, p := range ps {
			p.cnt--
			if p.cnt > 0 {
				Push(h, p)
			}
		}
	}
	Fprintln(out, len(ans))
	for _, rs := range ans {
		balls := rs[:]
		sort.Ints(balls)
		Fprintln(out, balls[2], balls[1], balls[0])
	}
}

//func main() {
//	Sol140C(os.Stdin, os.Stdout)
//}
