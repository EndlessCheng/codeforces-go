package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1572C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		pre := make([]int, n+1)
		last := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			pre[i] = last[a[i]]
			last[a[i]] = i
		}

		f := make([][]int, n+2)
		for i := range f {
			f[i] = make([]int, n+2)
		}
		for i := n; i > 0; i-- {
			for j := i + 1; j <= n; j++ {
				f[i][j] = f[i][j-1] + 1
				for k := pre[j]; k >= i; k = pre[k] {
					f[i][j] = min(f[i][j], f[i][k]+f[k+1][j])
				}
			}
		}
		Fprintln(out, f[1][n])
	}
}

//func main() { cf1572C(bufio.NewReader(os.Stdin), os.Stdout) }
