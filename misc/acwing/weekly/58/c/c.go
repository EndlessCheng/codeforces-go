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
	var n int
	Fscan(in, &n)
	fa := make([]int, n+1)
	for i := 2; i <= n; i++ {
		Fscan(in, &fa[i])
	}
	c := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &c[i])
	}
	ans := 1
	for i := 2; i <= n; i++ {
		if c[i] != c[fa[i]] {
			ans++
		}
	}
	Fprintln(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
