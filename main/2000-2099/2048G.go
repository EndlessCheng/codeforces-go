package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2048G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
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

	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &v)
		f := func(x, y int) int {
			sum, C := 0, 1
			for i := range n + 1 {
				x1 := pow(v, n-i) * pow(x-1, i) % mod
				x2 := pow(v-y, n-i) * pow(x-y-1, i) % mod
				x1 = (x1 - x2 + mod) % mod
				x1 = pow(x1, m)
				x1 = x1 * C % mod
				if i&1 != 0 {
					sum = (sum - x1 + mod) % mod
				} else {
					sum = (sum + x1) % mod
				}
				C = C * (n - i) % mod * pow(i+1, mod-2) % mod
			}
			return sum
		}

		ans := pow(v, n*m)
		for i := 1; i < v; i++ {
			ans = (ans - f(i+1, i) + mod) % mod
		}
		for i := 2; i < v; i++ {
			ans = (ans + f(i+1, i-1)) % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2048G(bufio.NewReader(os.Stdin), os.Stdout) }
