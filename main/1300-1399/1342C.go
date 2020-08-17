package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1342C(_r io.Reader, _w io.Writer) {
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

	var t, q int
	var a, b, l, r int64
	cnt := func(x int64) int64 {
		l := lcm(a, b)
		if l == b {
			return 0
		}
		ans := x / l * (l - b)
		x %= l
		if x >= b {
			ans += x - b + 1
		}
		return ans
	}
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fscan(in, &a, &b, &q)
		if a > b {
			a, b = b, a
		}
		for ; q > 0; q-- {
			Fscan(in, &l, &r)
			Fprint(out, cnt(r)-cnt(l-1), " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1342C(os.Stdin, os.Stdout) }
