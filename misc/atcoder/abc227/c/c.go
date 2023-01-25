package main

import (
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	for a := 1; a*a*a <= n; a++ {
		m := n / a
		for b := a; b*b <= m; b++ {
			ans += m/b - b + 1
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
