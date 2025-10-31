package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p12137(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	var n, v, s, ans int
	Fscan(in, &n)
	pow3 := pow(3, n-2)
	const inv3 = 333333336
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s ^= v
		if i < n-1 {
			ans = (ans + s*2*pow3) % mod
			pow3 = pow3 * inv3 % mod
		} else {
			ans = (ans + s) % mod
		}
	}
	Fprint(out, ans)
}

//func main() { p12137(bufio.NewReader(os.Stdin), os.Stdout) }
