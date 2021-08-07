package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1012B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, x, y int
	Fscan(in, &n, &m, &k)
	fa := make([]int, n+m)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	ans := n + m - 1
	for ; k > 0; k-- {
		Fscan(in, &x, &y)
		x, y = find(x-1), find(y-1+n)
		if x != y {
			fa[x] = y
			ans--
		}
	}
	Fprint(out, ans)
}

//func main() { CF1012B(os.Stdin, os.Stdout) }
