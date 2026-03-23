package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1864C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := []any{n}
		for n&(n-1) > 0 {
			n &= n - 1
			ans = append(ans, n)
		}
		for n > 1 {
			n /= 2
			ans = append(ans, n)
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1864C(os.Stdin, os.Stdout) }
