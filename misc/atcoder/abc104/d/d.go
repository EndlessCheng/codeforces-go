package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var s string
	Fscan(in, &s)
	n := len(s)
	p3 := make([]int, n)
	p3[0] = 1
	for i := 1; i < n; i++ {
		p3[i] = p3[i-1] * 3 % mod
	}
	c := strings.Count(s, "C")
	r := strings.Count(s, "?")
	var ans, a, l int
	for _, b := range s {
		if b == 'A' {
			a++
		} else if b == 'C' {
			c--
		} else {
			if b == '?' {
				r--
			}
			ans += a * c % mod * p3[l+r] % mod
			if l > 0 {
				ans += l * c % mod * p3[l-1+r] % mod
			}
			if r > 0 {
				ans += a * r % mod * p3[l+r-1] % mod
			}
			if l > 0 && r > 0 {
				ans += l * r % mod * p3[l+r-2] % mod
			}
			if b == '?' {
				l++
			}
		}
	}
	Fprint(out, ans%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
