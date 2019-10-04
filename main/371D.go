package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol371D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	water := make([]int, n)
	cp := make([]int, n)
	for i := range cp {
		Fscan(in, &cp[i])
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(from, to int) { fa[find(from)] = find(to) }

	var m, op, idx, x int
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &op, &idx)
		idx--
		if op == 1 {
			Fscan(in, &x)
			for idx = find(idx); ; idx = find(idx) {
				if water[idx]+x < cp[idx] {
					water[idx] += x
					break
				}
				x -= cp[idx] - water[idx]
				water[idx] = cp[idx]
				idx++
				if idx == n {
					break
				}
				merge(idx-1, idx)
			}
		} else {
			Fprintln(out, water[idx])
		}
	}
}

//func main() {
//	Sol371D(os.Stdin, os.Stdout)
//}
