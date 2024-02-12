package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF295B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	d := make([][]int, n)
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			Fscan(in, &d[i][j])
		}
	}
	ans := make([]int64, n)
	for i := range ans {
		Fscan(in, &ans[i])
		ans[i]--
	}

	vis := make([]bool, n)
	for q := n - 1; q >= 0; q-- {
		k := ans[q]
		vis[k] = true
		ans[q] = 0
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
				if vis[i] && vis[j] {
					ans[q] += int64(d[i][j])
				}
			}
		}
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF295B(os.Stdin, os.Stdout) }
