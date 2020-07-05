package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1025B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int64) int64 {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int64) int64 { return a / gcd(a, b) * b }

	var n int
	var x, y, a, b int64
	Fscan(in, &n, &x, &y)
	g := lcm(x, y)
	for n--; n > 0; n-- {
		Fscan(in, &a, &b)
		g = gcd(g, lcm(a, b))
	}
	f := func(n int64) int64 {
		for i := int64(2); i*i <= n; i++ {
			if n%i == 0 && g%i == 0 {
				return i
			}
			for ; n%i == 0; n /= i {
			}
		}
		if n > 1 && g%n == 0 {
			return n
		}
		return -1
	}
	ans := f(x)
	if ans == -1 {
		ans = f(y)
	}
	Fprint(out, ans)
}

//func main() { CF1025B(os.Stdin, os.Stdout) }
