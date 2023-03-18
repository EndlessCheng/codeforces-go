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
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			d[j]++
		}
	}
	for i := 1; i <= n; i++ {
		ans += d[i] * d[n-i]
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
