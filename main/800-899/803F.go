package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf803F(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	const mx int = 1e5
	var cnt, f [mx + 1]int
	var n, v int
	Fscan(in, &n)
	pow2 := make([]int, n+1)
	pow2[0] = 1
	for i := range n {
		Fscan(in, &v)
		cnt[v]++
		pow2[i+1] = pow2[i] * 2 % mod
	}

	for i := mx; i > 0; i-- {
		c := 0
		for j := i; j <= mx; j += i {
			c += cnt[j]
			f[i] -= f[j]
		}
		f[i] = (f[i] + pow2[c] - 1) % mod
	}
	Fprint(out, (f[1]+mod)%mod)
}

//func main() { cf803F(bufio.NewReader(os.Stdin), os.Stdout) }
