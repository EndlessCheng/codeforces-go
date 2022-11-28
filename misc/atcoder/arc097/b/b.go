package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, x, y, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	fa := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	for ; m > 0; m-- {
		Fscan(in, &x, &y)
		fa[find(x-1)] = find(y - 1)
	}
	for i, v := range a {
		if find(i) == find(v-1) {
			ans++
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
