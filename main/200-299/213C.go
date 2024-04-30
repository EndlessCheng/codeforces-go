package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf213C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = -1e9
		}
	}
	f[1][1] = a[0][0]
	for s := 1; s < n*2-1; s++ {
		for j := min(n-1, s); j >= max(s-n+1, 0); j-- {
			for k := min(n-1, s); k >= j; k-- {
				f[j+1][k+1] = max(f[j+1][k+1], f[j][k+1], f[j+1][k], f[j][k]) + a[s-j][j]
				if k != j {
					f[j+1][k+1] += a[s-k][k]
				}
			}
		}
	}
	Fprint(out, f[n][n])
}

//func main() { cf213C(os.Stdin, os.Stdout) }
