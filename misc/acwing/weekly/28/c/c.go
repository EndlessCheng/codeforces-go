package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, d int
	Fscan(in, &n)
	fa := make([]int, n)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		fa[i] = i
	}
	for i := 0; i < n; i++ {
		Fscan(in, &d)
		if i >= d {
			fa[find(i)] = find(i - d)
		}
		if i+d < n {
			fa[find(i)] = find(i + d)
		}
	}
	for i, v := range a {
		if find(i) != find(v-1) {
			Fprint(out, "NO")
			return
		}
	}
	Fprint(out, "YES")
}

func main() { run(os.Stdin, os.Stdout) }
