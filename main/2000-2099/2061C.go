package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2061C(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n+1)
		f := make([]int, n+1)
		f[0] = 1
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			if a[i] == a[i-1] {
				f[i] = f[i-1]
			}
			if i > 1 && a[i]-a[i-2] == 1 {
				f[i] = (f[i] + f[i-2]) % mod
			}
		}
		Fprintln(out, (f[n-1]+f[n])%mod)
	}
}

//func main() { cf2061C(bufio.NewReader(os.Stdin), os.Stdout) }
