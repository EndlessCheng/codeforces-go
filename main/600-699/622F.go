package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf622F(in io.Reader, out io.Writer) {
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
	var n, k, fn, s int
	Fscan(in, &n, &k)

	f := make([]int, k+3)
	f[0] = 1
	for i := 1; i <= k+2; i++ {
		f[i] = f[i-1] * i % mod
	}

	suf := make([]int, k+3)
	suf[k+2] = 1
	for i := k + 2; i > 0; i-- {
		suf[i-1] = suf[i] * (n - i) % mod
	}

	pre := 1
	for i := 1; i <= k+2; i++ {
		a := pre * suf[i] % mod
		b := f[i-1] * f[k+2-i] % mod
		if (k-i)%2 != 0 {
			b = -b
		}
		s = (s + pow(i, k)) % mod
		fn += a * pow(b, mod-2) % mod * s % mod
		pre = pre * (n - i) % mod
	}
	Fprint(out, (fn%mod+mod)%mod)
}

//func main() { cf622F(bufio.NewReader(os.Stdin), os.Stdout) }
