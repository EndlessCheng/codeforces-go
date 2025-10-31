package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3147(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int32
	Fscan(in, &n)
	f := [59][]int32{}
	for i := range f {
		f[i] = make([]int32, n+2)
	}
	for i := int32(1); i <= n; i++ {
		Fscan(in, &v)
		f[v][i] = i + 1
	}
	ans := 1
	for i := 2; i < len(f); i++ {
		for j := int32(1); j <= n; j++ {
			if f[i][j] == 0 {
				f[i][j] = f[i-1][f[i-1][j]]
			}
			if f[i][j] > 0 {
				ans = i
			}
		}
	}
	Fprint(out, ans)
}

//func main() { p3147(os.Stdin, os.Stdout) }
