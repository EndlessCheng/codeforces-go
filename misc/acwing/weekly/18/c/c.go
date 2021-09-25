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

	var n, v, w, ans int
	Fscan(in, &n)
	if n&1 > 0 { // 奇数无法拆分成若干偶数之和
		Fprintln(out, -1)
		return
	}

	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var f func(int, int) int
	f = func(v, fa int) int {
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v)
			}
		}
		if fa > 0 && sz&1 == 0 { // 由于偶数=偶数+偶数，只要子树大小是偶数就可以删边
			ans++
		}
		return sz
	}
	f(1, 0)
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
