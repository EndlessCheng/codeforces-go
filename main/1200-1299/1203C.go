package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1203C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, v, g, ans int64
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		g = gcd(g, v)
	}
	for d := int64(1); d*d <= g; d++ {
		if g%d == 0 {
			ans++
			if d*d < g {
				ans++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1203C(os.Stdin, os.Stdout) }
