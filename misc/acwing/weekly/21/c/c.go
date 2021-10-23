package main

import (
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	var T, a, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &m)
		Fprintln(out, phi(m/gcd(a, m)))
	}
}

func main() { run(os.Stdin, os.Stdout) }

func phi(n int) int {
	res := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			res = res / i * (i - 1)
			for ; n%i == 0; n /= i {
			}
		}
	}
	if n > 1 {
		res = res / n * (n - 1)
	}
	return res
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
