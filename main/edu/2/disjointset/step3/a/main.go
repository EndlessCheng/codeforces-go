package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, w int
	var op string
	Fscan(in, &n, &q)
	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	undo := []int{}
	checkPoints := []int{}
	cc := n
	find := func(x int) int {
		for ; x != fa[x]; x = fa[x] {
		}
		return x
	}
	merge := func(x, y int) {
		if x, y = find(x), find(y); x != y {
			if sz[x] > sz[y] {
				x, y = y, x
			}
			fa[x] = y
			sz[y] += sz[x]
			undo = append(undo, x)
			cc--
		}
	}
	rollback := func(tar int) {
		for len(undo) > tar {
			x := undo[len(undo)-1]
			undo = undo[:len(undo)-1]
			sz[fa[x]] -= sz[x]
			fa[x] = x
			cc++
		}
	}
	for ; q > 0; q-- {
		if Fscan(in, &op); op[0] == 'u' {
			Fscan(in, &v, &w)
			merge(v, w)
			Fprintln(out, cc)
		} else if op[0] == 'p' {
			checkPoints = append(checkPoints, len(undo))
		} else {
			rollback(checkPoints[len(checkPoints)-1])
			checkPoints = checkPoints[:len(checkPoints)-1]
			Fprintln(out, cc)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
