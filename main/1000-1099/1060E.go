package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1060E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, s int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := int64(0)
	var f func(v, fa, d int) int
	f = func(v, fa, d int) int {
		s += d
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v, d^1)
			}
		}
		ans += int64(sz) * int64(n-sz)
		return sz
	}
	f(1, 0, 0)
	Fprint(out, (ans+int64(s)*int64(n-s))/2) // 额外加上奇距离的个数（奇数层*偶数层）
}

//func main() { CF1060E(os.Stdin, os.Stdout) }
