package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf830D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n int
	Fscan(in, &n)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+2)
	}
	f[1][0] = 1
	f[1][1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; j+k <= n; k++ {
				tmp := f[i-1][j] * f[i-1][k] % mod
				if j+k > 0 {
					f[i][j+k-1] = (f[i][j+k-1] + (j+k)*(j+k-1)%mod*tmp) % mod
				}
				f[i][j+k] = (f[i][j+k] + (2*(j+k)+1)*tmp) % mod
				f[i][j+k+1] = (f[i][j+k+1] + tmp) % mod
			}
		}
	}
	Fprint(out, f[n][1])
}

//func main() { cf830D(os.Stdin, os.Stdout) }
