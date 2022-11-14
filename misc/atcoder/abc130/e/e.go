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
	const mod int = 1e9 + 7
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i, v := range a {
		for j, w := range b {
			if v == w {
				f[i+1][j+1] = (f[i][j+1] + f[i+1][j] + 1) % mod
			} else {
				f[i+1][j+1] = (f[i][j+1] + f[i+1][j] - f[i][j]) % mod // 二维前缀和
			}
		}
	}
	Fprint(out, (f[n][m]+1+mod)%mod)
}

func main() { run(os.Stdin, os.Stdout) }
