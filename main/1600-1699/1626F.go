package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1626F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, a, x, y, k, M int
	Fscan(in, &n, &a, &x, &y, &k, &M)

	m := 1
	w := k
	for i := 1; i < k; i++ {
		m = m * i / gcd26(m, i)
		w = w * n % mod
	}

	f := make([]int, m)
	pw := 1
	n--
	for i := k; i > 0; i-- {
		for j := m - 1; j >= 0; j-- {
			f[j] = (f[j]*n + f[j/i*i] + j*pw) % mod
		}
		pw = pw * (n + 1) % mod
	}

	ans := 0
	for i := 0; i <= n; i++ {
		ans = (ans + f[a%m] + a/m*m%mod*w) % mod
		a = (a*x + y) % M
	}
	Fprint(out, ans)
}

//func main() { cf1626F(bufio.NewReader(os.Stdin), os.Stdout) }

func gcd26(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
