package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf476A(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	ans := (n + 1) / 2
	ans += (m - ans%m) % m
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { cf476A(os.Stdin, os.Stdout) }
