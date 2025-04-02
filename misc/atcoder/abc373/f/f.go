package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, w, v int
	Fscan(in, &n, &m)
	g := make([]hp, m+1)
	for ; n > 0; n-- {
		Fscan(in, &w, &v)
		g[w].IntSlice = append(g[w].IntSlice, v-1)
	}

	f := make([]int, m+1)
	val := make([]int, m+1)
	for i, h := range g {
		if h.IntSlice == nil {
			continue
		}
		heap.Init(&h)
		for j := 1; j <= m/i; j++ {
			val[j] = val[j-1] + h.IntSlice[0]
			h.IntSlice[0] -= 2
			heap.Fix(&h, 0)
		}
		for j := m; j >= i; j-- {
			for k := 1; k*i <= j; k++ {
				f[j] = max(f[j], f[j-k*i]+val[k])
			}
		}
	}
	Fprint(out, f[m])
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
