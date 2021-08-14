package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1542C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := n
		for i, lcm := int64(2), int64(1); lcm <= n; i++ {
			ans += n / lcm
			lcm *= i / gcd(i, lcm)
		}
		Fprintln(out, ans%mod)
	}
}

//func main() { CF1542C(os.Stdin, os.Stdout) }
