package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2000F(in io.Reader, out io.Writer) {
	var T, n, k, a, b int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		f := make([]int, k+1)
		for i := 1; i <= k; i++ {
			f[i] = 1e9
		}
		for range n {
			Fscan(in, &a, &b)
			if a > b {
				a, b = b, a
			}
			for j := k; j > 0; j-- {
				w, v := 0, 0
				dp := func(dw, dv int) {
					w += dw
					v += dv
					f[j] = min(f[j], f[max(j-w, 0)]+v)
				}
				for range b - a {
					dp(1, a)
				}
				for i := a; i > 1; i-- {
					dp(1, i)
					dp(1, i-1)
				}
				dp(2, 1)
			}
		}
		if f[k] == 1e9 {
			f[k] = -1
		}
		Fprintln(out, f[k])
	}
}

//func main() { cf2000F(bufio.NewReader(os.Stdin), os.Stdout) }
