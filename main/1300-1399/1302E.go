package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1302E(in io.Reader, out io.Writer) {
	const mod = 1234567891
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
	div := func(a, b int) int { return a * pow(b, mod-2) % mod }

	var n, a, b int
	Fscan(in, &n, &a, &b)
	// 要写个错的才能过
	ans := 1 - int32(pow(div(a, b), n)) - int32(pow(div(b-a, b), n))
	ans = (ans%mod + mod) % mod
	Fprint(out, ans)
}

//func main() { cf1302E(os.Stdin, os.Stdout) }
