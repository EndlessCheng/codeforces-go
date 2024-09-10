package main

import (
	. "fmt"
	"io"
)

func cf1920E(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, k)
		}
		for j := range f[0] {
			f[0][j] = 1
		}
		for i := 1; i <= n; i++ {
			for j := 0; j < k; j++ {
				for x := 0; x+j < k && (x+1)*(j+1) <= i; x++ {
					f[i][j] = (f[i][j] + f[i-(x+1)*(j+1)][x]) % mod
				}
			}
		}
		ans := 0
		for _, v := range f[n] {
			ans += v
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { cf1920E(bufio.NewReader(os.Stdin), os.Stdout) }
