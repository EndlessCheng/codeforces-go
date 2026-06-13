package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var T, n, m uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := uint(0)
		for p10 := uint(1); p10 <= n; p10 *= 10 {
			m2 := m / gcd(p10*10-1, m)
			mx := min(p10*10-1, n)
			ans = (ans + n/m2%mod*((mx-p10+1)%mod)) % mod
		}
		Fprintln(out, ans)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

func gcd(a, b uint) uint {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
