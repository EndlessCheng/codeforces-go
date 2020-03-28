package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

// github.com/EndlessCheng/codeforces-go
func CF1183G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var t, n, tp, f int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		types := make([][2]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &tp, &f)
			types[tp][f]++
		}
		cnts := make([][]int, n+1)
		for _, t := range types {
			s := t[0] + t[1]
			cnts[s] = append(cnts[s], t[1])
		}
		h := &hp{}
		sum, sumF1 := 0, 0
		for i := n; i > 0; i-- {
			for _, f1 := range cnts[i] {
				Push(h, f1)
			}
			if len(h.IntSlice) > 0 {
				sum += i
				sumF1 += min(Pop(h).(int), i)
			}
		}
		Fprintln(out, sum, sumF1)
	}
}

//func main() { CF1183G(os.Stdin, os.Stdout) }
