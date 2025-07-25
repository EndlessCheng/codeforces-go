package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1677D(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		ans := 1
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if i <= k {
				ans = ans * i % mod
			}
			if i <= n-k {
				if v == -1 {
					ans = ans * (i + k) % mod
				}
				if v == 0 {
					ans = ans * (k + 1) % mod
				}
			} else if v > 0 {
				ans = 0
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1677D(bufio.NewReader(os.Stdin), os.Stdout) }
