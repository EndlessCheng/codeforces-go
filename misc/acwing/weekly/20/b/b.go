package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, k2, v, ans int
	Fscan(in, &n, &k, &k2)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range a {
		Fscan(in, &v)
		a[i] = abs(a[i] - v)
	}

	h := hp{a}
	heap.Init(&h)
	for k += k2; k > 0; k-- {
		if h.IntSlice[0] > 0 {
			h.IntSlice[0]--
		} else {
			h.IntSlice[0] = 1
		}
		heap.Fix(&h, 0)
	}
	for _, v := range a {
		ans += v * v
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
