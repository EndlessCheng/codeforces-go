package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, m, v int
	Fscan(in, &n, &q)
	p := make([]int, n+1)
	child := make([]int, n+1)
	for i := 2; i <= n; i++ {
		Fscan(in, &p[i])
		child[p[i]]++
	}

	for ; q > 0; q-- {
		Fscan(in, &m)
		has := make(map[int]bool, m)
		for ; m > 0; m-- {
			Fscan(in, &v)
			has[v] = true
		}
		ans := 0
		for v := range has {
			ans += child[v]
			if has[p[v]] {
				ans--
			} else {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
