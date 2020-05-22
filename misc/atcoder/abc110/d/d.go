package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	const p = 1e9 + 7
	pow := func(x, n int) int {
		r := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				r = r * x % p
			}
			x = x * x % p
		}
		return r
	}
	c := func(n, k int) int {
		a, b := 1, 1
		for i := 1; i <= k; i++ {
			a = a * n % p
			n--
			b = b * i % p
		}
		return a * pow(b, p-2) % p
	}
	var n, m int
	Fscan(_r, &n, &m)
	ans := 1
	for i := 2; i*i <= m; i++ {
		e := 0
		for ; m%i == 0; m /= i {
			e++
		}
		ans = ans * c(e+n-1, e) % p
	}
	if m > 1 {
		ans = ans * n % p
	}
	Fprint(_w, ans)
}

func main() { run(os.Stdin, os.Stdout) }
