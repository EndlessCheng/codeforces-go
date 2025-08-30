package main

import (
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1784(in io.Reader, out io.Writer) {
	a := [9][9]int{}
	for i := range a {
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	rowHas := [9][9]bool{}
	colHas := [9][9]bool{}
	subBoxHas := [3][3][9]bool{}
	emptyPos := [][2]int{}

	for i, row := range a {
		for j, b := range row {
			if b == 0 {
				emptyPos = append(emptyPos, [2]int{i, j})
			} else {
				x := b - 1
				rowHas[i][x] = true
				colHas[j][x] = true
				subBoxHas[i/3][j/3][x] = true
			}
		}
	}

	getCandidates := func(i, j int) int {
		candidates := 9
		for x := 0; x < 9; x++ {
			if rowHas[i][x] || colHas[j][x] || subBoxHas[i/3][j/3][x] {
				candidates--
			}
		}
		return candidates
	}

	emptyHeap := &hp1784{}
	for _, pos := range emptyPos {
		i, j := pos[0], pos[1]
		heap.Push(emptyHeap, tuple1784{getCandidates(i, j), i, j})
	}
	heap.Init(emptyHeap)

	var dfs func() bool
	dfs = func() bool {
		if emptyHeap.Len() == 0 {
			return true
		}

		t := heap.Pop(emptyHeap).(tuple1784)
		i, j := t.i, t.j

		candidates := 0
		for x := 0; x < 9; x++ {
			if rowHas[i][x] || colHas[j][x] || subBoxHas[i/3][j/3][x] {
				continue
			}

			a[i][j] = 1 + x
			rowHas[i][x] = true
			colHas[j][x] = true
			subBoxHas[i/3][j/3][x] = true

			if dfs() {
				return true
			}

			rowHas[i][x] = false
			colHas[j][x] = false
			subBoxHas[i/3][j/3][x] = false
			candidates++
		}

		heap.Push(emptyHeap, tuple1784{candidates, i, j})
		return false
	}
	dfs()

	for _, r := range a {
		for _, v := range r {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { p1784(os.Stdin, os.Stdout) }

type tuple1784 struct{ candidates, i, j int }
type hp1784 []tuple1784
func (h hp1784) Len() int           { return len(h) }
func (h hp1784) Less(i, j int) bool { return h[i].candidates < h[j].candidates }
func (h hp1784) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp1784) Push(v any)        { *h = append(*h, v.(tuple1784)) }
func (h *hp1784) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
