package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1202F(in io.Reader, out io.Writer) {
	var a, b, ans int
	Fscan(in, &a, &b)
	n := a + b
	for l, r := 1, 0; l <= n; l = r + 1 {
		q := n / l
		r = n / q
		if q > min(a, b) {
			continue
		}
		l1, r1 := (a+q)/(q+1), a/q
		l2, r2 := (b+q)/(q+1), b/q
		if l1 <= r1 && l2 <= r2 {
			ans += max(min(r1+r2, r)-max(l1+l2, l)+1, 0)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1202F(os.Stdin, os.Stdout) }
