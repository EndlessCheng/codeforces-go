package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1263C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := []any{}
		for l, r := 1, 0; l <= n; l = r + 1 {
			h := n / l
			ans = append(ans, h)
			r = n / h
		}
		ans = append(ans, 0)
		slices.Reverse(ans)
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1263C(os.Stdin, os.Stdout) }
