package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf235E(in io.Reader, out io.Writer) {
	const mod = 1 << 30
	var a, b, c int
	Fscan(in, &a, &b, &c)
	mu := make([]int, c+1)
	mu[1] = 1
	for i := 1; i <= c; i++ {
		if mu[i] != 0 {
			for j := i * 2; j <= c; j += i {
				mu[j] -= mu[i]
			}
		}
	}

	n := a * b
	f := make([]int, n+1)
	for i := 1; i <= c; i++ {
		cnt := 0
		for j := i; j <= c; j += i {
			cnt += c / j
		}
		for j := i; j <= n; j += i {
			f[j] = (f[j] + cnt*mu[i]) % mod
		}
	}

	ans := 0
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if gcd(i, j) == 1 {
				ans = (ans + (a/i)*(b/j)%mod*f[i*j]) % mod
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf235E(bufio.NewReader(os.Stdin), os.Stdout) }

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
