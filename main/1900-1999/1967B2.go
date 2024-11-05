package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1967B2(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		ans := 0
		for a := 1; a*a <= n; a++ {
			for b := 1; a+b <= min(n/a, m/b); b++ {
				if gcd(a, b) == 1 {
					ans += min(n/a, m/b) / (a + b)
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1967B2(bufio.NewReader(os.Stdin), os.Stdout) }
