package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1943D2(in io.Reader, out io.Writer) {
	var T, n, mx, mod int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &mx, &mod)
		f := make([][]int, n+3)
		for i := range f {
			f[i] = make([]int, mx+1)
		}
		f[1][0] = 1
		preS := 1
		for i := 2; i < n+3; i++ {
			var s, ss, t int
			for j := mx; j >= 0; j-- {
				f[i][j] = (preS - ss) % mod
				t += f[i][j]
				s += f[i-2][mx-j]
				ss += s
			}
			preS = t
		}
		Fprintln(out, (f[n+2][0]+mod)%mod)
	}
}

//func main() { cf1943D2(bufio.NewReader(os.Stdin), os.Stdout) }
