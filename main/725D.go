package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
	"sort"
)

type int64Heap725 []int64

func (h int64Heap725) Len() int              { return len(h) }
func (h int64Heap725) Less(i, j int) bool    { return h[i] < h[j] }
func (h int64Heap725) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *int64Heap725) Push(v interface{})   { *h = append(*h, v.(int64)) }
func (h *int64Heap725) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

// github.com/EndlessCheng/codeforces-go
func CF725D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type team struct{ balls, weight int64 }

	var n, ans int
	var myBalls, w, gives int64
	Fscan(in, &n, &myBalls, &w)
	ts := make([]team, n-1, n)
	for i := range ts {
		Fscan(in, &ts[i].balls, &ts[i].weight)
	}
	sort.Slice(ts, func(i, j int) bool { return ts[i].balls > ts[j].balls })
	ts = append(ts, team{})

	for _, t := range ts {
		if b := t.balls; b <= myBalls {
			gives += myBalls - b
			myBalls = b
			break
		}
		ans++
	}
	h := &int64Heap725{}
	for i := 0; ts[i].balls > 0; {
		for ; ts[i].balls > myBalls; i++ {
			Push(h, ts[i].weight-ts[i].balls+1)
		}
		for len(*h) > 0 {
			d := (*h)[0]
			if d > gives {
				break
			}
			Pop(h)
			gives -= d
		}
		if len(*h) < ans {
			ans = len(*h)
		}
		for _, t := range ts[i:] {
			if b := t.balls; b < myBalls {
				gives += myBalls - b
				myBalls = b
				break
			}
		}
	}
	Fprint(out, ans+1)
}

//func main() {
//	CF725D(os.Stdin, os.Stdout)
//}
