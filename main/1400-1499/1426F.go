package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1426F(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	var n int
	var s string
	Fscan(bufio.NewReader(in), &n, &s)
	pow3 := make([]int64, n)
	pow3[0] = 1
	for i := 1; i < n; i++ {
		pow3[i] = pow3[i-1] * 3 % mod
	}
	c := int64(strings.Count(s, "c"))
	r := int64(strings.Count(s, "?"))
	var ans, a, l int64
	for _, b := range s {
		if b == 'a' {
			a++
		} else if b == 'c' {
			c--
		} else {
			if b == '?' {
				r--
			}
			ans += a * c % mod * pow3[l+r] % mod
			if l > 0 {
				ans += l * c % mod * pow3[l-1+r] % mod
			}
			if r > 0 {
				ans += a * r % mod * pow3[l+r-1] % mod
			}
			if l > 0 && r > 0 {
				ans += l * r % mod * pow3[l+r-2] % mod
			}
			if b == '?' {
				l++
			}
		}
	}
	Fprint(out, ans%mod)
}

//func main() { CF1426F(os.Stdin, os.Stdout) }
